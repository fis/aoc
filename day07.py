#! /usr/bin/python3

import intcode
import itertools
import queue
import sys
import threading

prog = intcode.load('day07-input.txt' if len(sys.argv) < 2 else sys.argv[1])

# part 1

max_sig, max_phases = 0, None

if True:
    for phases in itertools.permutations((0, 1, 2, 3, 4)):
        sig = 0
        for phase in phases:
            out = []
            intcode.run(prog, stdin=[phase, sig], stdout=out)
            sig = out[0]
        if sig >= max_sig:
            max_sig, max_phases = sig, tuple(phases)

    print('{} -> {}'.format(max_phases, max_sig))

# part 2

max_sig, max_phases = 0, None

if True:
    for phases in itertools.permutations((5, 6, 7, 8, 9)):
        pipes = [queue.Queue() for _ in phases]
        for i, phase in enumerate(phases):
            pipes[i].put(phase)
        pipes[0].put(0)

        def amp(i):
            intcode.run(prog, stdin=pipes[i], stdout=pipes[(i+1) % len(pipes)])
        amps = [threading.Thread(target=amp, args=(i,)) for i in range(len(pipes))]

        for amp in amps:
            amp.start()
        for amp in amps:
            amp.join()

        sig = pipes[0].get()
        if sig >= max_sig:
            max_sig, max_phases = sig, tuple(phases)

    print('{} -> {}'.format(max_phases, max_sig))
