#! /usr/bin/python3

import sys

with open('day08-input.txt' if len(sys.argv) < 2 else sys.argv[1]) as f:
    data = f.readline().rstrip()

W = 25 if len(sys.argv) < 3 else int(sys.argv[2])
H = 6 if len(sys.argv) < 4 else int(sys.argv[3])

if len(data) % (W*H) != 0:
    raise RuntimeError('image size {} not a multiple of {}*{}'.format(len(data), W, H))

layers = [data[i*W*H:(i+1)*W*H] for i in range(len(data) // (W*H))]

if False:
    for layer in layers:
        for row in range(H):
            print(layer[row*W:(row+1)*W])
        print('')

# part 1

target, target_zeros = None, W*H+1

for layer in layers:
    zeros = layer.count('0')
    if zeros < target_zeros:
        target_zeros = zeros
        target = layer.count('1') * layer.count('2')

print(target)

# part 2

img = ['?'] * (W*H)

for layer in layers[::-1]:
    for i, pixel in enumerate(layer):
        if pixel == '0': img[i] = ' '
        if pixel == '1': img[i] = '#'

img = ''.join(img)

for row in range(H):
    print(img[row*W:(row+1)*W])
