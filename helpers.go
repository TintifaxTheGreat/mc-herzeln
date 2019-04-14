package main

import "fmt"

// misc helper methods
type Helper struct{}

// factory for helper
func (h *Helper) AllColors() [COLORS] bitmap {
	var res = [COLORS] bitmap{}
	index := uint(0)
	for i := uint(0); i < COLORS*FIGURES; i += FIGURES {
		temp := bitmap(1<<uint64(FIGURES) - 1)
		temp = temp << i
		res[index] = temp
		index++
	}
	return res
}

// array for each figure with the figure's bits set
func (h *Helper) AllFigures() [FIGURES] bitmap {
	var res = [FIGURES] bitmap{}
	for figure := uint(0); figure < FIGURES; figure++ {
		for color := uint(0); color < COLORS; color++ {
			res[figure].set(color*FIGURES + figure)
		}
	}
	return res
}

// array for each color with the color's bits set
func (h *Helper) Cardstrings() [COLORS * FIGURES]string {
	var result [COLORS * FIGURES] string
	for color := uint(0); color < COLORS; color++ {
		for figure := uint(0); figure < FIGURES; figure++ {
			result[color*FIGURES+figure] = string(COLOR_CHARS[color]) + string(FIGURE_CHARS[figure])
		}
	}
	return result
}

type CardValue struct {
	player uint
	value  uint
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

func info(s string) {
	if INFO {
		fmt.Println(s)
	}
}
