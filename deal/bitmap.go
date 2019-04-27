package deal

import (
	"math/bits"
	"math/rand"
)

// a hand (or deck) of cards, represented as a bit map
type bitmap uint64

// factory for Bitmap
func newBitmap(set bool) *bitmap {
	if set {
		newBitmap := bitmap(1<<uint64(COLORS*FIGURES) - 1)
		return &newBitmap
	}
	return new(bitmap)
}

// set bit at index
func (b *bitmap) set(index uint) {
	var i bitmap = 1 << index
	*b = *b | i
}

// unset bit at index
func (b *bitmap) unset(index uint) {
	var i bitmap = 1 << index
	*b = *b & ^i
}

// true if bit at index is set
func (b *bitmap) isSet(index uint) bool {
	c := bitmap(0)
	c.set(index)
	result := c & *b
	return 1 == bits.OnesCount64(uint64(result))
}

// find next set bit from given position
func (b *bitmap) next(pos uint) uint {
	pos++
	ret := uint(0)
	c := bitmap(1)
	c = c << pos
	for i := pos; i < BITMAP_SIZE; i++ {
		if *b&c != 0 {
			ret = i
			break
		}
		c = c << 1
	}
	return ret
}

// draw random bitcard
func (b *bitmap) drawRandom() uint {
	size := bits.OnesCount64(uint64(*b))
	if size == 0 {
		return 0
	}
	count := 1 + rand.Intn(size)
	index := uint(0)
	for i := 0; i < count; i++ {
		index = b.next(index)
	}
	return index
}

// string representation of Bitmap
func (b *bitmap) ToString() string {
	result := ""
	for index := uint(0); index < BITMAP_SIZE; index++ {
		if b.isSet(index) {
			result += CARDSTRINGS[index] + " "
		}
	}
	return result
}
