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


def plot_gobench():
    """Redraws all the plots based on the Go benchmark data."""

    data = aocdata.gobench()
    plot_gobench_time(data)


def plot_gobench_time(data):
    """Plots the times taken by all the solutions."""

    print('gobench_time')

    data = data.reset_index()

    x_title = 'Puzzle day'
    y_title = 'Time to solve puzzle (s)'
    y_scale = alt.Scale(type='log')

    alt.Chart(data) \
        .mark_line() \
        .encode(
            alt.X('day:O', title=x_title),
            alt.Y('runtime', title=y_title, scale=y_scale),
            alt.Color('year:N')) \
        .properties(width=800, height=600) \
        .save('out/gobench.time.html')
