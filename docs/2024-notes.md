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
ps2cotp)><^p{.-ab}Z]++
```

Part 2:

```
ps2cotp^p{jJ#rCN.*}j+]ms
```

Combined:

```
1:        )><   .-ab     Z]++
C: ps2cotp   ^p{        }
2:              jJ#rCN.* j+]ms
```

## [Day 2](https://adventofcode.com/2024/day/2): Red-Nosed Reports

Tried to do an "allow one error before returning false" style solution for part
2, but it failed a case like `[10 9 11 12]`, where the deletion would have had
to change an already established trajectory. In the end just making copies.

### Burlesque

Part 1:

```
ln{ps2CO{.-}^mJ)snsmj)abJ>]3<=j<]1>=&&&&}ms
```

Part 2:

```
ln{pssa-.rzj{jRA}j+]m[{2CO{.-}^mJ)snsmj)abJ>]3<=j<]1>=&&&&}ay}ms
```

Combined:

```
1:                                                            ms
C: ln{ps                  2CO{.-}^mJ)snsmj)abJ>]3<=j<]1>=&&&&}
2:      sa-.rzj{jRA}j+]m[{                                    ay}ms
```

## [Day 3](https://adventofcode.com/2024/day/3): Mull It Over

Went "where the fence is lowest", and just did the regex.

### Burlesque

If there's one thing Burlesque is not great for, it's anything the slightest bit
imperative. Would have probably done it with a regex-replace instead, but it
seems like the relevant regex library does not do non-greedy matching.

Part 1:

```
"mul\([0-9]+,[0-9]+\)"~?{",";;{:><ri}mp}ms
```

Part 2:

```
"mul\([0-9]+,[0-9]+\)|do(n't)?\(\)"~?0+]1Pp{J1!!'o=={2!!'n!=Pp}j{",";;{:><ri}mppP.*.+}jie}r[
```

Combined:

```
1:                                      {                                         }ms
C: "mul\([0-9]+,[0-9]+\)             "~?                            ",";;{:><ri}mp
2:                      |do(n't)?\(\)   0+]1Pp{J1!!'o=={2!!'n!=Pp}j{              pP.*.+}jie}r[
```
