package util

import "math/bits"

const (
	bitmapCoreSize = 8
	bitmapOffset   = 2
)

type Bitmap2D struct {
	core   [bitmapCoreSize][bitmapCoreSize]uint64
	astral map[P]uint64
}

func blockAddr(x, y int) (bx, by, ox, oy int) {
	bx, by = x>>3+bitmapOffset, y>>3+bitmapOffset
	ox, oy = x&7, y&7
	return bx, by, ox, oy
}

func coreBlock(bx, by int) bool {
	return bx&^(bitmapCoreSize-1) == 0 && by&^(bitmapCoreSize-1) == 0
}

func (bmp *Bitmap2D) astralize() {
	if bmp.astral == nil {
		bmp.astral = make(map[P]uint64)
	}
}

func (bmp *Bitmap2D) Get(x, y int) bool {
	bx, by, ox, oy := blockAddr(x, y)
	var b uint64
	if coreBlock(bx, by) {
		b = bmp.core[by][bx]
	} else {
		b = bmp.astral[P{bx, by}]
	}
	return b&(1<<(oy<<3|ox)) != 0
}

func (bmp *Bitmap2D) Set(x, y int) {
	bx, by, ox, oy := blockAddr(x, y)
	bit := uint64(1) << (oy<<3 | ox)
	if coreBlock(bx, by) {
		bmp.core[by][bx] |= bit
	} else {
		bmp.astralize()
		bmp.astral[P{bx, by}] |= bit
	}
}

func (bmp *Bitmap2D) Count() (sum int) {
	for _, row := range bmp.core {
		for _, b := range row {
			sum += bits.OnesCount64(b)
		}
	}
	for _, b := range bmp.astral {
		sum += bits.OnesCount64(b)
	}
	return sum
}
