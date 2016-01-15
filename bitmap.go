package bitmap

import (
	"fmt"
	"math/big"
	"sort"
)

type setBits map[int]bool

type BitMap struct {
	allBits    big.Int
	allSetBits setBits
}

func New() (m *BitMap) {
	m = new(BitMap)
	m.allSetBits = make(setBits)
	return m
}

func (b *BitMap) Set(offset int) *BitMap {
	b.allBits.SetBit(&b.allBits, offset, 1)
	b.allSetBits[offset] = true
	return b
}

func (b *BitMap) Clear(offset int) *BitMap {
	b.allBits.SetBit(&b.allBits, offset, 0)
	delete(b.allSetBits, offset)
	return b
}

func (b *BitMap) Test(offset int) bool {
	return b.allBits.Bit(offset) == 1
}

func (b *BitMap) BitLen() int {
	return b.allBits.BitLen()
}

func (b *BitMap) IsEmpty() bool {
	return b.allBits.BitLen() == 0
}

func (b *BitMap) Flip(offset int) *BitMap {
	if !b.Test(offset) {
		return b.Set(offset)
	}
	return b.Clear(offset)
}

func (b BitMap) String() string {
	return fmt.Sprintf("%b", &b.allBits)
}

func (b *BitMap) GetAllSetBits(needSort bool) (allSetBits []int) {
	for i, _ := range b.allSetBits {
		allSetBits = append(allSetBits, i)
	}
	if needSort {
		sort.Ints(allSetBits)
	}
	return allSetBits
}

func (x *BitMap) And(y *BitMap) (z *BitMap) {
	z = New()
	z.allBits.And(&x.allBits, &y.allBits)
	z.allSetBits.and(&x.allSetBits, &y.allSetBits)
	return
}

func (x *BitMap) Or(y *BitMap) (z *BitMap) {
	z = New()
	z.allBits.Or(&x.allBits, &y.allBits)
	z.allSetBits.or(&x.allSetBits, &y.allSetBits)
	return
}

func (x *BitMap) Xor(y *BitMap) (z *BitMap) {
	z = New()
	z.allBits.Xor(&x.allBits, &y.allBits)
	z.allSetBits.xor(&x.allSetBits, &y.allSetBits)
	return
}

func (x *BitMap) AndNot(y *BitMap) (z *BitMap) {
	z = New()
	z.allBits.AndNot(&x.allBits, &y.allBits)
	z.allSetBits.andNot(&x.allSetBits, &y.allSetBits)
	return
}

//find out the same mapkey from x and y
func (z *setBits) and(x, y *setBits) {
	for offset := range *x {
		if _, ok := (*y)[offset]; ok {
			(*z)[offset] = true
		}
	}
	return
}

//remove the same mapkey from x and y
func (z *setBits) or(x, y *setBits) {
	for offset := range *y {
		(*z)[offset] = true
	}
	setMap(x, y, z)
	return
}

//find out the different mapkey from (x and y)+(y and x)
func (z *setBits) xor(x, y *setBits) {
	setMap(x, y, z)
	setMap(y, x, z)
	return
}

// find out the different mapkey from (x and y)
func (z *setBits) andNot(x, y *setBits) {
	setMap(x, y, z)
	return
}

// set mapkey for pointer z
func setMap(x, y, z *setBits) {
	for offset := range *x {
		if _, ok := (*y)[offset]; !ok {
			(*z)[offset] = true
		}
	}
	return
}
