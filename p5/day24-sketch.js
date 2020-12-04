/**
 * Copyright 2019 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Initial state:
//   ###..
//   .##..
//   #....
//   ##..#
//   .###.
// Bits 0b0_1110_1001_1000_0100_1100_0111
// Value 0xe984c7

let states = [0, 0xe984c7, 0];
let state_offset = -1; // depth of first state in state array
let time = 0, bugs = 12;

const NEIGHS = [
    /* bit  0 */ [[0, 1, 5], [-1, 7, 11]],
    /* bit  1 */ [[0, 0, 2, 6], [-1, 7]],
    /* bit  2 */ [[0, 1, 3, 7], [-1, 7]],
    /* bit  3 */ [[0, 2, 4, 8], [-1, 7]],
    /* bit  4 */ [[0, 3, 9], [-1, 7, 13]],
    /* bit  5 */ [[0, 0, 6, 10], [-1, 11]],
    /* bit  6 */ [[0, 1, 5, 7, 11]],
    /* bit  7 */ [[0, 2, 6, 8], [1, 0, 1, 2, 3, 4]],
    /* bit  8 */ [[0, 3, 7, 9, 13]],
    /* bit  9 */ [[0, 4, 8, 14], [-1, 13]],
    /* bit 10 */ [[0, 5, 11, 15], [-1, 11]],
    /* bit 11 */ [[0, 6, 10, 16], [1, 0, 5, 10, 15, 20]],
    /* bit 12 */ [],
    /* bit 13 */ [[0, 8, 14, 18], [1, 4, 9, 14, 19, 24]],
    /* bit 14 */ [[0, 9, 13, 19], [-1, 13]],
    /* bit 15 */ [[0, 10, 16, 20], [-1, 11]],
    /* bit 16 */ [[0, 11, 15, 17, 21]],
    /* bit 17 */ [[0, 16, 18, 22], [1, 20, 21, 22, 23, 24]],
    /* bit 18 */ [[0, 13, 17, 19, 23]],
    /* bit 19 */ [[0, 14, 18, 24], [-1, 13]],
    /* bit 20 */ [[0, 15, 21], [-1, 11, 17]],
    /* bit 21 */ [[0, 16, 20, 22], [-1, 17]],
    /* bit 22 */ [[0, 17, 21, 23], [-1, 17]],
    /* bit 23 */ [[0, 18, 22, 24], [-1, 17]],
    /* bit 24 */ [[0, 19, 23], [-1, 13, 17]],
];

const GRID_SIZE = 15;
const MGRID_W = 17, MGRID_H = 12;

function setup() {
    createCanvas((MGRID_W * 6 + 1) * GRID_SIZE, (MGRID_H * 6 + 1) * GRID_SIZE);
    frameRate(5);
}

function draw() {
    background(255);

    fill(0);
    noStroke();
    text('t='+time, 20, 10);
    text('bugs='+bugs, 100, 10);

    for (let d = 101; d >= -101; d--) {
        let s = getState(d);
        // Draw the guidelines.
        stroke(128);
        for (let y = 0; y < 2; y++) {
            for (let x = 0; x < 2; x++) {
                let x1 = mapX(5*x, 5*y, d), y1 = mapY(5*x, 5*y, d);
                let x2 = mapX(2+x, 2+y, d-1), y2 = mapY(2+x, 2+y, d-1);
                if (x1 < 0 || x2 < 0 || y1 < 0 || y2 < 0) continue;
                line(x1, y1, x2, y2);
            }
        }
        // Draw the main grid.
        for (let y = 0; y < 5; y++) {
            for (let x = 0; x < 5; x++) {
                if (x == 2 && y == 2) continue;
                let px = mapX(x, y, d), py = mapY(x, y, d);
                if (px < 0 || py < 0) continue;
                fill((s >> (5*y+x)) & 1 ? 0 : 255);
                stroke(0);
                rect(px, py, GRID_SIZE, GRID_SIZE);
            }
        }
        // Draw the tiny recursive grid.
        for (let y = 0; y < 5; y++) {
            for (let x = 0; x < 5; x++) {
                let px = mapX(2+x/5, 2+y/5, d-1), py = mapY(2+x/5, 2+y/5, d-1);
                if (px < 0 || py < 0) continue;
                fill((s >> (5*y+x)) & 1 ? 0 : 255);
                stroke(0);
                rect(px, py, GRID_SIZE/5, GRID_SIZE/5);
            }
        }
    }

    time++;
    if (time == 201)
        noLoop();
    else
        step();
}

function step() {
    let nstates = [];
    for (let d = state_offset; d < state_offset + states.length; d++) {
        let ns = getState(d);
        for (let i = 0; i < NEIGHS.length; i++) {
            let count = 0;
            for (const bits of NEIGHS[i]) {
                let ps = getState(d + bits[0]);
                for (let j = 1; j < bits.length; j++) {
                    count += (ps >> bits[j]) & 1;
                }
            }
            if ((ns & (1 << i)) && count != 1) {
                ns &= ~(1 << i);
                bugs--;
            } else if (!(ns & (1 << i)) && (count == 1 || count == 2)) {
                ns |= 1 << i;
                bugs++;
            }
        }
        nstates.push(ns);
    }
    if (nstates[0] != 0) {
        nstates.unshift(0);
        state_offset--;
    }
    if (nstates[nstates.length-1] != 0) {
        nstates.push(0);
    }
    states = nstates;
}

function getState(depth) {
    let d = depth - state_offset;
    if (d >= 0 && d < states.length)
        return states[d];
    return 0;
}

function mapX(x, y, depth) {
    let d = 102 + depth;
    if (d < 0 || d >= MGRID_W*MGRID_H) return -1;
    let mx = Math.floor(d / MGRID_W) % 2 ? MGRID_W - 1 - (d % MGRID_W) : d % MGRID_W;
    return GRID_SIZE + mx*6*GRID_SIZE + x*GRID_SIZE;
}

function mapY(x, y, depth) {
    let d = 102 + depth;
    if (d < 0 || d >= MGRID_W*MGRID_H) return -1;
    return GRID_SIZE + Math.floor(d / MGRID_W)*6*GRID_SIZE + y*GRID_SIZE;
}
