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

Since picking digit $d_1$ over $d_2$, where $d_1 > d_2$, as the first digit is
strictly better, no matter what happens later on in the number, we can solve
this with a greedy algorithm that just always picks the earliest instance of the
largest digit within the first $n-r$ characters, where $n$ is the length of the
string, and $r$ the number of digits that still need to be selected after the
current one, then proceeding to apply this repeatedly on what remains of the
string after the first selected digit.

### Burlesque

By contrast, day 3 _was_ pretty Burlesque-friendly.

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

The solution here is pretty much exactly what the puzzle description states: it
iterates over the level, and for each `@` counts the number of other `@`s in the
Moore neighborhood. If it's less than 4, the coordinate is added to the lits of
rolls to remove. Then the level is updated, numbers tallied, rinse and repeat.

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

Again, quite straight-forward. We parse the list of intervals, sort them by
their lower endpoint, and then do a single pass over the list to merge the
intervals into a set of non-overlapping, maximal-length ones that cover the same
values.

After that, part 1 can be solved by just testing each ingredient (binary search
on the already-sorted interval list), while part 2 is just the sum of the
lengths of the merged intervals.

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

This is really just an input parsing problem, with nothing "algorithmic" to
speak of.

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

We just iterate row by row (in fact skipping every other row since they're
always empty), maintaining a counter for each position where an active tachyon
beam might be that tracks the number of paths required to get there.

If a beam hits empty space, it's just propagated forward. If it hits a splitter,
the number of paths it was "carrying" is added to both the neighboring cells,
reflecting the choice that each of those paths could now go one of two ways.

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

Today's solution revolves around a heap and a map tracking, for each junction
box, which circuit (using an arbitrary labeling of them) it's in.

The heap is initialized to contain all the pairwise distances between the boxes.
The circuit map is initialized to put each box in its own, singleton circuit.
The solution then keeps popping off pairs of boxes in increasing order of
distance from the heap, and "connecting" them by updating the circuit map, until
all boxes are in a single circuit, pausing after 1000 to also record the
solution to part 1.

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

Part 1 is simple: just iterate over all pairs and find the rectangle with the
largest area.

Part 2 got a bit complicated; not conceptually, just in the details. What the Go
code does first is to compute the signed area of the entire polygon, because
that reveals the winding (clockwise or counterclockwise), which the second part
of the solution relies on. We then iterate over the path and adjust the
coordinates to be, not the indices (conceptually, top-left corners of) red
tiles, but the actual corners of the infinitely thin lines that enclose the area
where valid rectangles must lie. These lines are then broken down into a set of
horizontal and vertical ones, after which we can again iterate over all possible
rectangles (defined by a pair of red tiles), only this time we disqualify any
rectangle that overlaps with one of the line segments from before.

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

## [Day 10](https://adventofcode.com/2025/day/10): Factory

Initial solution here just passed the problem to the `lp_solve` library (via
some `cgo` calls), but the external dependency was annoying me, so it got
replaced by a pure-Go solution.

What the Go code does is:

1. Find the lattice of all possible integer solutions to the system of (linear
   Diophantine) equations describing the button presses to achieve the required
   joltage levels. This is done via a decomposition into the
   [Hermite normal form](https://en.wikipedia.org/wiki/Hermite_normal_form), and
   finds all solutions, including the invalid (negative button presses) and
   non-optimal ones.
2. For the input problems, the resulting lattice has rank between 0 and 3
   inclusive. To find the optimal solution:
   - For rank 0, there's just the one solution, so we'll just assume it works.
   - For rank 1, there's one free variable $z \in \mathbb{Z}$. We derive bounds
     for valid solutions from the constraint that each button press count
     $x_i \ge 0$, and find the solution at one of the endpoints.
   - For rank 2, there's two free variables $z_0, z_1$. We again use the
     $x_i \ge 0$ constraints to find the vertices that define the boundary of
     valid solutions (ignoring integrality) and establish a bounding box for all
     the points. Then we iterate over all valid values of (say) $z_0$, and use
     the rank 1 solution to find the optimal $z_1$.
   - For rank 3, it's exactly the same, except calling the rank 2 solution on
     each iteration.

In a bit more detail:

Any given machine has $m$ joltage levels to achieve, and $n$ buttons. Let's
denote by $\mathbf{A}$ the $m$ by $n$ matrix of integers, where $a_{ij} = 1$ iff
joltage $i$ is increased when button $j$ is pressed, and 0 otherwise. Denoting
by $\mathbf{x}$ and $\mathbf{v}$ column vectors that indicate the button presses
and joltage levels, respectively, we're interested in finding an integer
solution for $\mathbf{A} \mathbf{x} = \mathbf{v}$, where $x_i \ge 0\ \forall i$,
that also minimizes $\sum_i x_i$.

Ignoring the constraints and optimality for now, we can find _all_ integer
solutions to this system of linear equations by decomposing $\mathbf{A}$ into
$\mathbf{H} = \mathbf{A} \mathbf{U}$, where $\mathbf{H}$ (the column-style
Hermite normal form) is lower triangular (and has a couple of other properties
besides), and $\mathbf{U}$ is unimodular.

It turns out that if we now find the (integer) solution
$\mathbf{y} = \mathbf{U}^{-1} \mathbf{x}$ to the system
$\mathbf{H} \mathbf{y} = \mathbf{v}$, then it is the case that

$$
\mathbf{y} = \begin{bmatrix} y_1 \cdots y_r \ 0 \cdots 0 \end{bmatrix}^T
$$

(where $r$ is the rank of $\mathbf{H}$), and all solutions $\mathbf{x}$ can be
obtained as $\mathbf{x} = \mathbf{U} \mathbf{y}'$, where $\mathbf{y}'$ is
$\mathbf{y}$ with arbitrary integers $z_1, \ldots, z_{n-r} \in \mathbb{Z}$ in
place of the zeroes. Or, put differently, all solutions have the form
$\mathbf{x} = \mathbf{x}_0 + \sum_i z_i \mathbf{b}_i$, where
$\mathbf{x}_0 = \mathbf{U} \mathbf{y}$ and the basis vectors $\mathbf{b}_i$ are
the $f=n-r$ last columns of $\mathbf{U}$. Since $\mathbf{H}$ is lower
triangular, finding $\mathbf{y}$ is quite easy.

Next, we need to take into account the nonnegativity constraints and optimality
of the desired solution. We can do this slighly differently depending on the
number of "free variables" ($f$).

When $f=0$, the system has a single unique solution, so that must be the optimal
one. For completeness, we could also check that it satisfies the constraints,
but here we just assume the puzzle authors are playing fair.

When $f=1$, all solutions have the form
$\mathbf{x} = \mathbf{x}_0 + z_1 \mathbf{b}_1$. We can use the set of
constraints $x_i \ge 0$ to derive lower and upper bounds for $z_1$ from every
component of $b_i$ with a positive or negative value, respectively.
Specifically:

- If $b_i < 0$: $z_1 \le \lfloor \frac{-x_i}{b_i} \rfloor$.
- If $b_i > 0$: $z_1 \ge \lceil \frac{-x_i}{b_i} \rceil$.

Intersecting all the constraints gives exact bounds for the range of $z_1$ where
the solution is valid (in practice, this seems to always be a finite range). Now
we just need to pick the one that minimizes $\sum_i x_i$. From the form of the
solution, $\sum_i x_i = \sum_i x_{0,i} + z_1 \sum_i b_i$. This means:

- If $\sum_i b_i > 0$: increasing $z_1$ also increases the number of button
  presses, so the optimal $z_1$ is at the lower bound:
  $z_1 = \max_j \lceil \frac{-x_j}{b_j} \rceil$ where $b_j > 0$.
- If $\sum_i b_i < 0$: increasing $z_1$ decreases the number of button presses,
  so the optimal solution is at the upper bound:
  $z_1 = \min_j \lfloor \frac{-x_j}{b_j} \rfloor$ where $b_j < 0$.
- If $\sum_i b_i = 0$: value of $z_1$ does not change the number of button
  presses, so we can pick any valid $z_1$.

When $f=2$, it would be in principle to do something very similar. Each
constraint $x_i = x_{0,i} + z_1 b_{1,i} + z_2 b_{2,i} \ge 0$ defines a (closed)
half-plane in the $(z_1, z_2)$ space where valid solutions could be found, and
the intersection of all those half-planes is exactly the set of valid solutions.

Likewise, we can again characterize the objective function as
$\sum_i x_i = \sum_i x_{0,i} + z_1 \sum_i b_{1,i} + z_2 \sum_i b_{2,i}$.
Ignoring integrality, we could directly find $z_1, z_2$ minimizing this function
at the point in the set of valid solutions that's furthest along in the
direction $(-\sum_i b_{1,i}, -\sum_i b_{2,i})$. This would have to be at one of
the points where the constraints' boundaries intersect, which we can find by
locating all the intersections of pairs of lines defined by
$b_{1,i} z_1 + b_{2,i} z_2 - x_i = 0$, filtering out those which the other
constraints rule out as invalid.

We are limited to integer values $z_1, z_2$, though, and locating the integer
point furthest in that direction still contained within the region doesn't seem
entirely trivial (except when the extreme intersection point happens to be one).
Instead, what we do is, establish a bounding box that encloses all the valid
intersections (the vertices of the polygon that defines the valid solutions),
and then iterate along one coordinate axis (say, $z_1$) for the integer
coordinates in the bounding box. With $z_1$ fixed, we can just call the solution
for the $f=1$ case to find the optimal $z_2$.

(There are hypothetical badly shaped regions for which this does not work, for
example ones that just have a single finite vertex, extending to infinity in
some direction. In practice, this does not seem to occur; it's possible that the
nature of the problem ensures it cannot.)

As an implementation detail, we pick the dimension $j$ with the maximal
$\left| \sum_i b_{j,i} \right|$ as the one to iterate over (as the one that's
most "in the direction of" the optimal solution), and start from the end closer
to the optimal solution. We also compute the "margin" as the best-case impact on
the objective function from selecting $z_2$ from within its bounds, and stop the
iteration early if we've found a solution that's no longer possible to improve
on. In practice, this early stopping does not seem to make a difference to the
runtime of the algorithm.

The case of $f=3$ is the natural extension to three dimensions, just using
triplets of constraints to find the intersection points, and calling the $f=2$
solution again with one coordinate fixed.

### Burlesque

For hopefully obvious reasons, part 2 hasn't really been golfed much. It's very
much a port of the Go solution, with some non-essentials (simplifying the
machines with ad-hoc rules, error checking) stripped off.

Part 1:

```
ln{WD~]g_~-{'#==}m[sas0<-2ugs1{XX{><}gB2en{\[ri2j**}ms}m[
S2L[2j**?drz{2dg++}Cmsb[-{2dg<-{}fIg2jsi{$$}r[g1==}fe2dg++}ms
```

Part 2:

```
%xP={JJ{!!}j+]g0-]jm[sarox/?-{0.>}m[?*}%xd={156048237246138057XXsi3co)pd3co)++^p
.-}%x1={g0-]g9ngz[:-]{-]0>=}ptp^{td./pd}m^>]j{td./av}m^<]r@{g0-]?*g9?+++}m[sa{<]
bx}if}%x2={g02.+g8ng[+tpS5Jcp{J\[1304XXsi2co)pdp^.-J{j\[24150523XXsi2co)pd2co{.-
}^mjtd?/0J}ifvvvv}m[NB{-1[+{{.*}Z]++0.0>=}j+]g5jal}f[sa{)[~J<]pdj>]avr@{g01!!?*
g8?+s9x1p^}m[sa{<]bx}if}if}ln{WD[-{XX{><}gB2en)\[ri}m[l_js2{g2L[[+{2j**}ms2dg[-
<-}m[sa-.rz)z?iR<>bxj+]s00{S3xP{0>=}m[2?*?d{{?*}Z]}j+]g0jm[s0xP++z?{JJ{RA}j+]{m[
}j+]g0j0aps0g2jRAs2}{{xPzi:[~([~)sc2.+tpp^p^s7s8p^./s9g0{Jg7!!g9?*{?-}j+]g8ap}m[
s0}{xP)nz++2>=}w!JxP{nz}fiS7!={g0{Jg7!!jJg3!!g7sajg3sa}m[s0}ifg0-]{g3!!}m[Jg3!!
?/0g3sa{Jg3!!{?*?-}j+]0jZ]}j0jr~g0jm[s0+.}x/ie}{g0-]saj-]L[<..<}w!{}s1Jg0-]tpj.+
saro{.+}Z]g2j{l_g1?*++x/j.-j./g1j[+s1}Z]vvg0[~Jg1?*tp)++s1j.-S0L[{{g1++}{g1s9x1
-]}{g1s8x2-]}{g0g1ng[+tpS53CB{J)~]\[xdJ{jtpJJ3rositp\[xdbxj320XXsitp\[xd[+j130XX
sitp\[xd[+jtd?/0J}ifvvvv}m[NB{-1[+{{.*}Z]++-0.1>=}j+]g5jal}f[)[~J<]pdj>]avr@{g0
[~?*g1?+s8x2p^}m[<]}}j!!e!}ms
```

## [Day 11](https://adventofcode.com/2025/day/11): Reactor

The first (and only) pure graph thing this year.

It wasn't too complicated. We assume it's a DAG, and could likely do this with a
traditional "first do a toposort, then count number of paths with a dynamic
programming approach" thing, but opted to a recursive backwards DFS with
memoization instead; it probably boils down to the same thing, really.

For part 2, the only wrinkle is to maintain separate counters for "number of
paths with neither", "number of paths with `fft`", "number of paths with `dac`"
and "number of paths with both", which can then be properly updated depending on
the node being explored.

Part 2 could again also solve part 1, except the source node is different.

### Burlesque

Found it much simpler than day 10 this time.

Part 1:

```
ln)WDs0{{"you"1}}s1%xP={J{~[}j+]Jg1jfi0>={/vg1jfe[~}{g0j{-]~]xP}FM++
Jx/bxj[+g1j[+s1}x/ie}"out"xP
```

Part 2:

```
ln)WDs08 2dg"svr"+]bxs1%xP={J{~[}j+]Jg1jfi0>={/vg1jfe[-}{g0j{-]~]xP}FMtp)++
jJ"fft"=={j{0 0 2 2}sij}ifJ"dac"=={j2.+J_+j}if+]J[-jg1j[+s1}x/ie}"out"xP[~
```

Combined:

```
1:        {{"you"1}}                                    [~
C: ln)WDs0              s1%xP={J{~[}j+]Jg1jfi0>={/vg1jfe  }{g0j{-]~]xP}FM   ++
2:        8 2dg"svr"+]bx                                [-               tp)

1: Jx/bxj[+
C:                                                     g1j[+s1}x/ie}"out"xP
2: jJ"fft"=={j{0 0 2 2}sij}ifJ"dac"=={j2.+J_+j}if+]J[-j                    [~
```

## [Day 12](https://adventofcode.com/2025/day/12): Christmas Tree Farm

Okay, I really didn't enjoy the joke. But at least it makes for a short
Burlesque solution...

### Burlesque

Just the one part on the last day, done two different ways corresponding to the
positive (sufficient area to fit N 3x3 squares) and negative (total area of all
shapes is larger than the rectangle) heuristics, which (for the puzzle inputs)
are just inverses of each other.

The former is slightly shorter, probably because it is able to ignore the shapes
completely.

Part 1, positive:

```
ln{""};;[~{':;;p^'x;;ri?i?i3?/pdjWDri++>=}ms
```

Part 1, negative:

```
ln{""};;l_{\['#CN}m[s0{':;;p^'x;;ripdjWDrig0?*++.>}ms
```

Combined:

```
+:         [~                         ?i?i3?/             >=
C: ln{""};;              {':;;p^'x;;ri       pdjWDri    ++  }ms
-:         l_{\['#CN}m[s0                           g0?*  .>
```

## Benchmarks

As usual, non-scientific benchmarks for the Go solutions:

```
BenchmarkAllDays/day=2025.01-32     81399     139491 ns/op
BenchmarkAllDays/day=2025.02-32    346077      34829 ns/op
BenchmarkAllDays/day=2025.03-32    168002      71821 ns/op
BenchmarkAllDays/day=2025.04-32      2550    4725766 ns/op
BenchmarkAllDays/day=2025.05-32    143902      84326 ns/op
BenchmarkAllDays/day=2025.06-32     71493     169607 ns/op
BenchmarkAllDays/day=2025.07-32     98474     120800 ns/op
BenchmarkAllDays/day=2025.08-32      2044    5876190 ns/op
BenchmarkAllDays/day=2025.09-32      3398    3520390 ns/op
BenchmarkAllDays/day=2025.10-32     13728     873956 ns/op
BenchmarkAllDays/day=2025.11-32     25617     466415 ns/op
BenchmarkAllDays/day=2025.12-32     79588     150799 ns/op
```

Relatively cheap again, with no solutions over the 10 millisecond mark, and a
median time of 0.16 ms.
