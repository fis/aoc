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

# util

def run(prog, a, b):
    prog = list(prog)
    prog[1] = a
    prog[2] = b
    ip = 0
    while True:
        op = prog[ip]
        if op == 99:
            break
        pa, pb, pc = prog[ip+1], prog[ip+2], prog[ip+3]
        a, b = prog[pa], prog[pb]
        if op == 1:
            prog[pc] = a + b
        elif op == 2:
            prog[pc] = a * b
        else:
            sys.exit('bah')
        ip += 4
    return prog[0]

# part 1

with open('day02-input.txt') as f:
    prog = [int(i) for i in f.readline().split(',')]

print(run(prog, 12, 2))

# part 2

for noun in range(100):
    for verb in range(100):
        if run(prog, noun, verb) == 19690720:
            print(100 * noun + verb)
