#! /usr/bin/python3
# Copyright 2019 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.


import sys

with open('day08-input.txt' if len(sys.argv) < 2 else sys.argv[1]) as f:
    data = f.readline().rstrip()

W = 25 if len(sys.argv) < 3 else int(sys.argv[2])
H = 6 if len(sys.argv) < 4 else int(sys.argv[3])

if len(data) % (W*H) != 0:
    raise RuntimeError(f'image size {len(data)} not a multiple of {W}*{H}')

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
