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


import altair as alt
import aocdata
import numpy as np
import pandas as pd


def plot_leaderboard():
    """Redraws all the plots based on the daily leaderboard data."""

    data = aocdata.leaderboard()
    plot_leaderboard_time(data, 'one_star', 'out/time.one.html')
    plot_leaderboard_time(data, 'two_stars', 'out/time.two.html')
    plot_leaderboard_twist(data)
    plot_leaderboard_dist(data)


def plot_leaderboard_time(data, series, file_name):
    """Plots the per-day time distribution for the leaderboard rankings.

    Args:
        data: Leaderboard data frame.
        series: Which series to use; 'one_star' or 'two_stars'.
        file_name: Output file.
    """

    print(f'leaderboard_time:{series}')

    quantiles = data.loc[(slice(None), slice(None), [1, 25, 50, 75, 100])].unstack()
    points = (quantiles[series] / 60).rename(columns=lambda r: f'r{r}').reset_index()

    y_title = f'Time to get {series.replace("_", " ")} (min)'
    y_scale = alt.Scale(type='log')

    base = alt.Chart(points).encode(x='year:O', color='year:N')
    rule = base.mark_rule().encode(alt.Y('r1:Q', title=y_title, scale=y_scale), alt.Y2('r100:Q'))
    bar = base.mark_bar().encode(alt.Y('r25:Q'), alt.Y2('r75:Q'))
    faceted = (rule + bar).facet(column=alt.Column('day:O', title='Day of contest'))
    faceted.configure_scale(bandPaddingInner=0.4).save(file_name)


def plot_leaderboard_twist(data):
    """Plots the per-day twistiness ranking, both as a bar chart and as a heatmap."""

    print('leaderboard_twist')

    totals = data.groupby(level=('year', 'day')).sum()
    twist = pd.DataFrame({'twist': totals['two_stars'] / totals['one_star']}).reset_index()

    alt.Chart(twist) \
        .encode(
            x='year:O',
            y=alt.Y('twist:Q', title='Twistiness'),
            color='year:N',
            column=alt.Column('day:O', title='Day of contest')) \
        .mark_bar() \
        .configure_scale(bandPaddingInner=0.2) \
        .save('out/twist.html')

    color_scale = alt.Scale(scheme='yelloworangered', type='log')
    alt.Chart(twist) \
        .encode(
            x=alt.X('day:O', title='Day of contest'),
            y='year:O',
            color=alt.Color('twist:Q', title='Twistiness (log)', scale=color_scale)) \
        .mark_bar() \
        .configure_scale(bandPaddingInner=0.1) \
        .save('out/twist.heat.html')


def plot_leaderboard_dist(data):
    """Plots the distribution of leaderboard entires over time."""

    print('leaderboard_dist')

    data = (data/60).rename(columns={'one_star': 'one star', 'two_stars': 'two stars'}).stack()
    data = data.rename_axis(index=['year','day','rank','stars'])
    data = pd.DataFrame({'time_min': data}).reset_index()

    x_scale = alt.Scale(type='log', base=2, domain=[0.25, 256])
    color_scale = alt.Scale(scheme='redpurple')
    alt.Chart(data) \
        .transform_density(
            'time_min',
            groupby=['year','day', 'stars'],
            extent=[0.25, 256],
            steps=2000) \
        .mark_line() \
        .encode(
            alt.X('value:Q', title='Time to solution (min)', scale=x_scale),
            alt.Y('density:Q'),
            alt.Color('day:O', scale=color_scale)) \
        .properties(width=600, height=250) \
        .facet(column='stars:N', row='year:O') \
        .resolve_scale(x='independent', y='independent') \
        .save('out/time.dist.html')


def plot_stats():
    """Redraws all the plots based on the stats data."""

    data = aocdata.stats()

    # Altair has an annoying tendency to localize times, but for the contest it makes sense to show
    # them as US/Eastern no matter what. It has an even more annoying tendency to, in some browsers,
    # even render naive (TZ-unaware) times as that. This terrible thing converts the timestamps into
    # allegedly-UTC timestamps that still show the US/Eastern calendar times. :/
    utc_times = pd.to_datetime(data.ts*1000000000).dt.tz_localize('UTC')
    naive_times = utc_times.dt.tz_convert('US/Eastern').dt.tz_localize(None)
    data['ts_utc'] = naive_times.dt.tz_localize('UTC')

    plot_stats_chart(data)
    plot_stats_trajectory(data)
    plot_stats_ratio(data)


def plot_stats_chart(data):
    """Plots the number of solutions over time."""

    print('stats_chart')

    counts = data.reset_index().drop(columns=['sidx', 'ts'])
    counts = counts.set_index(['year', 'day', 'ts_utc', 'since'])
    counts = counts.rename(columns={'one_star': 'just one star', 'two_stars': 'two stars'})
    counts = counts.stack().rename_axis(index=['year', 'day', 'ts_utc', 'since', 'stars'])
    counts = pd.DataFrame.from_dict({'count': counts}).reset_index()

    selection = alt.selection_multi(fields=['day'], bind='legend')
    hover = alt.selection_single(fields=['day'], on='mouseover')

    alt.Chart(counts) \
        .mark_line() \
        .encode(
            alt.X('utcyearmonthdatehoursminutes(ts_utc):T', title='Time of stats snapshot'),
            alt.Y('count:Q', title='Number of solutions'),
            alt.Color('day:O', scale=_rainbow_day_scale),
            opacity=alt.condition(selection | hover, alt.value(1), alt.value(0.2)),
            tooltip=['day']) \
        .add_selection(selection) \
        .add_selection(hover) \
        .facet(column='stars:N', row='year:O') \
        .resolve_scale(x='independent', y='independent') \
        .save('out/stats.chart.html')

    counts['since_d'] = counts.since / 86400

    alt.Chart(counts) \
        .mark_line() \
        .encode(
            alt.X('since_d:Q', title='Time since puzzle start (days)', scale=alt.Scale(type='sqrt')),
            alt.Y('count:Q', title='Number of solutions'),
            alt.Color('day:O', scale=_rainbow_day_scale),
            opacity=alt.condition(selection | hover, alt.value(1), alt.value(0.2)),
            tooltip=['day']) \
        .add_selection(selection) \
        .add_selection(hover) \
        .facet(column='stars:N', row='year:O') \
        .resolve_scale(x='independent', y='independent') \
        .save('out/stats.aligned.chart.html')


def plot_stats_trajectory(data):
    """Plots an X-Y line of solutions with just one star vs. solutions with both stars."""

    print('stats_trajectory')

    data = data.reset_index()
    data = data.loc[(data['one_star'] > 0) & (data['two_stars'] > 0)]

    selection = alt.selection_multi(fields=['day'], bind='legend')
    hover = alt.selection_single(fields=['day'], on='mouseover')

    for scale_type in ('linear', 'log'):
        alt.Chart(data) \
            .mark_line() \
            .encode(
                alt.X('one_star:Q', title='Solutions with just one star', scale=alt.Scale(type=scale_type)),
                alt.Y('two_stars:Q', title='Solutions with both stars', scale=alt.Scale(type=scale_type)),
                alt.Color('day:O', scale=_rainbow_day_scale),
                opacity=alt.condition(selection | hover, alt.value(1), alt.value(0.2)),
                tooltip=['day'],
                order='ts') \
            .add_selection(selection) \
            .add_selection(hover) \
            .properties(width=800, height=480) \
            .facet(row='year:O') \
            .resolve_scale(x='independent', y='independent') \
            .save('out/stats.trajectory.html' if scale_type == 'linear' else 'out/stats.trajectory.log.html')


def plot_stats_ratio(data):
    """Plots the ratio of solutions with one star to all solutions."""

    print('stats_ratio')

    ratios = pd.DataFrame.from_dict({'ts_utc': data.ts_utc, 'since': data.since, 'ratio': data.one_star / (data.one_star + data.two_stars)})
    ratios = ratios.reset_index()
    ratios = ratios.loc[~ratios['ratio'].isnull()]

    selection = alt.selection_multi(fields=['day'], bind='legend')
    hover = alt.selection_single(fields=['day'], on='mouseover')

    alt.Chart(ratios) \
        .mark_line() \
        .encode(
            alt.X('utcyearmonthdatehoursminutes(ts_utc):T', title='Time of stats snapshot'),
            alt.Y('ratio:Q', title='Fraction of solutions with just one star', scale=alt.Scale(type='log')),
            alt.Color('day:O', scale=_rainbow_day_scale),
            opacity=alt.condition(selection | hover, alt.value(1), alt.value(0.2)),
            tooltip=['day']) \
        .add_selection(selection) \
        .add_selection(hover) \
        .facet(row='year:O') \
        .resolve_scale(x='independent', y='independent') \
        .save('out/stats.ratio.html')

    ratios['since_d'] = ratios.since / 86400

    alt.Chart(ratios) \
        .mark_line() \
        .encode(
            alt.X('since_d:Q', title='Time since puzzle start (days)', scale=alt.Scale(type='sqrt')),
            alt.Y('ratio:Q', title='Fraction of solutions with just one star', scale=alt.Scale(type='log')),
            alt.Color('day:O', scale=_rainbow_day_scale),
            opacity=alt.condition(selection | hover, alt.value(1), alt.value(0.2)),
            tooltip=['day']) \
        .add_selection(selection) \
        .add_selection(hover) \
        .facet(row='year:O') \
        .resolve_scale(x='independent', y='independent') \
        .save('out/stats.aligned.ratio.html')


def plot_gobench():
    """Redraws all the plots based on the Go benchmark data."""

    data = aocdata.gobench()
    plot_gobench_time(data)


def plot_gobench_time(data):
    """Plots the times taken by all the solutions."""

    print('gobench_time')

    data = data.reset_index()
    data['yearday'] = data.apply(lambda row: f'{int(row.year)}-{int(row.day)}', axis=1).astype('string')

    alt.Chart(data) \
        .encode(
            x=alt.X('runtime:Q', title='Time to solve puzzle (s)'),
            y=alt.Y('yearday:O', sort='-x'),
            color='year:N') \
        .mark_bar() \
        .configure_scale(bandPaddingInner=0.2) \
        .save('out/gobench.time.html')


_rainbow_day_domain = list(range(1, 26))
_rainbow_day_range = [
    '#ff0000','#eeac06','#ceee0c','#13ff32','#19d1d1',
    '#1f6aff','#d326d3','#ff2c2c','#f1bb33','#d7f139',
    '#3fff59','#46dada','#4c88ff','#dc52dc','#ff5959',
    '#f4c95f','#e0f466','#6cff7f','#72e2e2','#79a5ff',
    '#e57fe5','#ff8585','#f7d88c','#e9f792','#99ffa6',
]
_rainbow_day_scale = alt.Scale(domain=_rainbow_day_domain, range=_rainbow_day_range)
