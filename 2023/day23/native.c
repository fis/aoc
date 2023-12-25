#include "native.h"

#pragma GCC push_options
#pragma GCC optimize("O2")

uint32_t brute_force(struct vertex *sg, uint32_t atV, uint32_t d, uint32_t toV) {
  if (atV == toV) {
    return d;
  }
  sg[atV].seen = true;
  uint32_t maxD = 0;
  for (uint32_t i = 0, deg = sg[atV].degree; i < deg; i++) {
    struct edge next = sg[atV].next[i];
    if (sg[next.v].seen) {
      continue;
    }
    uint32_t nextD = brute_force(sg, next.v, d+next.w, toV);
    if (nextD > maxD) {
      maxD = nextD;
    }
  }
  sg[atV].seen = false;
  return maxD;
}

#pragma GCC pop_options
