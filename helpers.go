package main

import (
	"fmt"
	"github.com/yourbasic/bit"
)

func info(s string) {
	fmt.Println(s)
}

/*
func card2string(index int) string {
	return CARDSTRINGS[index]
}
*/

func value(index int) int {
	return index % FIGURES
}

func allcolors() [COLORS] Bitmap {
	var res = [COLORS] Bitmap {}
	var index int = 0
	for i := 0; i < COLORS*FIGURES; i += FIGURES {
		res[index] = new(bit.Set).AddRange(i, i+FIGURES)
		index++
	}
	return res
}
