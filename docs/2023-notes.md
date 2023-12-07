# AoC 2023 solution notes

## Background

If you've ended up reading this, chances are you already know what this is all
about. If not, see the background sections of the various previous years.

Without further ado...

## [Day 1](https://adventofcode.com/2023/day/1): Trebuchet?!

As usual, there isn't much to say about the first day. Decided to scan from both
ends of the string separately. To handle the words, rather than coming up with a
custom string set search routine, just looking them up individually within the
span of string still remaining.

Feeling very rusty at Burlesque, which is also typical.

### Burlesque

Part 1:

```
ln{:><J-]j[~.+ri}ms
```

Part 2:

```
ln{iS{{j~!}j+]"1one2two3three4four5five6six7seven8eight9nine"{><}gBjfi
2./+.}m[:nzJ-]j[~_+10ug}ms
```

This is based on finding all the suffix strings of the line (`iS`), and for each
of them testing whether any of the digits (numeric or written) is a prefix.
Alternating the numbers and words allows using `{><}gB` to expand them into a
block, giving a semi-compact encoding.

## [Day 2](https://adventofcode.com/2023/day/2): Cube Conundrum

Most of the job is in parsing the input; actual tasks are much simpler.

### Burlesque

Part 1:

```
ln{": ";;p^:><rij" ";;2co{-]**3.%jri_+}^m><{-]}gB{)[~>]}m[2rz12?+{<=}Z]r&.*}ms
```

Part 2:

```
ln{": ";;[~" ";;2co{-]**3.%jri_+}^m><{-]}gB{)[~>]}m[pd}ms
```

Combined:

```
1:          p^:><rij                                         2rz12?+{<=}Z]r&.*
C: ln{": ";;        " ";;2co{-]**3.%jri_+}^m><{-]}gB{)[~>]}m[                 }ms
2:          [~                                               pd
```

## [Day 3](https://adventofcode.com/2023/day/3): Gear Ratios

Nothing much to say about the Go solution, except that it prompted me to add a
feature to the `util.Level` type so that it keeps a fixed-side region densely
allocated. This makes it perform a lot better for tasks where the area of
interest is predictable and densely packed, without having to give up the
conveniences (`LevelSolver`, nicer API).

For a lot of past problems, had to switch to more rudimentary types like naked
`[][]byte`s and such to get acceptable performance. Benchmarking some of the old
days that still used a `Level`, some of them got pretty nice (up to ~1 order of
magnitude) speedups, and nothing really got slower.

For the Z80 solution, went with a scheme where cell `(x, y)` of the input is
loaded into RAM at address `0x2000 | y<<8 | x`. For the 140-line puzzle input,
this consumes over half the device RAM, but makes it really easy to do the 2D
accessing needed for the task, by just manipulating the high/low halves of the
register pairs.

The Burlesque solutions were added later, and were sufficiently tedious to avoid
spending any time optimizing for them.

### Burlesque

The general idea is this:

- Turn the `abc123def` strings into `{'a 'b 'c 123 123 123 'd 'e 'f}` blocks. In
  other words, each number appears as a whole number for all the locations that
  its digits can be found in.
- Find the 2D `{row col}` indices for all symbols of interest; anything that's
  not a number or a `.` for part 1, or just the `*` symbols for part 2.
- For each index, expand it to its 3x3 neighborhood.
- Replace the index values with the corresponding contents of the schematic.
- Add a separator character between each row, and then remove consecutive
  duplicate elements. This gets rid of the repetition when the same number
  occupies more than one cell adjacent to a symbol.
- Filter away all the non-number contents of the groups.
- For part 2, also filter away all groups that don't have exactly 2 numbers.
- Compute the final values: overall sum for part 1, sum of products for part 2.

Part 1:

```
ln{XX{><j><&&}gb{J-]><{\[sa.*)ri}if}\m}m[JPp{{Jto-]'C==j'.!=&&}fI}m[zi
{{j_+}j+]m[}m^\[{J?dj?i{r@}Z]^pcp{pPjd!}m[3co'xIC=[)-]{to-]'I==}f[++}ms
```

Part 2:

```
ln{XX{><j><&&}gb{J-]><{\[sa.*)ri}if}\m}m[JPp{{'*==}fI}m[zi
{{j_+}j+]m[}m^\[{J?dj?i{r@}Z]^pcp{pPjd!}m[3co'xIC=[)-]{to-]'I==}f[}m[{L[2==}f[)pd++
```

Combined:

```
1:                                               Jto-]'C==j'.!=&&
C: ln{XX{><j><&&}gb{J-]><{\[sa.*)ri}if}\m}m[JPp{{                }fI}m[zi
2:                                               '*==

1:                                                                   ++}ms
C: {{j_+}j+]m[}m^\[{J?dj?i{r@}Z]^pcp{pPjd!}m[3co'xIC=[)-]{to-]'I==}f[
2:                                                                   }m[{L[2==}f[)pd++
```

## [Day 4](https://adventofcode.com/2023/day/4): Scratchcards

The Go solution this time has two versions of the code for parsing the cards: a
"simple" one written first that just relies on `strings.Split` and
`strconv.ParseInt`, and a "fast" one that assumes ASCII, at most two-digit
numbers, and so on.

The motivation for adding the second one came from running the Go CPU profiler
on the day benchmark. You can see the absolutely ludicrous results in the file
[2023-day14-prof.png](2023-day14-prof.png). According to synthetic benchmarks
for just the parser part, the other one is roughly 10x faster.

The set intersection builtin `IN` is the MVP of the Burlesque solution. Parsing
is again a large chunk of the total work. The part 2 solution does a reduce step
where the accumulator contains the remaining card counts, and totals are added
to the state stack as each card is processed.

The Z80 solution again uses an addressing trick, where each winning number `x`
is marked by setting the byte at `0x8000 | x`, so that set membership testing
is trivial. This time that takes only about 160 bytes, though. The code also
reads the numbers as hex, since the exact values don't really matter.

### Burlesque

Part 1:

```
ln{":";;[~"|";;)t[psp^INL[2j**2./}ms
```

Part 2:

```
lnsaro)nz+]{":";;[~"|";;)t[psp^INL[jg_JPpbxx/.*\[bxj+]tp)++}r[p\CL++
```

Combined:

```
1:                                    2j**2./                  ms
C: ln         {":";;[~"|";;)t[psp^INL[                        }
2:   saro)nz+]                        jg_JPpbxx/.*\[bxj+]tp)++ r[p\CL++
```

## [Day 5](https://adventofcode.com/2023/day/5): If You Give A Seed A Fertilizer

It's a bit of a stretch, but I'm counting today as this year's first instance of
the traditional AoC pattern, where the solution for part 1 could in theory be
used to solve part 2, but doing so naïvely would be computationally infeasible.

The natural solution here is to treat the input as a set of intervals, and
transform them as such (splitting where necessary), as opposed to trying to
iterate over each possible range.

Doing part 1 in Burlesque was okay enough, but part 2 was a pain. Again, not
feeling like golfing these solutions: it's enough to have them here.

### Burlesque

The two parts don't have all that much in common, so omitting the combined
section. The logic for part 2 this time is:

- Convert the `{dst src len}` ranges into `{start end offset}` ones.
- For each "layer" of mappings, add in synthetic ranges (with `0` offset) to
  make sure that the ranges cover the entire input domain, with no gaps between
  them.
- To solve the problem, do a reduce operation where the current set of input
  ranges meets each set of mappings in turn. The operation:
  - Takes the cross product of input `{is ie}` ranges and `{ms me mo}` mappings.
  - Turns each pair into `{max(is,ms) min(ie,me) mo}` pseudorange.
  - Filters to keep only the `max(...) < min(...)` cases. Together with the
    previous step, this effectively finds all the non-empty intersections the
    input ranges have with the mappings.
  - Add the offset to all the surviving pseudoranges to make them the mapped
    inputs of the next layer.

Part 1:

```
ln{""};;g_-]WD[-)rij)[-psPp{pPj+]
{jJJ_+1[+x/j[+jJ{Jx/[-iT[-)++rm==}j+]x/jfe~]^p.-.+}r[}m[<]
```

Part 2:

```
ln{""};;g_-]WD[-)ri2co{iT[-)++}m[j)[-ps{{J[-iT[-)++j~]^p.-[+}m[}m[
1rz9e9?*3.*tp{_+><J2CO{p^[--]0_+j-][]}m[_+}j+]m[
j+]{cp{tp{>]<]++}z[\[e!Cl}m[{~]^p.<}{l_?+}FM}r[FL<]
```

## [Day 6](https://adventofcode.com/2023/day/6): Wait For It

Let's start with some notation for a single race. Call the time limit of the
race `T`, and the time you hold down the button `h`. Now we can compute the
distance `d(h)` you will travel based on the hold time:

<!--
d(h) = (T - h) * h
-->
![The equation for distance: d(h) = (T - h) * h.](math/2023-notes-day06-d.png)

Let's also use `d_best` to denote the previous best distance of the race. The
brute force solution to determine `w`, the number of different ways to win, is
to simply try all the possible choices and count how many of them win:

<!--
     T                       T
w = sum [ d(h) > d_best ] = sum [ (T - h) * h > d_best ]
    h=0                     h=0
-->
![A sum for counting the winning combinations.](math/2023-notes-day06-w.png)

Normally, when there's a problem where part 1 has a simple brute force solution
like this, running it on part 2 tends to take a long time indeed. But in this
instance, the the runtime is measured in milliseconds even for part 2. So we
could even stop here.

That's not to say we can't do any better, though. Let's take a closer look at
that inequality. We can turn it into a quadratic form:

<!--
(T - h) * h > d_best
-h^2 + T*h - d_best > 0
-->
![Quadratic inequality for the winning hold times.](math/2023-notes-day06-ineq.png)

Since it's sad to even entertain the possibility we could not win a race, we can
probably just assume the parameters are always such that the equation has real
solutions. If so, counting the possible ways to win is equivalent to counting
how many integers are in the open interval between the two solutions `h_min` and
`h_max`. The values of these pop out of the quadratic formula:

<!--
h_min = (T - sqrt(T^2 - 4*d_best)) / 2
h_max = (T + sqrt(T^2 - 4*d_best)) / 2
-->
![The two solutions for the quadratic equation.](math/2023-notes-day06-minmax.png)

Since we need to beat instead of merely matching the earlier record (that's why
it's an open interval), we'll need to not include `h_min` or `h_max` themselves,
should they happen to be integer values. Therefore the equation for `w` becomes:

<!--
w = (⌈h_max⌉ - 1) - (⌊h_min⌋ + 1) + 1
  = ⌈h_max⌉ - ⌊h_min⌋ - 1
-->
![An alternative equation for w based on the solutions.](math/2023-notes-day06-w2.png)

### Burlesque

For Burlesque, there's two solutions to part 2 this time. The first uses the
same brute force approach as part 1, where we calculate the distance for every
possible hold time and then count how many beat the record. This is actually
(barely) computationally feasible even for the full puzzle input, but takes
about 8 gigabytes of RAM and 4 minutes to complete. The second uses the
closed-form solution, but is significantly longer (in terms of code size).

Part 1:

```
ln{WD[-)ri}m[tp{rzJ<-{.*}Z]j?-{0.>}fl}m^pd
```

Part 2 (brute force):

```
ln{:><ri}MPjrzJ<-{.*}Z]j?-{0.>}fl
```

Combined:

```
1:    WD[-)ri m[tp{                     }m^pd
C: ln{       }     rzJ<-{.*}Z]j?-{0.>}fl
2:    :><ri   MPj
```

Part 2 (faster alternative):

```
ln{:><rd}MPjJJ.*x/4.*.-r@2rz?d:nz?*?+2?/^pclj?iav.-ti
```

<!--math

%: day06-d

\vspace*{-3ex}
\begin{align*}
d(h) &= (T - h) \cdot h
\end{align*}

%: day06-w

\vspace*{-3ex}
\begin{align*}
w &= \sum_{h=0}^T [ d(h) > d_{\mathrm{best}} ]
= \sum_{h=0}^T [ (T - h) \cdot h > d_{\mathrm{best}} ]
\end{align*}

%: day06-ineq

\vspace*{-3ex}
\begin{align*}
(T - h) \cdot h &> d_{\mathrm{best}} \\
-h^2 + T h - d_{\mathrm{best}} &> 0
\end{align*}

%: day06-minmax

\vspace*{-3ex}
\begin{align*}
h_{\mathrm{min}} &= \frac{1}{2} \left( T - \sqrt{T^2 - 4 d_{\mathrm{best}}} \right) \\
h_{\mathrm{max}} &= \frac{1}{2} \left( T + \sqrt{T^2 - 4 d_{\mathrm{best}}} \right)
\end{align*}

%: day06-w2

\vspace*{-3ex}
\begin{align*}
w &= \left(\lceil h_{\mathrm{max}} \rceil - 1\right) - \left(\lfloor h_{\mathrm{min}} \rfloor + 1\right) + 1 \\
&= \lceil h_{\mathrm{max}} \rceil - \lfloor h_{\mathrm{min}} \rfloor - 1
\end{align*}

-->
