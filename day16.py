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

with open('day16-input.txt') as f:
    sig = f.readline().rstrip()

# examples part 1
#sig = '12345678'
#sig = '80871224585914546619083218645595'
#sig = '19617804207202209144916044189917'
#sig = '69317163492948606335995924319873'
# examples part 2
#sig = '03036732577212944063491565474664'
#sig = '02935109699940807407585447034323'
#sig = '03081770884921959731165446850517'

sig = [int(d) for d in sig]

# part 1

def part1(sig):
    sig = sig.copy()
    for phase in range(100):
        res = []
        for n in range(len(sig)):
            r = 0
            for i, d in enumerate(sig):
                m = (i+1) // (n+1) % 4
                if m == 1:   r += d
                elif m == 3: r -= d
            res.append(abs(r) % 10)
        sig = res
    return sig

print(''.join(str(d) for d in part1(sig)[:8]))

# part 2

def part2(sig):
    offset = int(''.join(str(d) for d in sig[:7]), 10)
    if offset < 10000*len(sig) / 2:
        raise RuntimeError('unable to extract digits from first half')
    rsig = []
    for i in range(10000*len(sig)-1, offset-1, -1):
        rsig.append(sig[i % len(sig)])
    for phase in range(100):
        cnt = 0
        for i in range(len(rsig)):
            cnt = (cnt + rsig[i]) % 10
            rsig[i] = cnt
    return rsig[-1:-9:-1]

print(''.join(str(d) for d in part2(sig)))
