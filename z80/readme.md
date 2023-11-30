# AoC on Z80

This directory is a step towards sufficient scaffolding to allow writing Advent
of Code puzzle solutions in Z80 assembly. As of this writing, it is quite
incomplete, and only contains:

- The `z80ex` package, which provides basic Cgo bindings to the
  [z80ex](https://sourceforge.net/projects/z80ex/) Z80 emulator library.
- The `cmd/z80ex` binary, which can use the library to run a Z80 program with
  I/O port 1 bound to the standard input/output streams.
- Few Z80 utility routines for input and output of unsigned 16-bit integers.
- A solution for the year 2022, day 1, part 1 puzzle in Z80 assembly. Sadly, the
  actual puzzle input solution is 69836, which does not fit in a 16-bit integer.

Prerequisites:

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

If I end up picking this up later, subsequent improvements should include:

- `//go:generate` lines, or some other automation, to assemble solutions without
  having to invoke `z80asm` manually.
- Integration with the shared `testdata` folder, so that it's possible to run
  `go test` to validate the generated programs.
- A solution for debugging. The Visual Studio Code
  [DeZog](https://github.com/maziac/DeZog) extension could be made to serve in
  this role, and its internal simulator should be capable enough for this,
  though some glue would be needed to support `z80asm` listing files.
  Alternatively, a homegrown TUI debugger might be an option.
- Potentially some sort of extended memory system, but only if some puzzle
  problems turn out to really need more RAM than that.
