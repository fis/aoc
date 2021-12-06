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

## [Day 4](https://adventofcode.com/2021/day/4): Giant Squid

Nothing particularly clever about the Go solution; it just goes through the
motions. It's probably reasonably efficient, not that it really matters for the
puzzle input.

### Burlesque

These are particularly bad, but I just can't be bothered, as long as they work.

Part 1:

```
{#a{{sm}ay}}hd',' r~psJ100.-25co{5coJtp.+}m[hd
-.{[-J-]#ajbc{{{}r~}j+]m[}Z]`a}{!bayn!}w!-]!bfeFL++.*2./
```

Part 2:

```
{}hd',' r~psJ100.-25co{5coJtp.+}m[hd
-.{[-J-]#ajbc{{{}r~}j+]m[}Z]J`a{{sm}ay}fI#b\\#b.+`b}{#bL[100.<}w!
-]#a#b-]!!FL++.*2./
```

## [Day 5](https://adventofcode.com/2021/day/5): Hydrothermal Venture

Is it a coincidence that "Hydrothermal Venture" and "Horizontal & Vertical" have
the same initial letters? Probably.

Just out of curiosity, I tried out whether doing a "pairwise" overlap test for
all the lines outperforms the simpler method of just walking every line. In
theory, if the lines are long enough but mostly just intersect at one point, it
should. For the puzzle input, though, it doesn't. Here's a representative
benchmark:

    BenchmarkOverlaps/arrayHV-16       3240      330777 ns/op     966659 B/op       1 allocs/op
    BenchmarkOverlaps/countingHV-16      86    13436727 ns/op    8209626 B/op    3875 allocs/op
    BenchmarkOverlaps/pairwiseHV-16     843     1488492 ns/op     348414 B/op     164 allocs/op

### Burlesque

Part 1:

```
ln{"[0-9]+"~?ri2co}m[{tp{sm}ay}f[{J#r?-J++abj)sn{J?+}[[jE!}m^sg{[-nz}fl
```

Part 2:

```
ln{"[0-9]+"~?ri2cop^J#r?-J)ab>]j)sn{J?+}[[jE!}m[sg{[-nz}fl
```

## [Day 6](https://adventofcode.com/2021/day/6): Lanternfish

The trick of the day is to ignore the way the puzzle description is leading you
towards simulating the school of fish as individuals, and just realize it's
sufficient to simply track the number of fishes `f_c` that have a specific
internal counter value `c`. This way, the counts of day `t+1` can be derived
from the counts of day `t` as:

<!--math:day06
\vspace*{-3ex}
\begin{align*}
f_c^{(t+1)} &= f_{c+1}^{(t)} \quad\textrm{for $c \in \{0, 1, 2, 3, 4, 5, 7\}$} \\
f_6^{(t+1)} &= f_7^{(t)} + f_0^{(t)} \\
f_8^{(t+1)} &= f_0^{(t)}
\end{align*}
-->
![day06.png](math/2021-notes-day06.png)

The first equation represents how the fish with non-zero counters will just have
their counters uniformly decremented by one. There are also two special cases:
fishes with counter 6 will include both those fish that decremented their
counter from 7, as well as all the fish who cycled from 0; and fishes with
counter 8 will be only the newly spawned ones.

The Go solution does one more (entirely unnecessary) optimization: by using an
offset value to track which field of a circular array represents count 0, the
cyclic update can be done by incrementing the offset, and the only other action
needed is to increment the number of the (new) counter 6 by that of the (old)
counter 0.

Finally, there's one more way of looking at the problem. If we combine all the
counts from above into a single 9-element column vector `f^(t)`, we can boil
down the daily update into a single matrix multiplication, and therefore use
matrix exponentiation to directly get the counts for any given day:

<!--math:day06-mat
\vspace*{-3ex}
\begin{align*}
\mathbf{f}^{(t)} &= \begin{pmatrix}
0 & 1 & 0 & 0 & 0 & 0 & 0 & 0 & 0 \\
0 & 0 & 1 & 0 & 0 & 0 & 0 & 0 & 0 \\
0 & 0 & 0 & 1 & 0 & 0 & 0 & 0 & 0 \\
0 & 0 & 0 & 0 & 1 & 0 & 0 & 0 & 0 \\
0 & 0 & 0 & 0 & 0 & 1 & 0 & 0 & 0 \\
0 & 0 & 0 & 0 & 0 & 0 & 1 & 0 & 0 \\
1 & 0 & 0 & 0 & 0 & 0 & 0 & 1 & 0 \\
0 & 0 & 0 & 0 & 0 & 0 & 0 & 0 & 1 \\
1 & 0 & 0 & 0 & 0 & 0 & 0 & 0 & 0
\end{pmatrix}^t \mathbf{f}^{(0)}
\end{align*}
-->
![day06-mat.png](math/2021-notes-day06-mat.png)

This could be used to calculate the result for day `t` in less than `O(t)` time.

As a demonstration, here is Octave calculating the result of the part 1 example,
where we start with initial fish `3,4,3,1,2`, which represented in the vector
form is `[0 1 1 2 1 0 0 0 0]` (one `1`, one `2`, two `3`s, one `4`):

```
octave:1> sum((diag(ones(8,1), 1) + [[0 0 0 0 0 0 1 0 1]' zeros(9,8)])^80 * [0 1 1 2 1 0 0 0 0]')
ans = 5934
```

### Burlesque

Part 1:

This one does simulate each fish indepedently, which is slow, but okay for 80
days and takes less commands.

```
',;;ri{{Jz?{6.+9}if-.}m[}80E!L[
```

Part 2:

This one does the count-by-counter-value thing.

```
',;;ribc8rz{CN}Z]{RTJ[~jJ6!!x/.+6sa}256E!++
```
