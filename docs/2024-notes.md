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

## [Day 6](https://adventofcode.com/2024/day/6): Guard Gallivant

The code looks quite convoluted, but the gist of it is:

- For part 1, it just walks the path, marking every position with `X`, and
  counting every time it hits a new one.
- For part 2, it does a similar walk, but every time just before stepping into a
  new `.` square, it considers what would happen if that were `#` instead. This
  consideration is done by a *third* copy of the guard-walking routine. The two
  related to part 2 also maintain a (x, y, direction) bitmap to detect loops.

## [Day 7](https://adventofcode.com/2024/day/7): Bridge Repair

My initial solution here was pretty elegant-looking loop over the number of
possibilities ($2^{N-1}$ for part 1, $3^{N-1}$ for part 2), using the base-2/3
digits of the loop counter as the operators. Unfortunately it turned out to be
significantly slower than the recursive approach, probably because it did not
reuse the partial results, but re-evaluated every possible choice from scratch.

## [Day 8](https://adventofcode.com/2024/day/8): Resonant Collinearity

No real surprises today.

Initially thought about processing all pairs with the typical `i = 0…N-2`,
`j = i+1…N-1` loops, but since you have to do a pretty "symmetric" thing in both
parts, instead just doing two nested identical loops and skipping the `i = j`
pair.

## [Day 9](https://adventofcode.com/2024/day/9): Disk Fragmenter

Pretty straight-forward algorithms.

For both parts, it all starts from expanding the initial "compact" disk map to a
full disk image: an `[]int` where each disk block has a corresponding value
indicating which file ID it is part of (or `-1` if free space). That's
sufficient for part 1; for part 2, it also builds a file map (just a list of the
start position and size of each file) and a free list (for each size 1..9, the
locations with free space gaps of that size, in order).

To compact by block (part 1), the code uses two pointers: one starting at the
beginning of the disk, and another at the end. The first pointer is incremented
to iterate over all free space, skipping any files; the second is decremented to
iterate over all file blocks, skipping free space. Compaction is then just a
matter of moving file blocks from the second pointer to the first until they
cross over.

The second part is a little more sophisticated. Each free list is treated as a
heap; conveniently, a sorted list (which is it by construction) already
satisfies the heap invariant. We then iterate over the file map in reverse (so
largest file ID first), and for each file, check the free lists of sizes equal
to or greater than the file size. The leftmost gap that's to the left of the
file position (if any) is chosen, and popped off its corresponding heap. Since
here we always move complete files, we can ignore the disk image and update the
positions in the file map instead.

Calculating the checksum for the file uses the closed form of the sum of
consecutive integers; given file $f$ of size $s$ at position $p$, we have:

$$\begin{aligned}
\mathrm{checksum} &= \sum_{i=p}^{p+s-1} f i \\
&= \frac{1}{2} f s (2p+s-1)
\end{aligned}$$

Interestingly, profiling the initial solution showed that a large chunk of the
total time was spent in `expandMap`, specifically in resizing the slice.
Preallocating the disk image (with a worst-case estimate of 9 blocks per file)
and file map cut the runtime down from 2.2 ms to 1.3 ms.

### Burlesque

Really behind on the Burlesque solutions, and only did part 1 today. Even that
was a pain: first two solutions (relying on `sa`-based updates to the list)
worked on the example but were too slow to be feasible on the actual puzzle
input; second even managed to use large amounts of memory somehow.

Part 1:

```
riXX0[+2cozi{+.bxj^p{0}j.*#r.*}m^FL
JJ{z?}fI-1+]2CO{.-}m^?dj:nz<-saPp{j{0}j.*}Z]FL?+?dPP.+zi)pd++
```

## [Day 10](https://adventofcode.com/2024/day/10): Hoof It

Pretty sure _everyone_ did the same thing for this: solved part 2 before solving
part 1. I definitely did.

## [Day 11](https://adventofcode.com/2024/day/11): Plutonian Pebbles

There isn't anything particularly interesting about the solution algorithmically
(it's just memoized recursion), but spent a little while making it relatively
performant.

Go maps continue to be not so great. The runtime was approximately 15ms when I
used a single (stone, steps) → count map for the memoization. Switching to 75
separate maps (one for each "remaining steps" count) dropped it to 10ms. Using a
simple handcrafted hash table of size 8192 (with open addressing and linear
probing, and the hash function being just taking the low bits) dropped it to
4.5ms. Finally, forgetting about hashing and just keeping a table of 1024 values
(and not caching higher values at all) was the fastest, at 0.7ms.

## [Day 12](https://adventofcode.com/2024/day/12): Garden Groups

It's just a flood fill. I'm doing the recursive DFS, and lowercasing the letters
of the map to mark them as visited to prevent cycles.

For part 2, the side-counting is done by only counting the "first" edge of each
side, checking whether the edge continues "backwards" or not; the definition of
that is fairly arbitrary, as long as it's done consistently.

## [Day 13](https://adventofcode.com/2024/day/13): Claw Contraption

Today's puzzle presents itself as a minimization problem, but really there's
only ever a single solution, if any.

If you think of it geometrically, it's just a matter of representing the
location of the prize point using the motions of the buttons as the basis
vectors of the coordinate system. If the two vectors are not collinear, there's
a single, unique representation.

Finding the answer involves solving the following for $a$, $b$:

$$
a \begin{pmatrix} A_x \\ A_y \end{pmatrix}
+ b \begin{pmatrix} B_x \\ B_y \end{pmatrix}
= \begin{pmatrix} P_x \\ P_y \end{pmatrix}
$$

After some maths™, this yields:

$$
\begin{aligned}
a &= \frac{P_x B_y - P_y B_x}{d} \\
b &= \frac{P_y A_x - P_x A_y}{d} \\
d &= A_x B_y - A_y B_x
\end{aligned}
$$

A prize is achievable iff both $a$ and $b$ are integers.

## [Day 14](https://adventofcode.com/2024/day/14): Restroom Redoubt

Part 1 was straight-forward, as usual: just evaluate the robots' new positions
at $t=100$ (modulo the region size), and count how many are in each quadrant.

Part 2, on the other hand, was the vaguest in a while. What exactly does it mean
to "arrange themselves into a picture of a Christmas tree"? Who knows.

Observing the behavior of the robots for a few hundred seconds, it looks like at
a certain time every 101 seconds (width of the area) the robots mostly all pack
up to a specific vertical strip in the image; and conversely, at a different
time repeating every 103 seconds (height of the area) they do the same along the
other axis. It stands to reason to assume that the tree gets formed when the two
cycles coincide.

How to quantify this? Well, if the $N$ robots were distributed uniformly
randomly over the $S$ rows (or, respectively, columns), then we should expect
the counts of robots on each row (or column) have a mean of $\frac{N}{S}$, and a
variance of $\frac{N(S-1)}{S^2}$ (from a multinomial distribution of $N$ trials
and a uniform probability of $\frac{1}{S}$ for each of the $S$ outcomes). If we
compute the counts of the robots, and the observed sample variance deviates
significantly, that's probably the key moment in the cycle.

After finding out the first times $t_H$, $t_V$ when all the robots bunch up
horizontally and vertically, we just need to find the target time $t_{HV}$ that
satisfies both cycles:

$$
\begin{aligned}
t_{HV} &= t_H \pmod{101} \\
t_{HV} &= t_V \pmod{103}
\end{aligned}
$$

The
[Chinese reminder theorem](https://en.wikipedia.org/wiki/Chinese_remainder_theorem)
article describes various ways of solving this, but since the numbers involved
are so small (and there are only two congruences), we take the easy way out and
simply test all values $t_H + k \cdot 101$ and return the first one that is
congruent to $t_V$ (modulo 103).

## [Day 15](https://adventofcode.com/2024/day/15): Warehouse Woes

No math today, just a matter of implementing things. Nothing particularly clever
in the implementation either. Part 1 just goes ahead and does it; for part 2,
specifically for vertical box moves, there's a little recursive function to
collect the affected boxes and also verify if they are all movable.

## [Day 16](https://adventofcode.com/2024/day/16): Reindeer Maze

Surprising it took this long to get to a "shortest path in a maze" program.

The solution here is a basic implementation of Dijkstra's algorithm, on top of
the existing bucket priority queue. To solve part 2, the usual `prev` link is
extended to a list of potential predecessors; to handle the "directionality"
(i.e., that turning costs more than going straight), the state includes the
current direction.

## [Day 17](https://adventofcode.com/2024/day/17): Chronospatial Computer

Disassembled program:

```
loop:
  bst A
  bxl 3
  cdv B
  adv 3
  bxc
  bxl 5
  out B
  jnz loop
```

Alternatively, simplified to a Go snippet:

```go
for A != 0 {  // technically a do-while loop, but let's assume A != 0 initially
  B = A & 7
  B ^= 3
  C = A >> B
  A >>= 3
  B ^= C
  B ^= 5
  out(B & 7)
}
```

In other words, for each iteration through the loop, we will output one value,
based on the low bits of `A`. More specifically, if we call the 3 lowest bits
`k`, we will output the value `k ^ 3 ^ (A >> (k ^ 3)) ^ 5`, which simplifies to
`k ^ (A >> (k ^ 3)) ^ 6`.

Since the "consumed" bits of `A` will not affect later outputs in any way, if we
want the program to output any specific sequence, we should encode it into `A`
starting from the last digit.

It might be possible to work this out from first principles, but the solution
here just tests all possible bits to add to `A` to get the right output. It also
has to backtrack sometimes, because depending on what you add, you may not be
able to get the right output "later".

## [Day 18](https://adventofcode.com/2024/day/18): RAM Run

Today was surprisingly simple, and I don't think it was just me. We already had
a shortest-path maze thing (with "directionality" in the cost) on day 16, while
today's puzzle was just the shortest path in a uniform grid.

Also, using the part 1 solution to brute-force part 2 (by just recomputing the
shortest path after every change) was surprisingly fast, clocking in at ~100 ms.
The solution here, where we only recompute the path if the modification hits one
of its squares, is of course still much faster, at ~3 ms.

## [Day 19](https://adventofcode.com/2024/day/19): Linen Layout

Again a pretty simple one for a day 19, so spent a bit of effort in trying to
make it reasonably fast.

The naïve dynamic-programming approach of just having a table of how many ways
each suffix of the design can be prepared, and for each suffix just iterating
over the list of towels in the original order, ran in 23.6 ms.

The slightly more clever method of sorting the towel list initially, and then
for each suffix (let's say `abcd`) using a binary search to incrementally narrow
down the range of possible towels to those starting with `a`, `ab`, `abc`, ...,
was significantly better, at 3.0 ms.

Benchmarking the solution, most (≥90%) of the runtime was in the binary search.
Putting the towels in a trie (so that each step just needs to descend one level)
cut the runtime further down to 0.53 ms.

As a quirky little tweak, replacing the switch that mapped `w`, `u`, `b`, `r`,
`g` down to 0, 1, 2, 3, 4 with a bitwise expression that just extracts three
bits that happen to be unique (and making the trie child pointer array size 8 to
compensate) further reduced the runtime to 0.32 ms.

Also considered using packed bitstrings, since they could do prefix checks much
faster. The maximum stripe count for a towel in my input is 8, which fits well
in a machine word (packed 3 or 4 bits per stripe). The maximum design size is
60, which doesn't, but it's still manageable. However, the trie solution
probably wouldn't benefit as much from this, so that's still unimplemented.

## [Day 20](https://adventofcode.com/2024/day/20): Race Condition

The gist of the solution is: for every square that's part of the single path
from start to end, consider what other parts of the path are within ≤ 20 steps,
and how much it would help to take the shortest possible path there, ignoring
walls.

The initial solution just checked all pairs (that were sufficiently far of each
other on the path), and compared their Manhattan distance to their on-path
distance. This was a little sluggish (~50 ms), so the current one instead keeps
track of "chunks" of paths, and only checks the chunks that could be possible.
It also uses the same coordinate transform as 2022 day 15, to turn the
diamond-shaped regions into axis-aligned boxes, for convenience.

The chunk size is mostly arbitrary. Benchmarking sizes $2^1$, $2^2$, $2^3$,
$2^4$ and $2^5$, the resulting times were, respectively, 9.6, 6.6, 6.0, 7.6 and
11.1 ms. So the solution uses an 8-by-8 chunk size (in the transformed
coordinates).

There's also one more speedup trick. Due to the construction, the items of a
chunk are sorted based on their on-path position. We only need to consider items
that are at least $N$ steps further along the path; otherwise the cheat could
not possibly save enough time, even if it was zero-length. Since the chunk is
consistently ordered, we can do a quick binary search to find the starting
position to start looking at a chunk in. This drops the runtime to 4.0 ms.
