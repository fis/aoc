#include <stdbool.h>
#include <stdint.h>

struct vertex {
  uint32_t degree;
	struct edge { uint32_t v, w; } next[4];
	bool seen;
};

uint32_t brute_force(struct vertex *sg, uint32_t atV, uint32_t d, uint32_t toV);
