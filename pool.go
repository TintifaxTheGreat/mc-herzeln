package main

type Pool struct {
	notDropped Bitcard
	onTable    Bitcard
	dropped    Bitcard
}

func NewCardpool() *Pool {
	c := new(Pool)
	c.notDropped = *NewBitcard(true)
	c.onTable = *NewBitcard(false)
	c.dropped = *NewBitcard(false)
	return c
}
