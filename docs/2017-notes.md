# AoC 2017 solution notes

## Background

The 2017 puzzle solutions here were written shortly before (and likely after,
but as of this writing that's still in the future) the 2021 contest, as a way to
get back in the AoC mindset. Unlike the later years, where the notes are more
like a diary, here entries will only be added here if there's anything to say
that's not immediately obvious from the code.

## Day 3

The only thing of note here is that part 1 was solved "analytically", rather
than iteratively walking the spiral.

Let's consider the spiral, for convenience ignoring square 1. We can subdivide
the grid into *shells* (or *circles*), the sets of squares that have the same
Chebyshev distance (later called *radius*, denoted with `r`) from square 1.
Numbering each square consecutively from 0 in the order they are allocated in
the spiral pattern, we have a grid that looks like this:

```
+----------------------------------+
| 11   10   09   08   07   06   05 |
|    +------------------------+    |
| 12 | 07   06   05   04   03 | 04 |
|    |    +--------------+    |    |
| 13 | 08 | 03   02   01 | 02 | 03 |
|    |    |    +----+    |    |    |
| 14 | 09 | 04 | ** | 00 | 01 | 02 | ..
|    |    |    +----+    |    |    |
| 15 | 10 | 05   06   07 | 00 | 01 | ..
|    |    +--------------+    |    |
| 16 | 11   12   13   14   15 | 00 | 01
|    +------------------------+    |
| 17   18   19   20   21   22   23 | 00
+----------------------------------+
```

As mentioned, each shell has a radius `r`; the above diagram shows the shells
with radii 1, 2 and 3, as well as the beginning of the shell with radius 4.
Let's also define the *diameter* `d`, where `d = 2r+1`; the diagram shows shells
with diameters 3, 5 and 7.

We can notice that there are a total of `d^2` squares in total in all the shells
with diameters `<= d`. Therefore if we're interested in the location of square
`s` (numbered for convenience from 0 rather than 1), we can figure out the
diameter `d` of the shell it's part of by rounding `sqrt(s)` down to the next
odd integer. If we let `i = s - d*d`, then the *index* `i` will denote the
position of that square within that shell, as shown in the diagram.

> For no particular reason, there's also an implementation of an integer-only
> square root included in the solution.

Let's look at the outermost shell shown in the diagram. It has a radius `r = 3`,
a diameter `d = 5`, and contains squares where `i = 0..23`. We can further
divide all those squares into four *sides*: squares 0..5 are the *right* side,
6..11 the *top*, 12..17 the *left* and 18..23 the *bottom*. Or, more generally,
for any shell with radius `r`, there will be a total of `8r` squares in it,
with the four sides being `0..2r-1`, `2r..4r-1`, `4r..6r-1` and `6r..8r-1`.

If we look at each of the four sides in isolation, we can compute the Manhattan
distance of a square based on its index and the radius of its shell. We have:

- In the left side, the distance along the X axis is `r` for all the squares.
  Along the Y axis, the square with index `r-1` has a zero distance, so the
  distance of the other squares is therefore `abs(i-(r-1))`. The total distance
  of each square is therefore `r + abs(i-(r-1))`.
- In the top side, this time the distance along the Y axis is the constant `r`.
  For the X axis, square with index `3r-1` is at the midpoint. This gives a
  total distance of `abs(i-(3r-1)) + r`.
- The right side is the mirror image of the left, with an overall distance of
  `r + abs(i-(5r-1))`.
- The bottom side has a distance of `abs(i-(7r-1)) + r`.

Finally, we can note that (due to the symmetry of the situation and the distance
metric) we can collapse all four cases by looking at the index of the square
modulo `2r`. The distance is then: `r + abs(i % 2r - (r - 1))`.

For part 2, the analytical solution (if one exists) is... harder. The sequence
of values of the spiral is [OEIS A141481](https://oeis.org/A141481), but only a
program is given. The solution here simply iterates along the spiral, though the
iteration is done in terms of the shells of part 1.

## Day 9

I don't know why this ended so ugly. It's a hand-written recursive-descent
parser, which are supposed to look reasonably elegant. Oh well, at least it
works.

## Day 11

The only slightly unusual thing here is the use of *axial coordinates* for the
hex grid, which makes the formula for the distance quite elegant. See the
["Hexagonal Grids" page by Red Blob Games](https://www.redblobgames.com/grids/hexagons/)
for a good explanation of how this works.

## Day 15

The solution here is the straight-forward brute-force one. It runs in
approximately 0.35s on my test system: this is just on the boundary of being
annoyingly long. But there doesn't seem to be an obvious speedup.

For linear congruential generators with a power-of-two modulus, the low 16 bits
are known to have a very short period. This would probably allow a much faster
solution. But the generators here use a prime (2**31-1) as the divisor.

## Day 16

A bad case of not remembering to predict the future here.

For part 1, the solution does the obvious: applies each of the dance steps in
order. The only "trick" applied is to not shuffle data around for the *spin*
move, by instead adjusting a logical start offset of the line (treating it as a
ring buffer), and then doing just one swap (if necessary) at the end.

Of course, in part 2 the task was to run the dance a ridiculous number of times,
so the part 1 solution wasn't really reusable. The key insight here is to notice
that the dance can be decomposed to two different kinds of operations: those
that operate on positions in the line (*spin*, *exchange*) and those that are
based on program names (*partner*). Further, the two kinds of operations can be
performed entirely independently, without maintaining their relative order: we
can first do all the shuffling of the line, and then apply all the program name
swaps.

To repeat the line order operations many times, we can use the usual trick of
successive squaring. The set of moves is first boiled down into a single
permutation. To permutation can then be *squared* by applying it on itself: the
resulting permutation has the same effect as applying the original twice. This
way we can figure out the permutations that correspond to performing the moves
1, 2, 4, 8, ... times. Then we simply need to apply those permutations that
correspond to set bits in the desired iteration count to the initial state.

Finally, because the program name relabeling is always its own inverse (applying
the same swaps a second time recovers the original labels), applying it `N`
times is either the same as applying it once (for odd `N`), or not doing it at
all (for even `N`). So we boil down all the swaps to a single permutation of
program names, and apply it (once) if the count is odd.
