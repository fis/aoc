package util

import "math/bits"

const (
	bitmapPageBits = 3
	bitmapPageSize = (1 << bitmapPageBits) // side length of a square in units of 8x8 blocks
	bitmapOffset   = 2                     // in units of blocks, used to offset the zero page
)

type Bitmap2D struct {
	zeroPage  bitmapPage
	nearPages [bitmapPageSize][bitmapPageSize]*bitmapPage
	farPages  map[P]*bitmapPage
}

type bitmapPage [bitmapPageSize][bitmapPageSize]uint64

func (bmp *Bitmap2D) findPage(x, y int) (page *bitmapPage, px, py, bx, by, ox, oy int) {
	ox, oy = x&7, y&7
	bx, by = x>>3+bitmapOffset, y>>3+bitmapOffset
	px, py = bx>>bitmapPageBits+bitmapOffset, by>>bitmapPageBits+bitmapOffset
	bx, by = bx&(bitmapPageSize-1), by&(bitmapPageSize-1)
	if px == bitmapOffset && py == bitmapOffset {
		page = &bmp.zeroPage
	} else if px&^(bitmapPageSize-1) == 0 && py&^(bitmapPageSize-1) == 0 {
		page = bmp.nearPages[py&(bitmapPageSize-1)][px&(bitmapPageSize-1)]
	} else if p, ok := bmp.farPages[P{px, py}]; ok {
		page = p
	}
	return page, px, py, bx, by, ox, oy
}

func (bmp *Bitmap2D) makePage(x, y int) (page *bitmapPage, bx, by, ox, oy int) {
	page, px, py, bx, by, ox, oy := bmp.findPage(x, y)
	if page != nil {
		return page, bx, by, ox, oy
	}
	page = &bitmapPage{}
	if px&^(bitmapPageSize-1) == 0 && py&^(bitmapPageSize-1) == 0 {
		bmp.nearPages[py&(bitmapPageSize-1)][px&(bitmapPageSize-1)] = page
		return page, bx, by, ox, oy
	}
	if bmp.farPages == nil {
		bmp.farPages = make(map[P]*bitmapPage)
	}
	bmp.farPages[P{px, py}] = page
	return page, bx, by, ox, oy
}

func (bmp *Bitmap2D) Get(x, y int) bool {
	p, _, _, bx, by, ox, oy := bmp.findPage(x, y)
	if p == nil {
		return false
	}
	b := p[by][bx]
	return b&(1<<(oy<<3|ox)) != 0
}

func (bmp *Bitmap2D) Set(x, y int) {
	page, bx, by, ox, oy := bmp.makePage(x, y)
	page[by][bx] |= uint64(1) << (oy<<3 | ox)
}

func (bmp *Bitmap2D) Count() (sum int) {
	sum = bmp.zeroPage.count()
	for y := 0; y < bitmapPageSize; y++ {
		for x := 0; x < bitmapPageSize; x++ {
			if p := bmp.nearPages[y][x]; p != nil {
				sum += p.count()
			}
		}
	}
	for _, page := range bmp.farPages {
		sum += page.count()
	}
	return sum
}

func (p *bitmapPage) count() (c int) {
	for y := 0; y < bitmapPageSize; y++ {
		for x := 0; x < bitmapPageSize; x++ {
			c += bits.OnesCount64(p[y][x])
		}
	}
	return c
}
