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

# part 1

def pspec(spec):
    q, chem = spec.split(' ')
    return chem, int(q)

deps, quanta = {}, {}
with open('day14-input.txt' if len(sys.argv) < 2 else sys.argv[1]) as f:
    for line in f.readlines():
        inspecs, outspec = line.rstrip().split(' => ')
        outchem, outq = pspec(outspec)
        deps[outchem] = dict(pspec(inspec) for inspec in inspecs.split(', '))
        quanta[outchem] = outq

def tsort(deps):
    deps = dict((outchem, set(outdeps.keys())) for outchem, outdeps in deps.items())
    rdeps = {}
    for outchem, outdeps in deps.items():
        for dep in outdeps:
            rdeps.setdefault(dep, set()).add(outchem)
    l = []
    s = set(('FUEL',))
    while s:
        n = s.pop()
        l.append(n)
        while deps.get(n, []):
            m = deps[n].pop()
            rdeps[m].remove(n)
            if not rdeps[m]:
                s.add(m)
    return l

odeps = tsort(deps)

def ore(fuel):
    want = {'FUEL': fuel}
    for ch in odeps:
        if ch == 'ORE' or ch not in want: continue
        n = want.pop(ch)
        q = quanta[ch]
        k = (n + q - 1) // q
        for depch, depn in deps[ch].items():
            want[depch] = want.get(depch, 0) + k*depn
    return want['ORE']

print(ore(1))

# part 2

target = 1000000000000
start, end = 1, target+1

while end - start >= 2:
    mid = start + (end - start) // 2
    got = ore(mid)
    if got > target:
        end = mid
    else:
        start = mid

print(f'{start} -> {ore(start)}')
print(f'{end} -> {ore(end)}')
