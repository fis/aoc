#! /usr/bin/python3

# part 1

with open('day01-input.txt') as f:
    modules = [int(ln.strip()) for ln in f.readlines()]

print(sum(m // 3 - 2 for m in modules))

# part 2

def rfuel(m):
    tf = 0
    while True:
        f = m // 3 - 2
        if f <= 0:
            break
        tf += f
        m = f
    return tf

print(sum(rfuel(m) for m in modules))
