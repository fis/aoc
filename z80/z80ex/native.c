#include <stdlib.h>
#include <stdint.h>
#include <string.h>
#include <z80ex/z80ex.h>

#include "native.h"

static Z80EX_BYTE mread_cb(Z80EX_CONTEXT *ex, Z80EX_WORD addr, int m1_state, void *ctx) {
	struct cpu *cpu = ctx;
	return cpu->ram[addr];
}

static void mwrite_cb(Z80EX_CONTEXT *ex, Z80EX_WORD addr, Z80EX_BYTE value, void *ctx) {
	struct cpu *cpu = ctx;
	cpu->ram[addr] = value;
}

static Z80EX_BYTE pread_cb(Z80EX_CONTEXT *ex, Z80EX_WORD port, void *ctx) {
	if ((port & 0xff) != 1)
		return 0;
	struct cpu *cpu = ctx;
	extern uint8_t readGoInput(uintptr_t);
	return readGoInput(cpu->streams);
}

static void pwrite_cb(Z80EX_CONTEXT *ex, Z80EX_WORD port, Z80EX_BYTE value, void *ctx) {
	if ((port & 0xff) != 1)
		return;
	struct cpu *cpu = ctx;
	extern void writeGoOutput(uintptr_t, uint8_t);
	writeGoOutput(cpu->streams, value);
}

static Z80EX_BYTE intread_cb(Z80EX_CONTEXT *ex, void *ctx) {
	return 0; // TODO: what's this?
}

struct cpu *go_z80ex_create(void) {
	struct cpu *cpu = malloc(sizeof *cpu);
	if (!cpu) return 0;
	cpu->ex = z80ex_create(mread_cb, cpu, mwrite_cb, cpu, pread_cb, cpu, pwrite_cb, cpu, intread_cb, cpu);
	if (!cpu->ex) { free(cpu); return 0; }
	z80ex_set_reg(cpu->ex, regSP, 0);
	return cpu;
}

void go_z80ex_destroy(struct cpu *cpu) {
	z80ex_destroy(cpu->ex);
	free(cpu);
}

void go_z80ex_reset(struct cpu *cpu, int reset_mem) {
	z80ex_reset(cpu->ex);
	z80ex_set_reg(cpu->ex, regSP, 0);
	if (reset_mem)
		memset(cpu->ram, 0, sizeof cpu->ram);
}

uint64_t go_z80ex_run(struct cpu *cpu, uintptr_t streams, uint64_t max_steps) {
	cpu->streams = streams;
	uint64_t steps = 0;
	while (steps < max_steps && !z80ex_doing_halt(cpu->ex))
		steps += z80ex_step(cpu->ex);
	if (z80ex_doing_halt(cpu->ex))
		steps |= 0x8000000000000000;
	return steps;
}

uint64_t go_z80ex_trace(struct cpu *cpu, uintptr_t streams) {
	// TODO: implement
	return 0;
}
