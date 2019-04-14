package main

type Pool struct {
	notDropped *Bitmap
	onTable    *Bitmap
	dropped    *Bitmap
}

func NewCardpool() *Pool {
	return &Pool{
		notDropped: NewBitmap(true),
		onTable:    NewBitmap(false),
		dropped:    NewBitmap(false),
	}
}

func (p *Pool) copy() *Pool {
	notDropped, onTable, dropped := *p.notDropped, *p.onTable, *p.dropped
	return &Pool{
		notDropped: &notDropped,
		onTable:    &onTable,
		dropped:    &dropped,
	}
}
