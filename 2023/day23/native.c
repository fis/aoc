#include "native.h"

#pragma GCC push_options
#pragma GCC optimize("O2")

static uint32_t max_d;

static void recurse(struct vertex *sg, uint32_t at_v, uint32_t d, uint32_t to_v) {
  if (at_v == to_v) {
    if (d > max_d)
      max_d = d;
    return;
  }
  sg[at_v].seen = true;
  for (uint32_t i = 0, deg = sg[at_v].degree; i < deg; i++) {
    struct edge next = sg[at_v].next[i];
    if (!sg[next.v].seen)
      recurse(sg, next.v, d+next.w, to_v);
  }
  sg[at_v].seen = false;
}

uint32_t brute_force(struct vertex *sg, uint32_t from_v, uint32_t to_v) {
  max_d = 0;
  recurse(sg, from_v, 0, to_v);
  return max_d;
}

#pragma GCC pop_options
