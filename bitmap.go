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

func (b *BitMap) GetAllSetBits(needSort bool) (allSetBits []int) {
	for i, _ := range b.allSetBits {
		allSetBits = append(allSetBits, i)
	}
	if needSort {
		sort.Ints(allSetBits)
	}
	return allSetBits
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

func (b1 *BitMap) And(b2 *BitMap) (b3 *BitMap) {
	b3 = New()
	b3.allBits.And(&b1.allBits, &b2.allBits)
	b3.allSetBits.and(&b1.allSetBits, &b2.allSetBits)
	return
}

func (b1 *BitMap) Or(b2 *BitMap) (b3 *BitMap) {
	b3 = New()
	b3.allBits.Or(&b1.allBits, &b2.allBits)
	b3.allSetBits.or(&b1.allSetBits, &b2.allSetBits)
	return
}

func (b1 *BitMap) Xor(b2 *BitMap) (b3 *BitMap) {
	b3 = New()
	b3.allBits.Xor(&b1.allBits, &b2.allBits)
	b3.allSetBits.xor(&b1.allSetBits, &b2.allSetBits)
	return
}

func (b1 *BitMap) AndNot(b2 *BitMap) (b3 *BitMap) {
	b3 = New()
	b3.allBits.AndNot(&b1.allBits, &b2.allBits)
	b3.allSetBits.andNot(&b1.allSetBits, &b2.allSetBits)
	return
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

//find out the same mapkey from x and y
func (z *setBits) and(x, y *setBits) {
	for offset := range *x {
		if _, ok := (*y)[offset]; ok {
			(*z)[offset] = true
		}
	}
}

//remove the same mapkey from x and y
func (z *setBits) or(x, y *setBits) {
	for offset := range *y {
		(*z)[offset] = true
	}
	setMap(x, y, z)
}

//find out the different mapkey from (x and y)+(y and x)
func (z *setBits) xor(x, y *setBits) {
	setMap(x, y, z)
	setMap(y, x, z)
}

// find out the different mapkey from (x and y)
func (z *setBits) andNot(x, y *setBits) {
	setMap(x, y, z)
}

// set mapkey for pointer z
func setMap(x, y, z *setBits) {
	for offset := range *x {
		if _, ok := (*y)[offset]; !ok {
			(*z)[offset] = true
		}
	}
}
