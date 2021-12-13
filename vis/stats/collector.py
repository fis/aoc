#! /usr/bin/env python3

import pytz
import requests
import time
from datetime import datetime, timedelta
from lxml import html


def main():
    while True:
        now, next = next_time()
        print(f'waiting for {next.isoformat()} for next update')
        time.sleep((next - now).total_seconds())

        try:
            fetch_start = datetime.now()
            content = fetch(next.year)
            fetch_end = datetime.now()

            counts = parse(content)

            with open(f'stats.{next.year}.txt', 'a') as f:
                columns = [next.timestamp(), (fetch_start.timestamp() + fetch_end.timestamp()) / 2]
                columns.extend(counts['first'])
                columns.extend(counts['both'])
                print(' '.join(str(c) for c in columns), file=f)

        except Exception as e:
            print(f'update failed: {e}')


def next_time():
    aoctz = pytz.timezone('US/Eastern')
    now = datetime.now(tz=aoctz)
    next = now.replace(minute=now.minute//5*5, second=0, microsecond=0)
    while next < now:
        next += timedelta(seconds=5*60)
    if next.month < 12:
        next = aoctz.localize(datetime(now.year, 12, 1))
    return now, next


def fetch(year):
    url = f'https://adventofcode.com/{year}/stats'
    r = requests.get(url, timeout=30)
    r.raise_for_status()
    return r.content


def parse(content):
    page = html.fromstring(content)
    entries_first = page.xpath('//pre[@class="stats"]//span[@class="stats-firstonly"]')
    entries_both = page.xpath('//pre[@class="stats"]//span[@class="stats-both"]')
    if len(entries_first) != 50 or len(entries_both) != 50:
        raise RuntimeError(f'unexpected number of spans: {len(entries_first)}, {len(entries_both)}')
    counts_first = [int(e.text_content()) for e in entries_first[48::-2]]
    counts_both = [int(e.text_content()) for e in entries_both[48::-2]]
    return dict(first=counts_first, both=counts_both)


if __name__ == '__main__':
    main()
