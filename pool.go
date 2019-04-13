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