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

Edited to add: due to some IRC conversations, also added a non-regex solution.
It's definitely faster:

```
BenchmarkAlgos/regexp-32      45457    774447 ns/op
BenchmarkAlgos/string-32    3362403     11299 ns/op
```

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

## [Day 4](https://adventofcode.com/2024/day/4): Ceres Search

Nothing to say here, haven't even done the Burlesque part yet. If it wasn't for
the diagonals it wouldn't be too bad; now it's going to be a pain. For part 2,
maybe I can do something like shifting and zipping.

## [Day 5](https://adventofcode.com/2024/day/5): Print Queue

Kinda frustrated by today's puzzle. Started by assuming that of course the
ruleset will not explicitly contain rules for *all* page pairs, and adapted the
existing graph library to deal with integer rather than string labels, so that I
could use the topological sort routine in it. Of course it turns out that:

- The full list of rules is *not* a DAG; it's just that there are no cycles
  within the pages of an individual update.
- The ruleset does in fact contain a rule for every possible page pair.

In the end it all boiled down to just checking the consecutive pages in part 1,
and a call to `slices.SortFunc` in part 2. Oh well, at least it's pretty fast.

### Burlesque

Burlesque has a slightly annoying quirk in that both the state stack as well as
the "global variables" (`s0`, `g0`) are inaccessible in the comparison function
of `sb`. So this code has to smuggle the list of rules into the function block
itself.

Part 1:

```
ln{""};;p^{"|";;)ri}m[s0{",";;)riJ2CO{g0j~[}aljsa2./!!.*}ms
```

Part 2:

```
ln{""};;p^{"|";;)ri}m[s0{",";;)riJ2CO{g0j~[}aln!j{_+j~[2.*-.}g01iasbsa2./!!.*}ms
```

Combined:

```
C: ln{""};;p^{"|";;)ri}m[s0{",";;)riJ2CO{g0j~[}al  j                   sa2./!!.*}ms
2:                                               n! {_+j~[2.*-.}g01iasb
```
