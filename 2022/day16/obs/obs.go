// Package obs contains an obsolete, overcomplicated two-agent search solution for AoC 2022 day 16.
// It's retained for historical interest only. See the main day16 for a much more performant solution.
package obs

/*
Few more things that could have been attempted in this obsolete solution:
- Make a better lower bound by using the greedy algorithm. Also notice when it's exact and stop short.
- Make a better upper bound by considering the minimum travel times further. Maybe precompute travel time sums for subsets? Something.
- If we know the partitioning, part 2 can be solved as the sum of two instances of part 1. And the instances should be much smaller, too.
  - But bruteforcing seems to still be slower than two-agent search. Try tweaking part 1 to be very fast? Benchmark first.
- Drop the locations from the state tracker, and just consider open vertices. But it would need a better state-tracker, and I don't think it'd work: the locations are meaningful for the future.
  - Better state tracker could be done by storing all the cusp points, then doing a binary search in the diagonal projection. But if the lists stay small, not worth it. Collect stats?
*/
