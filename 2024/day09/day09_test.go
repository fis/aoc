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

package day09

import (
	"testing"
)

var ex = "2333133121414131402"

func TestCompactBlocks(t *testing.T) {
	disk, _, _ := expandMap(ex)
	compactBlocks(disk)
	want := 1928
	if got := checksumDisk(disk); got != want {
		t.Errorf("checksum(compactBlocks(ex)) = %d, want %d", got, want)
	}
}

func TestCompactFiles(t *testing.T) {
	_, fileMap, freeList := expandMap(ex)
	compactFiles(fileMap, &freeList)
	want := 2858
	if got := checksumFiles(fileMap); got != want {
		t.Errorf("checksum(compactFiles(ex)) = %d, want %d", got, want)
	}
}
