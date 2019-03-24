package main

import (
	"fmt"
	"github.com/yourbasic/bit"
)

func main() {
	//var mygame = NewGame()

	//fmt.Println("Hello, Gopher!")
	//fmt.Println(reflect.TypeOf(mygame))

	// Add all elements in the range [0, 100) to the empty set.
	A := new(bit.Set).AddRange(0, 100) // {0..99}

	// Create a new set containing the two elements 0 and 200,
	// and then add all elements in the range [50, 150) to the set.
	B := bit.New(0, 200).AddRange(50, 150) // {0 50..149 200}

	// Compute the symmetric difference A △ B.
	X := A.Xor(B)

	// Compute A △ B as (A ∖ B) ∪ (B ∖ A).
	Y := A.AndNot(B).Or(B.AndNot(A))

	// Compare the results.
	if X.Equal(Y) {
		fmt.Println(X)
	}
}
