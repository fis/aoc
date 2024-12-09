// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package day09 solves AoC 2024 day 9.
package day09

import (
	"container/heap"
	"fmt"
	"math"

	"github.com/fis/aoc/glue"
)

func init() {
	glue.RegisterSolver(2024, 9, glue.LineSolver(solve))
}

func solve(lines []string) ([]string, error) {
	if len(lines) != 1 {
		return nil, fmt.Errorf("expected one line of input, got %d", len(lines))
	}
	diskMap := lines[0]
	disk, fileMap, freeLists := expandMap(diskMap)

	compactBlocks(disk)
	part1 := checksumDisk(disk)

	compactFiles(fileMap, &freeLists)
	part2 := checksumFiles(fileMap)

	return glue.Ints(part1, part2), nil
}

type inode struct {
	pos  int
	size int
}

type freeList []int // heap

func expandMap(diskMap string) (disk []int, fileMap []inode, freeLists [9]freeList) {
	// this prealloc cuts the runtime by close to 45%
	disk = make([]int, 0, (len(diskMap)+1)/2*9)
	fileMap = make([]inode, 0, (len(diskMap)+1)/2)

	for i, N := 0, len(diskMap); i < N; i += 2 {
		file := len(fileMap)

		size := int(diskMap[i] - '0')
		fileMap = append(fileMap, inode{pos: len(disk), size: size})
		for range size {
			disk = append(disk, file)
		}

		if i+1 < N {
			size := int(diskMap[i+1] - '0')
			if size > 0 {
				freeLists[size-1] = append(freeLists[size-1], len(disk))
				for range size {
					disk = append(disk, -1)
				}
			}
		}
	}

	return disk, fileMap, freeLists
}

func compactBlocks(disk []int) {
	nextFree, nextFile := 0, len(disk)-1
	for nextFile > nextFree {
		if disk[nextFree] != -1 {
			nextFree++
			continue
		}
		if disk[nextFile] == -1 {
			nextFile--
			continue
		}
		disk[nextFree], disk[nextFile] = disk[nextFile], -1
	}
}

func checksumDisk(disk []int) (sum int) {
	for i, f := range disk {
		if f < 0 {
			continue
		}
		sum += i * f
	}
	return sum
}

func compactFiles(fileMap []inode, freeLists *[9]freeList) {
	for fileID := len(fileMap) - 1; fileID >= 0; fileID-- {
		file := fileMap[fileID]
		slotAt, slotSize := math.MaxInt, 0
		for size := file.size; size <= 9; size++ {
			if len(freeLists[size-1]) == 0 {
				continue
			}
			nextSlot := freeLists[size-1][0]
			if nextSlot < slotAt && nextSlot < file.pos {
				slotAt, slotSize = nextSlot, size
			}
		}
		if slotSize == 0 {
			continue
		}
		fileMap[fileID].pos = slotAt
		heap.Pop(&freeLists[slotSize-1])
		if gap := slotSize - file.size; gap > 0 {
			heap.Push(&freeLists[gap-1], slotAt+file.size)
		}
	}
}

func checksumFiles(fileMap []inode) (sum int) {
	for fileID, file := range fileMap {
		sum += fileID * file.size * (2*file.pos + file.size - 1) / 2
	}
	return sum
}

// heap.Interface implementation

func (h freeList) Len() int           { return len(h) }
func (h freeList) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h freeList) Less(i, j int) bool { return h[i] < h[j] }

func (h *freeList) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *freeList) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
