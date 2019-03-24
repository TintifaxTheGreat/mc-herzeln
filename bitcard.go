package main

import (
	"fmt"
	"github.com/yourbasic/bit"
)

type Bitcard struct {
	numColors int
	numFigurs int
	c         [4]bit.Set
}

func NewBitcard(set bool) *Bitcard {
	b := new(Bitcard)
	// TODO change this later
	b.numColors = 4
	b.numFigurs = 8
	bitset := new(bit.Set).AddRange(0, b.numFigurs)
	for color := 0; color < b.numColors; color++ {
		b.c[color] = *bitset
	}
	return b
}

func main() {
	foo := NewBitcard()
	fmt.Print(foo)
}
