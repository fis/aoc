#! /usr/bin/python3

import sys

with open('day24-input.txt' if len(sys.argv) < 2 else sys.argv[1]) as f:
    bits = f.read().replace("\n", "").replace("#", "1").replace(".", "0")
    bits = bits[::-1]
    initial_state = int(bits, 2)

# part 1

neigh = [[] for _ in range(25)]
for i in range(25):
    x, y = i % 5, i // 5
    for nx, ny in ((x-1, y), (x+1, y), (x, y-1), (x, y+1)):
        if nx < 0 or nx >= 5 or ny < 0 or ny >= 5:
            continue
        neigh[i].append(5*ny + nx)

state = initial_state
seen = set((state,))

while True:
    new = state
    for i, bits in enumerate(neigh):
        count = 0
        for bit in bits:
            count += (state >> bit) & 1
        if (state & (1 << i)) != 0 and count != 1:
            new &= ~(1 << i)
        elif (state & (1 << i)) == 0 and (count == 1 or count == 2):
            new |= (1 << i)
    state = new
    if state in seen:
        print(state)
        break
    seen.add(state)

# part 2

steps = 200

neigh = [[(0, [b for b in bits if b != 12])] for bits in neigh]
neigh[12] = []
neigh[7].append((1, list(range(0, 5))))
neigh[11].append((1, list(range(0, 25, 5))))
neigh[13].append((1, list(range(4, 25, 5))))
neigh[17].append((1, list(range(20, 25))))
neigh[0].append((-1, [7, 11]))
neigh[4].append((-1, [7, 13]))
neigh[20].append((-1, [11, 17]))
neigh[24].append((-1, [13, 17]))
for i in range(1, 4): neigh[i].append((-1, [7]))
for i in range(5, 20, 5): neigh[i].append((-1, [11]))
for i in range(9, 24, 5): neigh[i].append((-1, [13]))
for i in range(21, 24): neigh[i].append((-1, [17]))

state = [initial_state]
total_bugs = '{:b}'.format(initial_state).count('1')

for _ in range(steps):
    new, add_below, add_above = state.copy(), None, None
    for li in range(-1, len(state)+1):
        ns = new[li] if li >= 0 and li < len(new) else 0
        for i, bitlists in enumerate(neigh):
            count = 0
            for ld, bits in bitlists:
                s = state[li+ld] if li+ld >= 0 and li+ld < len(state) else 0
                for bit in bits:
                    count += (s >> bit) & 1
            if (ns & (1 << i)) != 0 and count != 1:
                total_bugs -= 1
                ns &= ~(1 << i)
            elif (ns & (1 << i)) == 0 and (count == 1 or count == 2):
                total_bugs += 1
                ns |= (1 << i)
        if li == -1:
            if ns != 0: add_below = ns
        elif li == len(state):
            if ns != 0: add_above = ns
        else:
            new[li] = ns
    state = new
    if add_below is not None:
        state.insert(0, add_below)
    if add_above is not None:
        state.append(add_above)

print(total_bugs)
