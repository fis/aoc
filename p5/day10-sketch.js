const W = 23, H = 23;
const INPUT = [
    ".###..#......###..#...#",
    "#.#..#.##..###..#...#.#",
    "#.#.#.##.#..##.#.###.##",
    ".#..#...####.#.##..##..",
    "#.###.#.####.##.#######",
    "..#######..##..##.#.###",
    ".##.#...##.##.####..###",
    "....####.####.#########",
    "#.########.#...##.####.",
    ".#.#..#.#.#.#.##.###.##",
    "#..#.#..##...#..#.####.",
    ".###.#.#...###....###..",
    "###..#.###..###.#.###.#",
    "...###.##.#.##.#...#..#",
    "#......#.#.##..#...#.#.",
    "###.##.#..##...#..#.#.#",
    "###..###..##.##..##.###",
    "###.###.####....######.",
    ".###.#####.#.#.#.#####.",
    "##.#.###.###.##.##..##.",
    "##.#..#..#..#.####.#.#.",
    ".#.#.#.##.##########..#",
    "#####.##......#.#.####.",
];

const CELL = 16;
const SWEEP_MS = 5000;
const FADE_MS = 200;

let asteroids = new Map();
let asteroidList = [];

let stage = 'part1';
let current = 0;

let bestPos = undefined;
let bestVis = 0;

let sweeps = undefined;

function setup() {
    for (let y = 0; y < INPUT.length; y++) {
        const row = INPUT[y];
        for (let x = 0; x < row.length; x++) {
            if (row.charAt(x) == '#') {
                asteroids.set(y*W+x, true);
                asteroidList.push({'x': x, 'y': y});
            }
        }
    }

    createCanvas(W*CELL + 160, H*CELL + 20); // TODO sizes
    frameRate(25);
}

function draw() {
    background(0);
    if (stage == 'part1')
        part1();
    else if (stage == 'done1')
        done1();
    else if (stage == 'part2')
        part2();
    else
        noLoop();
}

function part1() {
    // Draw asteroid field.
    let visible = 0;
    push();
    translate(10, 10);
    drawGrid();
    noStroke();
    const c = asteroidList[current];
    for (const p of asteroidList) {
        let dx = p.x - c.x, dy = p.y - c.y;
        if (dx == 0 && dy == 0) {
            fill(255, 255, 0);
        } else {
            const d = gcd(dx, dy);
            dx /= d; dy /= d;
            let tx = c.x + dx, ty = c.y + dy;
            while ((tx != p.x || ty != p.y) && !asteroids.has(ty*W+tx)) {
                tx += dx; ty += dy;
            }
            if (tx == p.x && ty == p.y) {
                visible++;
                fill(0, 255, 0);
            } else {
                fill(255, 0, 0);
            }
        }
        drawAsteroid(p.x, p.y);
    }
    pop();
    if (visible > bestVis) {
        bestPos = current;
        bestVis = visible;
    }
    // Draw status text.
    fill(255);
    textSize(12);
    push();
    translate(W*CELL+20, 20);
    text('Part 1:', 0, 0);
    text('Locating best asteroid.', 0, 20);
    text('Asteroid ' + (current+1) + '/' + asteroidList.length + ':', 0, 50);
    text('' + visible + ' visible.', 0, 70);
    text('Current best: ' + bestPos, 0, 100);
    text('' + bestVis + ' visible.', 0, 120);
    pop();
    // Update state.
    current++;
    if (current == asteroidList.length) {
        current = 0;
        stage = 'done1';
    }
}

function done1() {
    fill(255);
    textSize(12);
    push();
    translate(W*CELL+20, 20);
    text('Part 1 done!', 0, 0);
    text('Best asteroid: ' + bestPos, 0, 20);
    text('Charging laser: ' + round(current/50) + '%', 0, 50);
    pop();
    push();
    translate(10, 10);
    drawGrid();
    noStroke();
    for (let i = 0; i < asteroidList.length; i++) {
        const p = asteroidList[i];
        if (i == bestPos)
            fill(255, 255, 0);
        else
            fill(128);
        drawAsteroid(p.x, p.y);
    }
    pop();
    current += deltaTime;
    if (current > 5000) {
        current = 0;
        stage = 'part2';
    }
}

function part2() {
    if (sweeps === undefined) {
        makeSweeps();
    }
    current += deltaTime;
    // Draw the carnage.
    push();
    translate(10, 10);
    drawGrid();
    noStroke();
    let loopNeeded = false, beamNeeded = false;
    for (const si in sweeps) {
        const dull = floor(current/SWEEP_MS) != floor(sweeps[si][0].time/SWEEP_MS);
        for (const a of sweeps[si]) {
            if (a.time < current) {
                const f = 1 - (current - a.time) / FADE_MS;
                if (f < 0) continue;
                fill(255*f, 0, 0);
            } else {
                beamNeeded = true;
                fill(dull ? 128 : 255);
            }
            loopNeeded = true;
            const p = asteroidList[a.idx];
            drawAsteroid(p.x, p.y);
        }
    }
    const b = asteroidList[bestPos];
    if (beamNeeded) {
        stroke(128, 128, 255);
        strokeWeight(3);
        drawBeam(b.x, b.y, current / SWEEP_MS % 1);
    }
    noStroke();
    fill(255, 255, 0);
    drawAsteroid(b.x, b.y);
    pop();
    // Count vaporized.
    let vaporized = 0, nx = '?', ny = '?';
    for (const sweep of sweeps) {
        for (const a of sweep) {
            if (current > a.time) {
                vaporized++;
                if (a.nth == 200) {
                    const p = asteroidList[a.idx];
                    nx = p.x; ny = p.y;
                }
            }
        }
    }
    // Draw the status text
    fill(255);
    textSize(12);
    push();
    translate(W*CELL+20, 20);
    text('Part 2:', 0, 0);
    text('Vaporizing asteroids.', 0, 20);
    text('Blown up ' + vaporized + '/' + (asteroidList.length-1) + '.', 0, 50);
    text('Vaporized 200th:', 0, 80);
    text('('+nx+', '+ny+')', 0, 100);
    pop();
    if (!loopNeeded) {
        noLoop();
    }
}

function drawGrid() {
    stroke(64);
    for (let x = 0; x <= W; x++)
        line(x*CELL, 0, x*CELL, H*CELL);
    for (let y = 0; y <= H; y++)
        line(0, y*CELL, W*CELL, y*CELL);
}

function drawAsteroid(x, y) {
    const x0 = (x+0.5)*CELL, y0 = (y+0.5)*CELL, r = CELL/2-3;
    beginShape();
    for (let i = 0; i < 5; i++) {
        const th = (i/5)*2*PI + (13578*(y*W+x) % 10007) + millis() / 1000;
        const dx = r*cos(th), dy = r*sin(th);
        vertex(x0+dx, y0+dy);
    }
    endShape(CLOSE);
}

function drawBeam(sx, sy, dir) {
    sx = (sx + 0.5) * CELL;
    sy = (sy + 0.5) * CELL;
    const dx = sin(dir*2*PI), dy = -cos(dir*2*PI), len = (W+H)*CELL;
    line(sx, sy, sx+len*dx, sy+len*dy);
}

function makeSweeps() {
    sweeps = [];
    const b = asteroidList[bestPos];
    let available = [];
    for (let i = 0; i < asteroidList.length; i++)
        if (i != bestPos)
            available.push(i);
    let n = 0;
    while (available.length > 0) {
        let sweep = [], remaining = [];
        for (let i = 0; i < available.length; i++) {
            const p = asteroidList[available[i]];
            let dx = p.x - b.x, dy = p.y - b.y;
            const d = gcd(dx, dy);
            dx /= d; dy /= d;
            let tx = b.x + dx, ty = b.y + dy;
            while ((tx != p.x || ty != p.y) && !asteroids.has(ty*W+tx)) {
                tx += dx; ty += dy;
            }
            if (tx == p.x && ty == p.y)
                sweep.push({
                    'idx': available[i],
                    'time': (sweeps.length + angle(b, p)) * SWEEP_MS,
                });
            else
                remaining.push(available[i]);
        }
        sweep.sort(function(a, b) { return a.time - b.time; });
        for (const a of sweep) {
            const p = asteroidList[a.idx];
            asteroids.delete(p.y*W+p.x);
            a.nth = ++n;
        }
        sweeps.push(sweep);
        available = remaining;
    }
}

function angle(p1, p2) {
    const dx = p2.x - p1.x, dy = p2.y - p1.y;
    return (atan2(dx, -dy) + 2*PI) % (2*PI) / (2*PI);
}

function gcd(a, b) {
    if (a < 0) a = -a;
    if (b < 0) b = -b;
    while (b != 0) {
        const t = b;
        b = a % b;
        a = t;
    }
    return a;
}
