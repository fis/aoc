# AoC 2025 solution notes

## Background

Sounds like we're getting less AoC this year. Well, it is what it is.

## [Day 1](https://adventofcode.com/2025/day/1): Secret Entrance

There's never much to say about day 1. The solution (for part 2) just decomposes
every move into the number of full rotations (one click at 0 per full turn) and
the remainder.

### Burlesque

By December of year X, I've always completely forgotten all the little tips and
tricks of how to do Burlesque efficiently I learned in year X-1. (And this
comment I can just copy over from a past year, because it happens every time.)

I've always thought it's odd Burlesque doesn't have a "scan" builtin (like
fold/reduce, except keeping all the intermediate elements). These solutions use
the global state stack as an alternative. The `**2&&-.` sequence is a trick to
go from `'L` and `'R` to `-1` and `1`, respectively.

Part 1:

```
ln50Pp{g_**2&&-.jri.*pP.+100.%Pp}m[p\CL:z?L[
```

Part 2:

```
ln50Pp{g_**2&&-.jri.*pP.+Pp}m[p\CLJ2CO{><p^+.j-.r@}m[+]FL{100dv}fl
```

Combined:

```
1:                          100.%         :z?L[
C: ln50Pp{g_**2&&-.jri.*pP.+     Pp}m[p\CL
2:                                        J2CO{><p^+.j-.r@}m[+]FL{100dv}fl
```

## [Day 2](https://adventofcode.com/2025/day/2): Gift Shop

For day 2, there's two solutions here. The simple(r) one just iterates over the
range, testing each number in it for whether it matches one of the patterns.

The more subtle one uses the following approach:

- For every range, split it into intervals where all the numbers are of the
  same length. For example, the range 123-45678 would be split into subranges
  123-999, 1000-9999, and 10000-45678, respectively.
- For every subrange, iterate over the possible "base" lengths that could
  potentially form a repeated pattern in a number of that length. For example,
  for 6-digit numbers, these are 1 (pattern `aaaaaa`), 2 (pattern `ababab`) and
  3 (pattern `abcabc`).
- For a given subrange and base length $n$, take the first $n$ digits of the
  start and end values, and iterate over those values. For example, for the
  range 1234-5678 and $n=2$, this would mean iterating from 12 to 56 inclusive.
- For each such base value, repeat it enough times to construct a number of the
  correct length; in the above example, we would test the values 1212, 1313,
  and so on through to 5656 inclusive.
- For each of those candidate numbers, add them to a map if they're contained
  within the range. The map is used to avoid overcounting: otherwise the value
  12121212 would be counted both when checking length 2 (12, 4 times) and length
  4 (1212, 2 times).

### Burlesque

Neither of the first two days have been superbly Burlesque-friendly.

Part 1:

```
',;;{ps)ab^pr@{XXsa+.2./cosa1.>jsm&&}f[++}ms
```

Part 2:

```
',;;{ps)ab^pr@{XXsafCNBj{sax/./cosa1.>jsm&&}j+]ay}f[++}ms
```

Combined:

```
1:                    +.2
C: ',;;{ps)ab^pr@{XXsa          ./cosa1.>jsm&&}      f[++}ms
2:                    fCNBj{sax/               j+]ay}
```

## [Day 3](https://adventofcode.com/2025/day/3): Lobby

### Burlesque

Part 1:

```
ln{JJ~]>]Fi.-g_j>]_+ri}ms
```

Part 2:

```
ln{{JJ<-g0-.S0.->]Fi.-g_j}12S0E!vvCl\[ri}ms
```

Combined:

```
1:       ~]                 >]_+
C: ln{ JJ          >]Fi.-g_j             ri}ms
2:    {  <-g0-.S0.-         }12S0E!vvCl\[
```

## [Day 4](https://adventofcode.com/2025/day/4): Printing Department

### Burlesque

Part 1:

```
lnS0L[-.S1rzJcp{Jg0jd!'@==j{2rz?d?+}MPcp{{J0>=jg1<=}al}{g0jd!}FM'@CN4<=&&}fl
```

Part 2:

```
lnS0L[-.S1rzJcp{g0jd!'@==}f[1Pp{{{2rz?d?+}MPcp{{J0>=jg1<=}al}{g0jd!}FM'@CN4<=}pt
p^saPp{{'xD!}j+]}\mg0je!s0}{pP}w!p\CL++-.
```

Combined:

```
1:                 J         j                                                   && fl
C: lnS0L[-.S1rzJcp{ g0jd!'@==        {2rz?d?+}MPcp{{J0>=jg1<=}al}{g0jd!}FM'@CN4<=  }
2:                           }f[1Pp{{                                               pt

2: p^saPp{{'xD!}j+]}\mg0je!s0}{pP}w!p\CL++-.
```
