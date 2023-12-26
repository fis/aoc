#include <stdbool.h>
#include <stdint.h>

struct vertex {
  uint32_t degree;
	struct edge { uint32_t v, w; } next[4];
	bool seen;
};

uint32_t brute_force(struct vertex *sg, uint32_t from_v, uint32_t to_v);
