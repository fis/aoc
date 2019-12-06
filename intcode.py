import collections
import enum

Opcode = collections.namedtuple('Opcode', 'name nargs action')

opcodes = {
    1: Opcode(
        name='add',
        nargs=3,
        action=lambda p, a: write(p, a[2], read(p, a[0]) + read(p, a[1])),
    ),
    2: Opcode(
        name='mul',
        nargs=3,
        action=lambda p, a: write(p, a[2], read(p, a[0]) * read(p, a[1])),
    ),
    3: Opcode(
        name='in',
        nargs=1,
        action=lambda p, a: write(p, a[0], prompt()),
    ),
    4: Opcode(
        name='out',
        nargs=1,
        action=lambda p, a: output(read(p, a[0])),
    ),
    # TODO better
    5: Opcode(
        name='jnz',
        nargs=2,
        action='jnz',
    ),
    6: Opcode(
        name='jz',
        nargs=2,
        action='jz',
    ),
    7: Opcode(
        name='setlt',
        nargs=3,
        action=lambda p, a: write(p, a[2], int(read(p, a[0]) < read(p, a[1]))),
    ),
    8: Opcode(
        name='seteq',
        nargs=3,
        action=lambda p, a: write(p, a[2], int(read(p, a[0]) == read(p, a[1]))),
    ),
    99: Opcode(
        name='halt',
        nargs=0,
        action=None,
    ),
}

Arg = collections.namedtuple('Arg', 'mode value')

class ArgMode(enum.Enum):
    INDIRECT = 0
    IMMEDIATE = 1

def load(path):
    with open(path) as f:
        return [int(i) for i in f.readline().split(',')]

def run(prog):
    ip = 0
    while True:
        opcode, args = decode(prog, ip)
        #print('{}: {}'.format(ip, encode(opcode, args)))
        if opcode.action is None:
            return
        # TODO better
        if opcode.action == 'jnz':
            if read(prog, args[0]) != 0: ip = read(prog, args[1])
            else: ip += 3
        elif opcode.action == 'jz':
            if read(prog, args[0]) == 0: ip = read(prog, args[1])
            else: ip += 3
        else:
            opcode.action(prog, args)
            ip += 1 + opcode.nargs

def decode(prog, ip):
    op = prog[ip]
    opcode = opcodes[op % 100]
    args, argpos = [], 100
    for i in range(opcode.nargs):
        args.append(Arg(
            mode=ArgMode(op // argpos % 10),
            value=prog[ip+1+i],
        ))
        argpos *= 10
    return opcode, args

def encode(opcode, args):
    out = opcode.name
    for arg in args:
        out += ' '
        out += ' {}{}'.format('@' if arg.mode == ArgMode.INDIRECT else '', arg.value)
    return out

def read(prog, arg):
    if arg.mode == ArgMode.INDIRECT:
        return prog[arg.value]
    elif arg.mode == ArgMode.IMMEDIATE:
        return arg.value

def write(prog, arg, n):
    if arg.mode == ArgMode.INDIRECT:
        prog[arg.value] = n

def prompt():
    return int(input('? '))

def output(n):
    print(n)
