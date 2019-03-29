package main

import "fmt"

func info(s string) {
	fmt.Println(s)
}

func value(index int) int {
	// TODO change this later, or make it more elegant
	return index % FIGURES
}
