package util

import "math/bits"

type FixedBitmap1D []uint64

func MakeFixedBitmap1D(size int) FixedBitmap1D {
	wsize := (size + 63) >> 6
	return make(FixedBitmap1D, wsize)
}

func (bmp FixedBitmap1D) Clear() {
	for i := range bmp {
		bmp[i] = 0
	}
}

func (bmp FixedBitmap1D) Get(i int) bool { return bmp[i>>6]&(1<<(i&63)) != 0 }
func (bmp FixedBitmap1D) Set(i int)      { bmp[i>>6] |= 1 << (i & 63) }

type FixedBitmap2D [][]uint64

func MakeFixedBitmap2D(w, h int) FixedBitmap2D {
	ww := (w + 63) >> 6
	data := make([]uint64, ww*h)
	bmp := make(FixedBitmap2D, h)
	for y := 0; y < h; y++ {
		bmp[y] = data[y*ww : (y+1)*ww]
	}
	return bmp
}

func (bmp FixedBitmap2D) Size() (w, h int) {
	return len(bmp[0]) << 6, len(bmp)
}

func (bmp FixedBitmap2D) Get(x, y int) bool {
	wx, ox := x>>6, x&63
	return (bmp[y][wx] & (1 << ox)) != 0
}

func (bmp FixedBitmap2D) GetN(x, y, n int) uint64 {
	wx, ox := x>>6, x&63
	if ox+n <= 64 {
		return (bmp[y][wx] >> ox) & ((uint64(1) << n) - 1)
	}
	n1, n2 := 64-ox, ox+n-64
	b := (bmp[y][wx] >> ox) & ((uint64(1) << n1) - 1)
	b |= (bmp[y][wx+1] & ((uint64(1) << n2) - 1)) << n1
	return b
}

func (bmp FixedBitmap2D) Set(x, y int) {
	wx, ox := x>>6, x&63
	bmp[y][wx] |= 1 << ox
}

func (bmp FixedBitmap2D) Clear(x, y int) {
	wx, ox := x>>6, x&63
	bmp[y][wx] &^= 1 << ox
}

func (bmp FixedBitmap2D) RotateR(w int) {
	for y := range bmp {
		wc, oc := (w-1)>>6, (w-1)&63
		out := (bmp[y][wc] >> oc) & 1
		bmp[y][wc] &^= 1 << oc
		for x := wc; x >= 0; x-- {
			bmp[y][x] <<= 1
			if x-1 >= 0 {
				bmp[y][x] |= bmp[y][x-1] >> 63
			}
		}
		bmp[y][0] |= out
	}
}

func (bmp FixedBitmap2D) RotateL(w int) {
	for y := range bmp {
		wc, oc := (w-1)>>6, (w-1)&63
		out := bmp[y][0] & 1
		for x := range bmp[y] {
			bmp[y][x] >>= 1
			if x+1 < len(bmp[y]) {
				bmp[y][x] |= bmp[y][x+1] << 63
			}
		}
		bmp[y][wc] |= out << oc
	}
}

func (bmp FixedBitmap2D) Clone() (clone FixedBitmap2D) {
	clone = make(FixedBitmap2D, len(bmp))
	for i, row := range bmp {
		clone[i] = make([]uint64, len(row))
		copy(clone[i], row)
	}
	return clone
}

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

func bitmapCoords(x, y int) (px, py, bx, by, ox, oy int) {
	ox, oy = x&7, y&7
	bx, by = x>>3+bitmapOffset, y>>3+bitmapOffset
	px, py = bx>>bitmapPageBits+bitmapOffset, by>>bitmapPageBits+bitmapOffset
	bx, by = bx&(bitmapPageSize-1), by&(bitmapPageSize-1)
	return
}

func (bmp *Bitmap2D) findPage(px, py int) (page *bitmapPage) {
	if px == bitmapOffset && py == bitmapOffset {
		page = &bmp.zeroPage
	} else if px&^(bitmapPageSize-1) == 0 && py&^(bitmapPageSize-1) == 0 {
		page = bmp.nearPages[py&(bitmapPageSize-1)][px&(bitmapPageSize-1)]
	} else if p, ok := bmp.farPages[P{px, py}]; ok {
		page = p
	}
	return page
}

func (bmp *Bitmap2D) makePage(x, y int) (page *bitmapPage, bx, by, ox, oy int) {
	px, py, bx, by, ox, oy := bitmapCoords(x, y)
	page = bmp.findPage(px, py)
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
	px, py, bx, by, ox, oy := bitmapCoords(x, y)
	p := bmp.findPage(px, py)
	if p == nil {
		return false
	}
	b := p[by][bx]
	return b&(1<<(oy<<3|ox)) != 0
}

func (bmp *Bitmap2D) GetR(x0, y0, x1, y1 int) (result uint64) {
	px0, py0, bx0, by0, ox0, oy0 := bitmapCoords(x0, y0)
	px1, py1, bx1, by1, ox1, oy1 := bitmapCoords(x1, y1)
	w, h := x1-x0+1, y1-y0+1
	page := bmp.findPage(px0, py0)
	if px0 == px1 && py0 == py1 {
		if page == nil {
			return 0
		}
		if bx0 == bx1 && by0 == by1 {
			for y := oy0; y <= oy1; y++ {
				for x := ox0; x <= ox1; x++ {
					result = result<<1 | (page[by0][bx0]>>(y<<3|x))&1
				}
			}
			return result
		}
		for y, yi := by0<<3|oy0, 0; yi < h; y, yi = y+1, yi+1 {
			by, oy := y>>3, y&7
			for x, xi := bx0<<3|ox0, 0; xi < w; x, xi = x+1, xi+1 {
				bx, ox := x>>3, x&7
				result = result<<1 | (page[by][bx]>>(oy<<3|ox))&1
			}
		}
		return result
	}
	for y := y0; y <= y1; y++ {
		for x := x0; x <= x1; x++ {
			px, py, bx, by, ox, oy := bitmapCoords(x, y)
			if px != px0 || py != py0 {
				page = bmp.findPage(px, py)
				px0, py0 = px, py
			}
			result <<= 1
			if page != nil {
				result |= (page[by][bx] >> (oy<<3 | ox)) & 1
			}
		}
	}
	return result
}

func (bmp *Bitmap2D) Set(x, y int) {
	page, bx, by, ox, oy := bmp.makePage(x, y)
	page[by][bx] |= uint64(1) << (oy<<3 | ox)
}

func (bmp *Bitmap2D) GetSet(x, y int) (old bool) {
	page, bx, by, ox, oy := bmp.makePage(x, y)
	oldBlock := page[by][bx]
	mask := uint64(1) << (oy<<3 | ox)
	page[by][bx] |= mask
	return oldBlock&mask != 0
}

func (bmp *Bitmap2D) Clear(x, y int) {
	page, bx, by, ox, oy := bmp.makePage(x, y)
	page[by][bx] &^= uint64(1) << (oy<<3 | ox)
}

func (bmp *Bitmap2D) GetClear(x, y int) (old bool) {
	page, bx, by, ox, oy := bmp.makePage(x, y)
	oldBlock := page[by][bx]
	mask := uint64(1) << (oy<<3 | ox)
	page[by][bx] &^= mask
	return oldBlock&mask != 0
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
