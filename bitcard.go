package main

import (
	"math/bits"
	"math/rand"
)

type CardValue struct {
	player uint
	value  uint
}

type Bitmap uint64

func NewBitmap(set bool) *Bitmap {
	b := new(Bitmap)
	if set {
		newBitmap := Bitmap(1<<uint64(COLORS*FIGURES) - 1)
		return &newBitmap
	}
	return b
}

func (b *Bitmap) ToString() string {
	result := ""
	for index := uint(0); index < BITMAP_SIZE; index++ {
		if b.IsSet(index) {
			result += CARDSTRINGS[index] + " "
		}
	}
	return result
}

func (b *Bitmap) Set(index uint) {
	var i Bitmap = 1 << index
	*b = *b | i
}

func (b *Bitmap) Unset(index uint) {
	var i Bitmap = 1 << index
	*b = *b ^ i // FIXME
}

// true if bit at index is set
func (b *Bitmap) IsSet(index uint) bool {
	c := Bitmap(0)
	c.Set(index)
	result := c & *b
	return 1 == bits.OnesCount64(uint64(result))
}

// find next set bit from given position
func (b *Bitmap) Next(pos uint) uint {
	pos++
	ret := uint(0)
	c := Bitmap(1)
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
func (b *Bitmap) DrawRandom() uint {
	size := bits.OnesCount64(uint64(*b))
	if size == 0 {
		return 0
	}
	count := 1 + rand.Intn(size)
	index := uint(0)
	for i := 0; i < count; i++ {
		index = b.Next(index)
	}
	return index
}

// given a lead card, calculate all cards legal to pass
func (b *Bitmap) LegalCards(leadCard uint, followSuit bool) (*Bitmap, bool) {
	color := uint(leadCard / FIGURES)
	legalCards := *b & ALLCOLORS[color]
	size := bits.OnesCount64(uint64(legalCards))
	if size == 0 {
		return b, false
	}
	return &legalCards, true
}
