# Advent of Code data visualization

This directory contains code to scrape and parse the Advent of Code daily
leaderboards into a Pandas data frame, and some Altair plots that hope to put
contest puzzles into the historical context.

## Leaderboard time

The `time.one` and `time.two` plots are straightforward box plots of the time it
took the first 100 people to get the first or both stars, respectively. The
endpoints of the rule show the times used by the first and the 100th rank, and
the edges of the box show the 25th and the 75th rank.

## Twistiness

The twistiness of a puzzle can be measured by comparing how long it took the
general public to get both stars, compared to just the first one.

The `twist` plot shows this as a similar grouped bar chart as the leaderboard
time plots, while the `twist.heat` plot uses a heat map. The latter is better
for overall comparisons, especially across years. The former is better if you're
interested in the exact twistiness values.

The specific metric used here is based on the total time used by all 100
leaderboard ranks to get one or two stars; the metric is simply the ratio of the
latter and the former. This means it's never below one.
