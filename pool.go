package main

type Pool struct {
	notDropped Bitcard
	onTable    Bitcard
	dropped    Bitcard
}

func NewCardpool() *Pool {
	return &Pool{
		notDropped: *NewBitcard(true),
		onTable:    *NewBitcard(false),
		dropped:    *NewBitcard(false),
	}
}
