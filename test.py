#! /usr/bin/python3

import os.path
import subprocess

day = 1
while True:
    cmd = './day{:02}.py'.format(day)
    if not os.path.exists(cmd):
        break

    with open('day{:02}-output.txt'.format(day)) as f:
        want = f.read()

    got = subprocess.run([cmd], check=True, capture_output=True, encoding='utf-8').stdout

    if got == want:
        print('day{} ... pass'.format(day))
    else:
        print('day{} ... FAIL'.format(day))
        print('got:\n---\n{}\n---'.format(got))
        print('want:\n---\n{}\n---'.format(want))
        break

    day += 1
