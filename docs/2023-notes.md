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
