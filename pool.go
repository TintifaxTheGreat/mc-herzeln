package main

/*
a pool of card, gives
	cards that are in the deck
	cards on the table
	cards that were already played and are now in some player's tricks
*/
type Pool struct {
	notDropped *bitmap
	onTable    *bitmap
	dropped    *bitmap
}

// factory for Pool
func NewCardpool() *Pool {
	return &Pool{
		notDropped: newBitmap(true),
		onTable:    newBitmap(false),
		dropped:    newBitmap(false),
	}
}

// deep copy
func (p *Pool) copy() *Pool {
	notDropped, onTable, dropped := *p.notDropped, *p.onTable, *p.dropped
	return &Pool{
		notDropped: &notDropped,
		onTable:    &onTable,
		dropped:    &dropped,
	}
}
