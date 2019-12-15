import collections
import enum
import queue

def load(path):
    with open(path) as f:
        return [int(i) for i in f.readline().split(',')]

def run(prog, stdin=None, stdout=None, trace=False):
    vm = VM(prog)
    vm.run(stdin=stdin, stdout=stdout, trace=trace)

Opcode = collections.namedtuple('Opcode', 'name nargs action jump', defaults=(0, None, False))

Arg = collections.namedtuple('Arg', 'mode value')

class ArgMode(enum.Enum):
    INDIRECT = 0
    IMMEDIATE = 1
    RELATIVE = 2

    @classmethod
    def format(cls, arg):
        return {
            ArgMode.INDIRECT: '{}',
            ArgMode.IMMEDIATE: '#{}',
            ArgMode.RELATIVE: 'B{:+}',
        }[arg.mode].format(arg.value)

class VM:
    def __init__(self, prog):
        self._ip = 0
        self._base = 0
        self._prog = list(prog)
        self._stdin = None
        self._stdout = None

    def run(self, stdin=None, stdout=None, trace=False):
        self._stdin = stdin
        self._stdout = stdout
        while True:
            if not self.step(trace=trace):
                return

    def step_out(self, stdin=None, trace=False):
        self._stdin = stdin
        self._stdout = []
        while not self._stdout:
            if not self.step(trace=trace):
                return None
        return self._stdout[0]

    def step(self, trace=False):
        opcode, args = self._fetch(self._ip)
        if trace: self._trace(opcode, args)
        if opcode.action is None:
            return False
        opcode.action(self, *args)
        if not opcode.jump:
            self._ip += 1 + opcode.nargs
        return True

    def _trace(self, opcode, args):
        out = '{:-4d}: {}'.format(self._ip, opcode.name)
        for arg in args:
            out += ' {}'.format(ArgMode.format(arg))
        print(out)

    def _fetch(self, ip):
        if ip < 0 or ip >= len(self._prog):
            raise IndexError('fetch ip {} beyond [0, {})'.format(ip, len(self._prog)))
        op = self._prog[ip]
        opcode = VM._opcodes.get(op % 100, None)
        if opcode is None:
            raise IndexError('invalid opcode: {} ({})'.format(op % 100, op))
        args, argpos = [], 100
        for i in range(opcode.nargs):
            if ip + 1 + i >= len(self._prog):
                raise IndexError('fetch arg {} beyond [0, {})'.format(ip + 1 + i, len(self._prog)))
            args.append(Arg(
                mode=ArgMode(op // argpos % 10),
                value=self._prog[ip + 1 + i],
            ))
            argpos *= 10
        return opcode, args

    def _read(self, arg):
        if arg.mode == ArgMode.IMMEDIATE:
            return arg.value
        elif arg.mode == ArgMode.INDIRECT or arg.mode == ArgMode.RELATIVE:
            p = self._base if arg.mode == ArgMode.RELATIVE else 0
            p += arg.value
            if p < 0: raise RuntimeError('read p {}'.format(p))
            if p >= len(self._prog): self._prog.extend([0] * (p - len(self._prog) + 1))
            return self._prog[p]
        else:
            raise RuntimeError('invalid read arg: ' + repr(arg))

    def _write(self, arg, n):
        if arg.mode == ArgMode.INDIRECT or arg.mode == ArgMode.RELATIVE:
            p = self._base if arg.mode == ArgMode.RELATIVE else 0
            p += arg.value
            if p < 0: raise RuntimeError('read p {}'.format(p))
            if p >= len(self._prog): self._prog.extend([0] * (p - len(self._prog) + 1))
            self._prog[p] = n
        else:
            raise RuntimeError('invalid write arg: ' + repr(arg))

    def _input(self):
        if self._stdin is None:
            return int(input('? '))
        elif self._stdin:
            if type(self._stdin) == queue.Queue:
                n = self._stdin.get()
            elif callable(self._stdin):
                n = self._stdin()
            else:
                n = self._stdin[0]
                self._stdin = self._stdin[1:]
            if self._stdout is None:
                print('? -> {}'.format(n))
            return n
        else:
            raise RuntimeError('read past provided input')

    def _op_add(self, a1, a2, dst):
        self._write(dst, self._read(a1) + self._read(a2))

    def _op_mul(self, a1, a2, dst):
        self._write(dst, self._read(a1) * self._read(a2))

    def _op_in(self, dst):
        self._write(dst, self._input())

    def _op_out(self, a):
        n = self._read(a)
        if self._stdout is None:
            print(n)
        elif type(self._stdout) == queue.Queue:
            self._stdout.put(n)
        elif callable(self._stdout):
            self._stdout(n)
        else:
            self._stdout.append(n)

    def _op_jnz(self, a, tgt):
        if self._read(a) != 0:
            self._ip = self._read(tgt)
        else:
            self._ip += 3

    def _op_jz(self, a, tgt):
        if self._read(a) == 0:
            self._ip = self._read(tgt)
        else:
            self._ip += 3

    def _op_setlt(self, a1, a2, dst):
        self._write(dst, int(self._read(a1) < self._read(a2)))

    def _op_seteq(self, a1, a2, dst):
        self._write(dst, int(self._read(a1) == self._read(a2)))

    def _op_setb(self, a):
        self._base += self._read(a)

    _opcodes = {
        1:  Opcode(name='add',   nargs=3, action=_op_add),
        2:  Opcode(name='mul',   nargs=3, action=_op_mul),
        3:  Opcode(name='in',    nargs=1, action=_op_in),
        4:  Opcode(name='out',   nargs=1, action=_op_out),
        5:  Opcode(name='jnz',   nargs=2, action=_op_jnz, jump=True),
        6:  Opcode(name='jz',    nargs=2, action=_op_jz, jump=True),
        7:  Opcode(name='setlt', nargs=3, action=_op_setlt),
        8:  Opcode(name='seteq', nargs=3, action=_op_seteq),
        9:  Opcode(name='setb',  nargs=1, action=_op_setb),
        99: Opcode(name='halt'),
    }
