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


import os.path
import subprocess
import sys

def output_file(cmd):
    base = os.path.splitext(cmd)[0]
    return base + '-output.txt'

if len(sys.argv) > 1:
    tests = sys.argv[1:]
else:
    tests = []
    for day in range(1, 26):
        cmd = './day{:02}.py'.format(day)
        if not os.path.exists(cmd) or not os.path.exists(output_file(cmd)):
            continue
        tests.append(cmd)

for cmd in tests:
    want_file = output_file(cmd)
    if not os.path.exists(want_file):
        raise RuntimeError('no known good output: ' + want_file)
    with open(want_file) as f:
        want = f.read()

    got = subprocess.run([cmd], check=True, capture_output=True, encoding='utf-8').stdout

    if got == want:
        print('{} ... pass'.format(cmd))
    else:
        print('{} ... FAIL'.format(cmd))
        print('got:\n---\n{}\n---'.format(got))
        print('want:\n---\n{}\n---'.format(want))
        break
