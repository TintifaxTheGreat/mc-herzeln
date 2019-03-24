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

func (b *Bitcard) ToString(cardstrings [COLORS * FIGURES]string) string {
	var s string = ""
	var thisItem int = 0
	var nextItem int
	for {
		nextItem = b.c.Next(thisItem)
		if -1 == nextItem {
			break
		}
		s += cardstrings[nextItem-1] + " "
		thisItem = nextItem

	}
	return s
}

// draw random bitcard and move it to other bitcard
func (bFrom *Bitcard) DrawRandom(bTo *Bitcard) {
	size := bFrom.c.Size()-1
	// TODO improve this
	count := 1+rand.Intn(size)
	var index int = 0
	for i:=0; i<count; i++ {
		index = bFrom.c.Next(index)
	}
	bFrom.c.Delete(index)
	bTo.c.Add(index)
}
