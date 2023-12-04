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

Burlesque is still missing at this point. Not looking forward to it: "2D"
problems are one of the worst kind to do in it.
