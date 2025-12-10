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

## [Day 5](https://adventofcode.com/2025/day/5): Cafeteria

### Burlesque

Part 1:

```
ln{""};;p^{ps)ab}m[s0)ri{{Jx/rm==}j+]g0jay}ms
```

Part 2:

```
lnJ""Fi.+{ps)ab}m[><g_Pp{J-]pP[~+..>{JPp}if[~PP^px/>._+Pp}m[p\CL{p^.-+.}ms
```

Combined:

```
1:   {  };;p^         s0)ri  {Jx/rm==}j+]g0jay
C: ln ""     {ps)ab}m[      {                                              }ms
2:   J  Fi.+          ><g_Pp J-]pP[~+..>{JPp}if[~PP^px/>._+Pp}m[p\CL{p^.-+.
```

## [Day 6](https://adventofcode.com/2025/day/6): Trash Compactor

### Burlesque

Part 1:

```
lnl_ttpstpjWD{'.+]psr[}Z]++
```

Part 2:

```
lnl_)XXtp)\[tt{""};;rijWD{'.+]psr[}Z]++
```

Combined:

```
1:     ttps
C: lnl_    tp             jWD{'.+]psr[}Z]++
2:     )XX   )\[tt{""};;rijWD{'.+]psr[}Z]++
```

## [Day 7](https://adventofcode.com/2025/day/7): Laboratories

### Burlesque

Rather different approaches this time: part 1 uses a list containing the column
indices of active beams (and the set operations `\\`, `UN` and `IN` to update),
while part 2 just has a list of counters for each column (with pairwise
"coercing" operators `?*` and `?+` to update).

The part 2 solution can also pretty much solve part 1, but the extra code to
count the number of non-zero results out of the splitters means it's not
actually any shorter

Part 1:

```
lng_'SFibx+]{{'.!=}fIjJx/INS0saPp\\g0?dg0?iUNUN}r[p\CL++
```

Part 2:

```
ln{{'.!=}m[}m[{S0)n!jJx/?*jg0?*JrtjRT?+?+}r[++
```

Combined:

```
1:   g_'SFibx+]       fIjJx/INS0saPp\\g0?dg0?iUNUN       p\CL
C: ln          {{'.!=}                                }r[    ++
2:                    m[}m[{S0)n!jJx/?*jg0?*JrtjRT?+?+
```

Bonus solution! Alternative part 2, more efficient but longer:

```
lng_'SFi1_+bx+]{{'^==}fIs0{s1Jg0j~[{J+.g1_+j-.}ifg1_+}^m{-]j-]==}gb{tpp^-]j++_+}m[}r[)[~++
```

This uses an alist-like structure for tracking the beams, but I couldn't figure
out a good merge operation for them.

## [Day 8](https://adventofcode.com/2025/day/8): Playground

### Burlesque

We've reached the point where I can't be bothered to try golf these...

Part 1:

```
',' r~SPS0zi)-]S1J)bxs2Jcp{^p.<}{J{g0j!!}MP?-J?*+++]}FM><1000.+{vvS3g1j!!S4jS5g1j!!S6!=
{g2{g2g6!!{Jg1jg4jsas1}m[_+}g4ap{}g6sas2}if}m^vvg2)L[<>3.+pd
```

Part 2:

```
',' r~SPS0zi)-]S1J)bxs2Jcp{^p.<}{J{g0j!!}MP?-J?*+++]}FM><{g_p^vvS3g1j!!S4jS5g1j!!S6!=
{g2{g2g6!!{Jg1jg4jsas1}m[_+}g4ap{}g6sas2}if}{g2:L[L[-.}w!vvg0g3!!-]g0g5!!-].*
```

Combined:

```
1:                                                          1000.+
C: ',' r~SPS0zi)-]S1J)bxs2Jcp{^p.<}{J{g0j!!}MP?-J?*+++]}FM><      {    vvS3g1j!!S4jS5g1j!!S6!=
2:                                                                 g_p^

1:                                             m^vvg2)L[<>3.+pd
C: {g2{g2g6!!{Jg1jg4jsas1}m[_+}g4ap{}g6sas2}if}
2:                                             {g2:L[L[-.}w!vvg0g3!!-]g0g5!!-].*
```

## [Day 9](https://adventofcode.com/2025/day/9): Movie Theater

### Burlesque

Following the lead of the Go solution here. Part 2 is abysmally slow.

Part 1:

```
',' r~SPJcp{tpJ)>]j)<]?-?ipd}m[>]
```

Part 2:

```
',' r~SPJJ2.+_+3CO{J2CO{?-}m^FL^p0.>x/0.>||j0.<x/0.<||_+j1!!?+}m[J-][+2COs0
Jcp{tpJ)>]S2j)<]S1?-?ipdg0{><Jtp)sm-]Jn!{p^J0Jg10<=jg20.>||j1g21.>||j1g11<=||}
j(-])r~j([~)r~e!}al.*}m[>]
```

Combined:

```
C: ',' r~SP
2:         JJ2.+_+3CO{J2CO{?-}m^FL^p0.>x/0.>||j0.<x/0.<||_+j1!!?+}m[J-][+2COs0

C: Jcp{tpJ}>]  j)<]  ?-?ipd
2:           S2    S1      g0{><Jtp)sm-]Jn!{p^J0Jg10<=jg20.>||j1g21.>||j1g11<=||}

C:                      }m[>]
2: j(-])r~j([~)r~e!}al.*
```
