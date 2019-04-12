package main

import (
	"github.com/yourbasic/bit"
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
	value int
}

type Bitcard struct {
	c bit.Set
}

func NewBitcard(set bool) *Bitcard {
	b := new(Bitcard)
	if set {
		bitset := new(bit.Set).AddRange(0, 1+COLORS*FIGURES)
		b.c = *bitset
	}
	return b
}

func (b *Bitcard) ToString() string {
	var s string = ""
	var thisItem int = 0
	var nextItem int
	for {
		nextItem = b.c.Next(thisItem)
		if -1 == nextItem {
			break
		}
		s += CARDSTRINGS[nextItem-1] + " "
		thisItem = nextItem

	}
	return s
}

func (b *Bitcard) Set(index int) {
	b.c.Add(index)
}

func (b *Bitcard) Unset(index int) {
	b.c.Delete(index)
}

// draw random bitcard
func (b *Bitcard) DrawRandom() int {
	size := b.c.Size() - 1
	// TODO improve this
	if size == 0 {
		return 0
	}
	count := 1 + rand.Intn(size)
	var index int = 0
	for i := 0; i < count; i++ {
		index = b.c.Next(index)
	}
	b.c.Delete(index)
	return index
}

// given a lead card, calculate all cards legal to pass
func(b* Bitcard) LegalCards(leadCard int, followSuit bool) (*Bitcard, bool) {
	color := int(leadCard / COLORS)
	check := *b.c.And(ALLCOLORS[color])
	if check.Size() == 0 {
		return b, false
	}
	legalCards := NewBitcard(false)
	legalCards.c = check
	return legalCards, true
}
