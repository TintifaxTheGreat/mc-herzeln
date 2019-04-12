package main

import (
	"math/bits"
	"math/rand"
)

var CARDSTRINGS = [COLORS * FIGURES]string{
	"HA", "HK", "HO", "HU", "HX", "H9", "H8", "H7",
	"SA", "SK", "SO", "SU", "SX", "S9", "S8", "S7",
	"PA", "PK", "PO", "PU", "PX", "P9", "P8", "P7",
	"EA", "EK", "EO", "EU", "EX", "E9", "E8", "E7",
}

type CardValue struct {
	player int
	value  int
}

type Bitmap uint64

func (b *Bitmap) ToString() string {
	s := ""
	c := Bitmap(1)
	for i := 0; i < 64; i++ {
		if *b&c != 0 {
			s += CARDSTRINGS[i] + " "
		}
		c = c << 1
	}
	return s
}

func (b *Bitmap) Set(index uint64) {
	var i Bitmap = 1 << index
	*b = *b | i
}

func (b *Bitmap) Unset(index uint64) {
	var i Bitmap = 1 << index
	*b = *b ^ i
}

// find next set bit from given position
func (b *Bitmap) Next(pos uint64) uint64 {
	pos++
	ret := uint64(0)
	c := Bitmap(1)
	c = c << pos
	for i := pos; i < 64; i++ {
		if *b&c != 0 {
			ret = i
			break
		}
		c = c << 1
	}
	return ret
}

// draw random bitcard
func (b *Bitmap) DrawRandom() uint64 {
	size := bits.OnesCount64(uint64(*b))
	if size == 0 {
		return 0
	}
	count := 1 + rand.Intn(size)
	index := uint64(0)
	for i := 0; i < count; i++ {
		index = b.Next(index)
	}
	b.Unset(index)
	return index
}

// given a lead card, calculate all cards legal to pass
func (b *Bitmap) LegalCards(leadCard int, followSuit bool) (*Bitmap, bool) {
	color := int(leadCard / COLORS)
	check := *b & ALLCOLORS[color]
	size := bits.OnesCount64(uint64(check))
	if size == 0 {
		return b, false
	}
	legalCards := NewBitcard(false)
	legalCards = check
	return legalCards, true
}
*/
