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
import os.path
import pandas as pd
import pytz
import requests
from datetime import datetime, timedelta
from lxml import html


_FIRST_YEAR = 2015
_LEADERBOARD_FILE = 'cache/leaderboard.pickle'


def leaderboard():
    """Loads the daily leaderboard data frame.

    The returned frame will have a three-level hierarchical index, consisting of the integer columns
    `year`, `day` and `rank` with the obvious semantics. It will have two data series, `one_star`
    and `two_stars`, representing the time (in seconds) it took to get that specific leaderboard
    position for the first part or the full puzzle, respectively.
    """
    return pd.read_pickle(_LEADERBOARD_FILE)


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
    out = _leaderboard_cachefile(year, day)
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
        A Pandas dataframe containing that day's data. It will have one index, `rank`, with values
        ranging from 1 to 100, and two series, `one_star` and `two_stars`, containing the time (in
        seconds) that yielded that given ranking, for the first part or both parts respectively.
    """
    page = html.parse(_leaderboard_cachefile(year, day))
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


def _leaderboard_cachefile(year, day):
    """Returns the cache file name for the daily leaderboard."""
    return f'cache/leaderboard.{year:04}.{day:02}.html'


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
