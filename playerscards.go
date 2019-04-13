package main

type PlayersCards struct {
	hand   *Bitmap
	tricks *Bitmap
}

func NewPlayersCards() *PlayersCards {
	p := new(PlayersCards)
	p.hand = NewBitmap(false)
	p.tricks = NewBitmap(false)
	return p
}

func (p *PlayersCards) Show(leadplayer bool) string {
	var s string = ""
	s += p.hand.ToString() + "[" +
		p.tricks.ToString() + "]"
	if leadplayer == true {
		s += "*"
	}
	return s
}
