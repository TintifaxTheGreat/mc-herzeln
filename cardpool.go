package main

type Cardpool struct {
	notDropped Bitcard
	onTable    Bitcard
	dropped    Bitcard
}

func NewCardpool() *Cardpool {
	c := new(Cardpool)
	c.notDropped = *NewBitcard(true)
	c.onTable = *NewBitcard(false)
	c.dropped = *NewBitcard(false)
	return c
}

/*
func main() {
	foo := NewCardpool()
	fmt.Print(foo)
}
*/