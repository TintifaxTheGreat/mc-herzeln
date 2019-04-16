package deal

/*
a pool of card, gives
	cards that are in the deck
	cards on the table
	cards that were already played and are now in some player's tricks
*/
type Pool struct {
	NotDropped *bitmap
	OnTable    *bitmap
	Dropped    *bitmap
}

// factory for Pool
func NewCardpool() *Pool {
	return &Pool{
		NotDropped: newBitmap(true),
		OnTable:    newBitmap(false),
		Dropped:    newBitmap(false),
	}
}

// deep copy
func (p *Pool) copy() *Pool {
	notDropped, onTable, dropped := *p.NotDropped, *p.OnTable, *p.Dropped
	return &Pool{
		NotDropped: &notDropped,
		OnTable:    &onTable,
		Dropped:    &dropped,
	}
}
