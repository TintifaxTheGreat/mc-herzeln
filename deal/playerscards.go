package deal

// a player's hand, and tricks
type PlayersCards struct {
	hand   *bitmap
	tricks *bitmap
}

// factory for PlayersCards
func NewPlayersCards() *PlayersCards {
	p := &PlayersCards{}
	p.hand = newBitmap(false)
	p.tricks = newBitmap(false)
	return p
}

// deep copy
func (p *PlayersCards) copy() *PlayersCards {
	hand, tricks := *p.hand, *p.tricks

	return &PlayersCards{
		hand:   &hand,
		tricks: &tricks,
	}
}

// string representation of a player's hand, and tricks
func (p *PlayersCards) Show(leadplayer bool) string {
	var s string = ""
	s += p.hand.ToString() + "[" +
		p.tricks.ToString() + "]"
	if leadplayer == true {
		s += "*"
	}
	return s
}
