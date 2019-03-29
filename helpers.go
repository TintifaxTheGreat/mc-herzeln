package main

import (
	"fmt"
	"github.com/yourbasic/bit"
)

func info(s string) {
	fmt.Println(s)
}

func value(index int) int {
	return index % FIGURES
}

func allcolors() [COLORS]*bit.Set {
	var res = [COLORS] *bit.Set{}
	var index int = 0
	for i := 0; i < COLORS*FIGURES; i += FIGURES {
		res[index] = new(bit.Set).AddRange(i, i+FIGURES)
		index++
	}
	return res
}
