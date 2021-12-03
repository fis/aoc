# AoC 2021 solution notes

## Background

Another year, another Advent of Code. I started doing these "seriously" in 2019,
so see [2020-notes.md](./2020-notes.md) and [2019-notes.md](./2019-notes.md) for
more of the background. The *tl;dr* is that I'm (a) writing in Go with unit
tests and all, and (b) writing these notes as a sort of a diary.

Last year, I solved some of the problems with Befunge-98. This time, I'm doing
the same with [Burlesque](https://esolangs.org/wiki/Burlesque). I'm not very
good at it.

## [Day 1](https://adventofcode.com/2021/day/1): Sonar Sweep

As usual, there isn't much to say about the first day. Although it's maybe worth
noting that you don't have to explicitly sum up the contents of each overlapping
window, because (gratuitous use of math):

<!--math:day01
\vspace*{-3ex}
\begin{align*}
\sum_{i=j-N}^{j-1} s_i &< \sum_{i=j-N+1}^j s_i \\
s_{j-N} + \sum_{i=j-N+1}^{j-1} s_i &< \sum_{i=j-N+1}^{j-1} s_i + s_j \\
s_{j-N} &< s_j
\end{align*}
-->
![day01.png](math/2021-notes-day01.png)

### Burlesque

Part 1:

```
psJ[-{.<}Z]++
```

Part 2:

```
psJ3.-{.<}Z]++
```

## [Day 2](https://adventofcode.com/2021/day/2): Dive!

No comment. Other than apparently puzzle titles this year won't have the same
alliterative appeal they did last year.

### Burlesque

Part 1:

```
ln<>{-]'f==}gB{{wd[~ri}ms}MPx/.-.*
```

Part 2:

```
0Ppln{wd^prijL[J7=={vvJpP.*CL}j{3.-.*PP.+Pp}jie}m[tp{++}mp
```

## [Day 3](https://adventofcode.com/2021/day/3): Binary Diagnostic

Continuing with straight-forward tasks. The solution here is more "bitwise" than
is really reasonable (especially the `keep := (2 * ones / N) ^ keepLCB` thing),
but sometimes you just have to have a little fun with it.

### Burlesque

Part 1:

```
ln)XXtpJ)n!j)fcCL{\[b2}mp
```

Part 2 (quite terrible):

```
ln)XXtpbc{n!fc}{{JPPJ+.Pp!!JJ.+'1+]{==}j+]fI{si}j+]m[}j11ia0Pp{-~nz}w!}Z]{FL\[b2}mp
```
