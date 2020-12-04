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

def test(n):
    digits = list(str(n))
    double = False
    for i in range(len(digits)-1):
        a, b = digits[i], digits[i+1]
        if b < a:
            return False
        if b == a:
            double = True
    return double

count = 0
for n in range(256310, 732736+1):
    if test(n):
        count += 1
print(count)

# part 2

def test2(n):
    digits = list(str(n))
    double = False
    for i in range(len(digits)-1):
        a, b = digits[i], digits[i+1]
        if b < a:
            return False
        if b == a and (i == 0 or digits[i-1] != a) and (i == len(digits)-2 or digits[i+2] != b):
            double = True
    return double

count = 0
for n in range(256310, 732736+1):
    if test2(n):
        count += 1
print(count)
