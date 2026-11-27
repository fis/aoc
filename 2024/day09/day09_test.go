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
