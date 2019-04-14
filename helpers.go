package main

import "fmt"

func info(s string) {
	if INFO {
		fmt.Println(s)
	}
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

type Helper struct{}

func (h *Helper) AllColors() [COLORS] Bitmap {
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

func (h *Helper) AllFigures() [FIGURES] Bitmap {
	var res = [FIGURES] Bitmap{}
	for figure := uint(0); figure < FIGURES; figure++ {
		for color := uint(0); color < COLORS; color++ {
			res[figure].Set(color*FIGURES + figure)
		}
	}
	return res
}

func (h *Helper) Cardstrings() [COLORS * FIGURES]string {
	var result [COLORS * FIGURES] string
	for color := uint(0); color < COLORS; color++ {
		for figure := uint(0); figure < FIGURES; figure++ {
			result[color*FIGURES+figure] = string(COLOR_CHARS[color]) + string(FIGURE_CHARS[figure])
		}
	}
	return result
}
