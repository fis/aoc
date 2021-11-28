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
