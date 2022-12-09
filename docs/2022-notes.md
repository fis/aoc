# AoC 2022 solution notes

## Background

It's that time of the year again. Back to the grind. See previous years'
introductions for more about what this is all about. In the meanwhile...

## [Day 1](https://adventofcode.com/2022/day/1): Calorie Counting

Typically soft landing. The Go code just sums and sorts to find the top three.

I've played around with extending the quickselect algorithm in the `util`
package to support the unordered partial sorting needed to find the _k_ lowest
(or highest) elements in arbitrary order in linear time. If I get around to
cleaning it up, I'll update this day's solution to use it.

### Burlesque

Part 1:

```
ln{""};;)ri)++>]
```

Part 2:

```
ln{""};;)ri)++<>3co-]++y
```

Combined:

```
1:               >]
C: ln{""};;)ri)++
2:               <>3co-]++y
```

## [Day 2](https://adventofcode.com/2022/day/2): Rock Paper Scissors

Oh, we're not doing
[Added Alliterative Appeal](https://tvtropes.org/pmwiki/pmwiki.php/Main/AddedAlliterativeAppeal)
(like [in 2020](2020-notes.md)) this year either? A shame.

Nothing interesting to say about today's Go solution. The Burlesque one does a
trick where it takes the A/B/C and X/Y/Z characters' code points modulo 3 and
treats them as the digits of a two-digit base-3 number, and uses that to just
look up the score from a map. This means both parts' solution is essentially the
same, except for the mapping.

### Burlesque

Part 1:

```
ln{XXRT[-{**3.%}m[3ug963174528XXj!!}ms
```

Part 2:

```
ln{XXRT[-{**3.%}m[3ug978123564XXj!!}ms
```

Combined:

```
1:                      963174528
C: ln{XXRT[-{**3.%}m[3ug         XXj!!}ms
2:                      978123564
```

## [Day 3](https://adventofcode.com/2022/day/3): Rucksack Reorganization

Oh, maybe we're alliterative on alternating days? I may be reading too much into
this.

The Go version uses an `[2]uint64` as the representation of a rucksack, with the
i'th bit set if an item with priority i exists in the corresponding compartment.
This makes finding the intersections real easy with bitwise operations.

Today's problem is a relatively good fit for Burlesque, what with its built-in
for finding set intersection (`IN`). Item priorities are computed as `b-38 mod
58`, which avoids needing a conditional.

### Burlesque

Part 1:

```
ln{sa2./cop^IN-]**38.-58.%}ms
```

Part 2:

```
ln3co{p^ININ-]**38.-58.%}ms
```

Combined:

```
1:       sa2./co
C: ln   {       p^IN  -]**38.-58.%}ms
2:   3co            IN
```

## [Day 4](https://adventofcode.com/2022/day/4): Camp Cleanup

Oh, great, now the titles are just actively teasing me.

Anyway, not much to say about the solutions. The Burlesque ones feel inelegant,
but they're short enough to let pass this time.

### Burlesque

Part 1:

```
ln{',;;{'-;;ri^pr@}m[Jp^~~j^p~~||}ms
```

Part 2:

```
ln{',;;{'-;;ri^pr@}MPINL[nz}ms
```

Combined:

```
1:                    m[Jp^~~j^p~~||
C: ln{',;;{'-;;ri^pr@}              }ms
2:                    MPINL[nz
```

## [Day 5](https://adventofcode.com/2022/day/5): Supply Stacks

Nothing really to say about either solution. Went very stateful with Burlesque.

### Burlesque

Part 1:

```
ln{""};;p^~])XXtp4co{1!!:rd}m[+]
{wd2enri?d^ps2s1{{g_s3}g1ap{g3+]}g2ap}j+.E!}r[)-]\[
```

Part 2:

```
ln{""};;p^~])XXtp4co{1!!:rd}m[+]
{wd2enrig_s0?d^p#r{g0cog_s1\[}apj{g1j.+}ap}r[)-]\[
```

Combined:

```
C: ln{""};;p^~])XXtp4co{1!!:rd}m[+]

1:                 s2s1{{g_s3}g1ap{g3+]}g2ap}j+.E!
C: {wd2enri    ?d^p                               }r[)-]\[
2:         g_s0    #r{g0cog_s1\[}apj{g1j.+}ap
```

## [Day 6](https://adventofcode.com/2022/day/6): Tuning Trouble

Oddly, day 6 feels simpler to me than day 5; and judging from the statistics,
I'm not the only one. It was definitely simpler to do in Burlesque.

To pad things out a little, the Go version has two alternative solutions: one
which is the naive `O(n*k)` solution but using an `uint64` as a bitset (just the
right amount of bits for upper- and lowercase ASCII letters), and another
implementing the "windowed" `O(n)` approach.

The latter does appear to outperform the former, at least on my puzzle input.
Though both are already ridiculously inexpensive.

```
$ go test -run='^$' -bench=. -benchtime=60s ./2022/day06
BenchmarkFindMarker/size=4/algo=bitset-16      17119893    4080 ns/op
BenchmarkFindMarker/size=4/algo=windowed-16    26869724    2631 ns/op
BenchmarkFindMarker/size=14/algo=bitset-16      4383006   15993 ns/op
BenchmarkFindMarker/size=14/algo=windowed-16   17832987    4008 ns/op
```

### Burlesque

Part 1:

```
4CO{><U_}fi4.+
```

Part 2:

```
14CO{><U_}fi14.+
```

Combined:

```
C:  4CO{><U_}fi 4.+
2: 1           1
```

## [Day 7](https://adventofcode.com/2022/day/7): No Space Left On Device

Today's task was really much more about parsing the input than about computing
the result.

The approach taken here (both in Go and in Burlesque) is to, instead of the
arguably more natural recursive structure, keep an explicit stack of parent
directory sizes as we go along. That way a simple scan over the input lines is
sufficient: file sizes are summed up to the size of the current directory, and
every time we ascend up to a parent directory, the subdirectory's size is added
to it also.

### Burlesque

Part 1:

```
ln{}+]{J"$ c"~!{J'.~[0j{vvPPJ#Rj_+#rPP.+}ifPp}if:><b0PP.+Pp}r[
p\CLiT)++PP.+{1e5<=}f[++
```

Part 2:

```
ln{}+]{J"$ c"~!{J'.~[0j{vvPPJ#Rj_+#rPP.+}ifPp}if:><b0PP.+Pp}r[
p\CLiT)++l_PP.+j4e7.-{>=}j+]f[<]
```

Combined:

```
C: ln{}+]{J"$ c"~!{J'.~[0j{vvPPJ#Rj_+#rPP.+}ifPp}if:><b0PP.+Pp}r[

1:                {1e5<=}        ++
C: p\CLiT)++  PP.+             f[
2:          l_    j4e7.-{>=}j+]  <]
```

## [Day 8](https://adventofcode.com/2022/day/8): Treetop Tree House

The Burlesque solutions are not great, but at least they're there. Both parts
have the same overall structure: defining a "subprogram" that finds the relevant
quantity (0/1 map of visible trees, or the viewing distances) when looking
from/to the left. The `J!ajJ)<-!a)<-jtpJ!atpj)<-!a)<-tp` sequence then applies
this to the input with the correct flips and transposes to handle all four
directions (and remap the results back to the original orientation), followed by
a short coda that combines the results from the four directions.

### Burlesque

Part 1:

```
ln)XXJtpJ)<-#RJ)<-CL{{iT[-{l_-1+]>].>}m[}m[}MP
tpj)<-tp#R)<-CL)\[tp)r|++
```

Part 2:

```
ln)XXJtpJ)<-#RJ)<-CL{{iT[-{<-sa-.jRTJ[~{>=}j+]fi+.<.}m[}m[}MP
tpj)<-tp#R)<-CL)\[tp)pd>]
```

Combined:

```
1:                            l_-1+]>].>
C: ln)XXJtpJ)<-#RJ)<-CL{{iT[-{                         }m[}m[}MP
2:                            <-sa-.jRTJ[~{>=}j+]fi+.<.

1:                     )r|++
C: tpj)<-tp#R)<-CL)\[tp
2:                     )pd>]
```

## [Day 9](https://adventofcode.com/2022/day/9): Rope Bridge

Both Go and Burlesque use the same building block: the way the tail moves
towards the head (or more generally, the next knot) is just taking the
[sign function](https://en.wikipedia.org/wiki/Sign_function) of each of the
components of the difference between the head and the tail, with one slight
adjustment that the tail doesn't move if it's already next to the head.

The Burlesque solution also uses the remarkable fact that the ASCII codes for
the letters 'U', 'L', 'R' and 'D' happen to be 0, 1, 2 and 3 mod 5,
respectively.

### Burlesque

Part 1:

```
0J_+Jx/ln{g_2336 3dg?d2coj**5.%!!{?+J#RJ#r
?-J)ab>]2>=?*)sn#R?+JPpj}j+]jri.*p^}\me!p\CL><gl
```

Part 2:

```
0J_+10.*jln{g_2336 3dg?d2coj**5.%!!{jg_x/?++]{J2J[P^p
?-J)ab>]2>=?*)snjRTg_x/?++]}9E!g_JPp[+}j+]jri.*p^}\me!p\CL><gl
```

Combined:

```
1:     Jx/                             ?+J#RJ#r
C: 0J_+     ln{g_2336 3dg?d2coj**5.%!!{
2:     10.*j                           jg_x/?++]{J2J[P^p

1:                 #R       JPpj
C: ?-J)ab>]2>=?*)sn       ?+             }j+]jri.*p^}\me!p\CL><gl
2:                 jRTg_x/  +]}9E!g_JPp[+
```
