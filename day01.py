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


# part 1

with open('day01-input.txt') as f:
    modules = [int(ln.rstrip()) for ln in f.readlines()]

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
