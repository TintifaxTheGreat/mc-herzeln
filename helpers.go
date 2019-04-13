package main

import (
	"fmt"
)

func info(s string) {
	fmt.Println(s)
}

func IndexOfCard(s string) (uint, bool) {
	found := false
	index := uint(0)
	for i := uint(0); i < COLORS*FIGURES; i++ {
		if s == CARDSTRINGS[i] {
			found = true
			index = i
			break
		}
	}
	return index, found
}

func value(index uint) uint {
	return index % FIGURES
}

func allcolors() [COLORS] Bitmap {
	var res = [COLORS] Bitmap{}
	index := uint(0)
	for i := uint(0); i < COLORS*FIGURES; i += FIGURES {
		temp := Bitmap(1<<uint64(FIGURES) - 1)
		temp = temp << i
		res[index] = temp
		index++
	}
	return res
}
