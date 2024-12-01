# AoC 2024 solution notes

## Background

It's that time of the year again. See previous years for the context.

## [Day 1](https://adventofcode.com/2024/day/1): Historian Hysteria

There's never anything particularly insightful to say about the first few days.
For the Go solution, in part 2, the "natural" approach would probably be to turn
the right list into `map[int]int` frequency counts and then do one pass over the
left list, but since I've got them sorted for part 1 anyway, I've opted to do a
merge-sort-merge-steppish kind of a thing instead.

### Burlesque

By December of year X, I've always completely forgotten all the little tips and
tricks of how to do Burlesque efficiently I learned in year X-1.

Part 1:

```
WD)ri2cotp)><^p{.-ab}Z]++
```

Part 2:

```
WD)ri2cotp^p{jJ#rCN.*}j+]ms
```

Combined:

```
1:           )><   .-ab     Z]++
C: WD)ri2cotp   ^p{        }
2:                 jJ#rCN.* j+]ms
```
