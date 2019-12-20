#! /usr/bin/python3

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
