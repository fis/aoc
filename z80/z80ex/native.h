#ifndef GO_Z80EX_NATIVE_H_
#define GO_Z80EX_NATIVE_H_ 1

#include <stdint.h>
#include <z80ex/z80ex.h>

struct cpu {
	Z80EX_CONTEXT *ex;
	Z80EX_BYTE ram[65536];
	uintptr_t streams;
};

struct cpu *go_z80ex_create(void);
void go_z80ex_destroy(struct cpu *cpu);
uint64_t go_z80ex_run(struct cpu *cpu, uintptr_t streams, uint64_t max_steps);
uint64_t go_z80ex_trace(struct cpu *cpu, uintptr_t streams);

#endif
