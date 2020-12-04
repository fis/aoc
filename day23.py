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


import intcode
import itertools
import queue
import sys
import threading

class Switch:
    def __init__(self, nports):
        self._inq = queue.Queue()
        self._recvq = [[] for _ in range(nports)]
        self._recv_count = [0 for _ in range(nports)]
        self._nat_packet = None
        self._nat_send_history = (None, None)
        self._shutdown = None

    def run(self):
        while True:
            cmd, addr_from, addr_to, arg = self._inq.get()
            if cmd == 'send':
                if self._shutdown is not None:
                    pass  # shutting down, dropping packets
                elif addr_to == 255:
                    if self._nat_packet is None:
                        print(arg)
                    self._nat_packet = arg
                else:
                    self._recvq[addr_to].append(arg)
                    self._recv_count[addr_from] = 0
            elif cmd == 'recv':
                recvq = self._recvq[addr_to]
                if self._shutdown is not None:
                    self._shutdown[addr_to] = True
                    arg.put('shutdown')
                elif not recvq:
                    self._recv_count[addr_to] += 1
                    arg.put(None)
                else:
                    packet = recvq[0]
                    self._recvq[addr_to] = recvq[1:]
                    self._recv_count[addr_to] = 0
                    arg.put(packet)
            else:
                raise RuntimeError('invalid command: ' + cmd)
            if self._shutdown is not None and all(self._shutdown):
                return
            self._nat()

    def send(self, addr_from, addr_to, x, y):
        self._inq.put(('send', addr_from, addr_to, (x, y)))

    def recv(self, addr_to, outq):
        self._inq.put(('recv', None, addr_to, outq))

    def _nat(self):
        for recvq in self._recvq:
            if recvq:
                # not idle: this computer has data to receive
                return
        for recv_count in self._recv_count:
            if recv_count < 5:
                # not idle: this computer hasn't tried to receive packets
                # that many times in a row without sending
                return
        # network idle
        if self._nat_packet is None:
            raise RuntimeError('nat triggered before receiving any packets')
        self._recvq[0].append(self._nat_packet)
        self._recv_count = [0 for _ in self._recv_count]
        if self._nat_send_history[1] == self._nat_packet[1]:
            print(f'{self._nat_send_history}, {self._nat_packet}')
            self._shutdown = [False for _ in self._recvq]
        self._nat_send_history = self._nat_packet

class ShutdownException(Exception): pass

class Computer:
    def __init__(self, addr, switch, prog):
        self._addr = addr
        self._switch = switch
        self._prog = prog
        self._state = 'booting'
        self._inq = queue.Queue()
        self._incoming = None
        self._outgoing = []
        self._thread = threading.Thread(target=self._run)
        self._thread.start()

    def _run(self):
        try:
            intcode.run(self._prog, stdin=self._in, stdout=self._out)
        except ShutdownException:
            return

    def _in(self):
        if self._state == 'booting':
            self._state = 'running'
            return self._addr
        if self._incoming is not None:
            y = self._incoming
            self._incoming = None
            return y
        self._switch.recv(self._addr, self._inq)
        packet = self._inq.get()
        if packet == 'shutdown':
            raise ShutdownException()
        if packet is None:
            return -1
        x, y = packet
        self._incoming = y
        return x

    def _out(self, n):
        self._outgoing.append(n)
        if len(self._outgoing) == 3:
            to, x, y = self._outgoing
            self._outgoing.clear()
            self._switch.send(self._addr, to, x, y)

prog = intcode.load('day23-input.txt')

# part 1 & 2

N = 50

switch = Switch(N)
for a in range(N):
    Computer(a, switch, prog)
switch.run()
