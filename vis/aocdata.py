# Copyright 2020 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.


import numpy as np
import os
import os.path
import pandas as pd
import pytz
import re
import requests
import subprocess
from datetime import datetime, timedelta
from lxml import html


_FIRST_YEAR = 2015
_LEADERBOARD_FILE = 'cache/leaderboard.pickle'
_STATS_FILE = 'cache/stats.pickle'
_GOBENCH_FILE = 'cache/gobench.pickle'


def leaderboard():
    """Loads the daily leaderboard data frame.

    The returned frame will have a three-level hierarchical index, consisting of the integer columns
    `year`, `day` and `rank` with the obvious semantics. It will have two data series, `one_star`
    and `two_stars`, representing the time (in seconds) it took to get that specific leaderboard
    position for the first part or the full puzzle, respectively.
    """
    return pd.read_pickle(_LEADERBOARD_FILE)


def stats():
    """Loads the "solutions over time" statistics data frame.

    The returned frame will have a three-level index consisting of the integer columns `year`, `ts`
    (nominal timestamp of when the sample was taken) and `day`. There will be records for all 25
    days for each of the timestamps, though of course puzzles not yet unlocked will contain zeros.
    It will have two data series, `one_star` and `two_stars`, representing the number of solutions
    that have obtained only one star, or both stars, respectively.
    """
    return pd.read_pickle(_STATS_FILE)


def gobench():
    """Loads the daily Go benchmark data frame.

    The returned frame will have a two-level index consisting of the integer columns `year` and
    `day`, with the obvious meaning. It will have a single data series, `runtime`, denoting the time
    (in seconds) it takes for the corresponding Go solution to solve the puzzle input.
    """
    return pd.read_pickle(_GOBENCH_FILE)


def leaderboard_update():
    """Ensures the leaderboard data is up to date.
    
    Returns:
        False, if no changes were detected. True, if the data was regenerated.
    """

    new_data = False
    for year, days in contest_days():
        for day in days:
            if _leaderboard_fetch(year, day):
                new_data = True
    
    if not new_data and os.path.exists(_LEADERBOARD_FILE):
        print('leaderboard: no changes')
        return False
    print('leaderboard: regenerating data')

    data = pd.DataFrame()
    for year, days in contest_days():
        year_frame = pd.DataFrame()
        for day in days:
            year_frame = year_frame.append(_leaderboard_parse(year, day))
        data = data.append(year_frame)

    data.to_pickle(_LEADERBOARD_FILE)
    return True


def _leaderboard_fetch(year, day):
    """Fetches the daily leaderboard to the cache.

    Args:
        year: Contest year, 2015 to current.
        day: Contest day, 1 to 25 (unless in the future).
    
    Returns:
        True if a new file was fetched, False if it was unnecessary.
    
    Raises:
        requests.exceptions.RequestException: Network problems.
        IOError: Problems writing the cache file.
    """
    url = f'https://adventofcode.com/{year}/leaderboard/day/{day}'
    out = _cachefile('leaderboard', year, day)
    if os.path.exists(out):
        return False
    print(f'leaderboard: {url} -> {out}')
    r = requests.get(url, timeout=30)
    r.raise_for_status()
    with open(out, 'wb') as f:
        f.write(r.content)
    return True


def _leaderboard_parse(year, day):
    """Parses a cached leaderboard file into a dataframe.

    Args:
        year: Contest year, 2015 to current.
        day: Contest day, 1 to 25 (unless in the future).
    
    Returns:
        A Pandas dataframe containing that day's data. It will have a three-level index with the
        columns `year`, `day` and `rank` (the latter with values ranging from 1 to 100), and two
        series, `one_star` and `two_stars`, containing the time (in seconds) that yielded that given
        ranking, for the first part or both parts respectively.
    """
    page = html.parse(_cachefile('leaderboard', year, day))
    entries = page.xpath('//div[@class="leaderboard-entry"]')
    if len(entries) != 200:
        raise RuntimeError('number of entries not 200')

    times = np.empty((2, 100))
    times.fill(np.nan)
    for i, e in enumerate(entries):
        text = e.xpath('span[@class="leaderboard-time"]')[0].text_content()
        d = datetime.strptime(text, '%b %d  %H:%M:%S') - datetime(1900, 12, day, 0, 0, 0)
        times[i//100, i%100] = d.total_seconds()
    
    ranks = pd.MultiIndex.from_product(
        [[year], [day], range(1, 101)], names=('year', 'day', 'rank'))
    return pd.DataFrame({'one_star': times[1,:], 'two_stars': times[0,:]}, index=ranks)


def stats_update():
    """Ensures the stats data is up to date.

    Returns:
        False, if no changes were detected. True, if the data was regenerated.
    """

    input_times = []
    for year, _ in contest_days():
        src = _stats_source(year)
        if os.path.exists(src):
            input_times.append((year, src, os.stat(src).st_mtime))
    output_time = 0
    if os.path.exists(_STATS_FILE):
        output_time = os.stat(_STATS_FILE).st_mtime

    if not input_times or max(it[2] for it in input_times) < output_time:
        print('stats: no changes')
        return False
    print('stats: regenerating data')

    data = pd.DataFrame()
    for year, src, _ in input_times:
        year_frame = _stats_parse(year, src)
        data = data.append(year_frame)

    data.to_pickle(_STATS_FILE)
    return True


def _stats_parse(year, src):
    """Parses a collected stats file into a dataframe.

    Args:
        year: Contest year, 2015 to current.
        src: Source file path to read statistics from.

    Returns:
        A Pandas dataframe containing that year's statistics. It will have a three-level index with
        the columns `year`, `ts` (timestamp in seconds) and `day`, and two series, `one_star` and
        `two_stars`, containing the number of solutions for one or both parts of that day's puzzle.
    """

    epoch = pd.to_datetime(f'{year}-12-01').tz_localize('US/Eastern').value // 1000000000

    with open(src) as f:
        lines = list(f)

    index = []
    data = []
    for i, line in enumerate(lines):
        columns = line.split()
        if len(columns) != 52:
            raise RuntimeError(f'malformatted statistics: {columns}')
        ts = int(float(columns[0]))
        for day0 in range(25):
            since = ts - (epoch + day0*86400)
            if since < 0:
                continue
            index.append((year, day0+1, _stats_sampleidx(since), i))
            data.append((ts, since, int(columns[2+day0]), int(columns[27+day0])))

    index = pd.MultiIndex.from_tuples(index, names=('year', 'day', 'sidx', 'row'))
    data = pd.DataFrame(data, index=index, columns=('ts', 'since', 'one_star', 'two_stars'))

    return data.groupby(['year', 'day', 'sidx']).mean()


def _stats_sampleidx(since):
    if since < 86400:
        return since // 300
    elif since < 7*86400:
        return 288 + (since-86400) // 3600
    else:
        return 456 + (since-7*86400) // 86400



def _stats_source(year):
    """Returns the path to the stats source file of the given year."""
    return f'stats/stats.{year}.txt'


def gobench_update():
    """Ensures the Go benchmark data is up to date.

    Returns:
        False, if no changes were detected. True, if the data was regenerated.
    """

    new_data = False
    for year, days in _gobench_days():
        for day in days:
            if _gobench_compute(year, day):
                new_data = True

    if not new_data and os.path.exists(_GOBENCH_FILE):
        print('gobench: no changes')
        return False
    print('gobench: regenerating data')

    data = pd.DataFrame()
    for year, days in _gobench_days():
        year_frame = pd.DataFrame()
        for day in days:
            year_frame = year_frame.append(_gobench_parse(year, day))
        data = data.append(year_frame)

    data.to_pickle(_GOBENCH_FILE)
    return True


def _gobench_compute(year, day):
    """Benchmarks a single day's Go solution.

    Args:
        year: Contest year, 2015 to current.
        day: Contest day, 1 to 25 (unless in the future).

    Returns:
        True if a new file was fetched, False if it was unnecessary.

    Raises:
        requests.exceptions.RequestException: Network problems.
        IOError: Problems writing the cache file.
    """
    out = _cachefile('gobench', year, day)
    if os.path.exists(out):
        return False
    print(f'gobench: ({year}, {day}) -> {out}')
    cmd = [
        'go', 'test', '-run=$.^',
        f'-bench=BenchmarkAllDays/day={day:02d}',
        f'./{year:04d}/days',
    ]
    proc = subprocess.run(cmd, cwd='..', capture_output=True, check=True)
    with open(out, 'wb') as f:
        f.write(proc.stdout)
    return True


def _gobench_parse(year, day):
    """Parses a cached Go benchmark result file.

    Args:
        year: Contest year, 2015 to current.
        day: Contest day, 1 to 25 (unless in the future).

    Returns:
        A Pandas dataframe containing that day's data. TODO: describe format.
    """
    path = _cachefile('gobench', year, day)
    with open(path, 'r') as f:
        for line in f:
            m = re.search(r'BenchmarkAllDays/day=\S*\s+\d+\s+(\d+) ns/op', line)
            if m:
                runtime = float(m.group(1)) / 1000000000
                break
        else:
            raise RuntimeError(f'missing benchmark results for {year}, {day}')
    idx = pd.MultiIndex.from_product([[year], [day]], names=('year', 'day'))
    return pd.DataFrame({'runtime': runtime}, index=idx)


def _gobench_days():
    """Generates all the AoC puzzle days that have a Go solution present.

    Yields:
        A pair: the year, and a generator for the valid puzzle days on that year.
    """
    for year, days in contest_days():
        present = list(day for day in days if _gobench_present(year, day))
        if present:
            yield year, present


def _gobench_present(year, day):
    """Checks if a Go solution for the given year, day is present.

    Args:
        year: Contest year, 2015 to current.
        day: Contest day, 1 to 25 (unless in the future).

    Returns:
        True, if a Go source file with the expected name exists. False, otherwise.
    """
    d = f'day{day:02d}'
    path = f'../{year:4d}/{d}/{d}.go'
    return os.path.exists(path)


def _cachefile(dataset, year, day):
    """Returns the cache file name for a given day of data."""
    return f'cache/{dataset}.{year:04}.{day:02}.html'


def contest_days():
    """Generates all the valid AoC puzzle days.

    A new puzzle day is returned 2 hours after US/Eastern midnight, to give some time for the
    leaderboards to populate.

    Yields:
        A pair: the year, and generator for the valid puzzle days on that year.
    """
    last_date = datetime.now(tz=pytz.timezone('US/Eastern')) - timedelta(hours=2)
    for year in range(_FIRST_YEAR, last_date.year+1):
        last_day = 25
        if year == last_date.year and last_date.month == 12:
            last_day = last_date.day
        yield year, range(1, last_day+1)
