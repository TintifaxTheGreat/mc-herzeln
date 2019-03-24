package main

import (
	"fmt"
	"github.com/yourbasic/bit"
)

var CARDSTRINGS = [COLORS][FIGURES]string{
	{"HA", "HK", "HO", "HU", "HX", "H9", "H8", "H7"},
	{"SA", "SK", "SO", "SU", "SX", "S9", "S8", "S7"},
	{"PA", "PK", "PO", "PU", "PX", "P9", "P8", "P7"},
	{"EA", "EK", "EO", "EU", "EX", "E9", "E8", "E7"},
}

type Bitcard struct {
	c         [COLORS]bit.Set
}

func NewBitcard(set bool) *Bitcard {
	b := new(Bitcard)
	if set {

		bitset := new(bit.Set).AddRange(0, 1+FIGURES)
		for color := 0; color < COLORS; color++ {
			b.c[color] = *bitset
		}
	}
	return b
}

// TODO change this later
func (b *Bitcard) ToString(cardstrings [COLORS][FIGURES]string) string {
	var s string = ""
	for color := 0; color < COLORS; color++ {
		var thisItem int = 0
		var nextItem int
		for {
			nextItem = b.c[color].Next(thisItem)
			if -1 == nextItem {
				break
			}
			s += cardstrings[color][nextItem-1] + " "
			thisItem = nextItem

		}
	}
	return s
}

// draw random bitcard and move it to other bitcard
//func (bFrom *Bitcard) DrawRandom(bTo *Bitcard) {
//	countSet = bits.OnesCount(bFrom)
//
//}

func main() {
	foo := NewBitcard(true)
	fmt.Print(foo.ToString(CARDSTRINGS))
}
