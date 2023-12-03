// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
void go_z80ex_reset(struct cpu *cpu, int reset_mem);
uint64_t go_z80ex_run(struct cpu *cpu, uintptr_t streams, uint64_t max_steps);
uint64_t go_z80ex_trace(struct cpu *cpu, uintptr_t streams);

#endif
