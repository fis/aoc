#! /usr/bin/python3

import os.path
import subprocess
import sys

if len(sys.argv) > 1:
    tests = sys.argv[1:]
else:
    tests = []
    for day in range(1, 26):
        cmd = './day{:02}.py'.format(day)
        if not os.path.exists(cmd):
            break
        tests.append(cmd)

for cmd in tests:
    cmd_base, ext = os.path.splitext(cmd)
    want_file = cmd_base + '-output.txt'
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
