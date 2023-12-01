# AoC on Z80

This directory is a step towards sufficient scaffolding to allow writing Advent
of Code puzzle solutions in Z80 assembly. As of this writing, it is pretty
incomplete, and only contains:

- The `z80ex` package, which provides basic Cgo bindings to the
  [z80ex](https://sourceforge.net/projects/z80ex/) Z80 emulator library.
- The `validate_test.go` unit test that runs all existing solutions against the
  shared puzzle inputs and expected outputs.
- The `cmd/z80ex` binary, which can use the library to run a Z80 program with
  I/O port 1 bound to the standard input/output streams, for manual testing.
- Few Z80 utility routines for input and output of unsigned 16-bit and 32-bit
  integers.
- Solutions for 2022-01 and 2023-01 in Z80 assembly, and `//go:generate` lines
  to assemble them.

Prerequisites for using this:

- The [z80ex](https://sourceforge.net/projects/z80ex/) library, installed so
  that its include file is available as `<z80ex/z80ex.h>` and the library
  available as `-lz80ex` in the library search path.
- The [z80asm](https://git.savannah.nongnu.org/cgit/z80asm.git) assembler, and
  specifically a version including
  [commit 320cce79](https://git.savannah.nongnu.org/cgit/z80asm.git/commit/?id=320cce79f8ec862fc5d750d05519113d741871b2),
  which fixes the "label defined" test expression (`? foo`). Notably, this is
  still broken in the version packaged in Debian.

The emulated Z80 machine model is as simple as it could possibly be. The machine
has 64 kiB of RAM, filling the address space. The program is loaded at the
beginning, and the stack pointer is set to make stack start growing from the top
of the address space. Writing a byte into the I/O port 1 transmits it to
standard output, while reading a byte reads it from standard input, returning 0
if at end of file. That is all.

If I end up writing more solutions, subsequent improvements could include:

- A pure Go Z80 assembler. This would allow quality-of-life stuff like automated
  importing of library dependencies, and maybe even fancier stuff like inlining
  only-used-once routines.
- A solution for debugging. The Visual Studio Code
  [DeZog](https://github.com/maziac/DeZog) extension could be made to serve in
  this role, and its internal simulator should be capable enough for this,
  though some glue would be needed to support `z80asm` listing files.
  Alternatively, a homegrown TUI debugger might be an option.
- Potentially some sort of extended memory system, but only if some puzzle
  problems turn out to really need more RAM than that.
