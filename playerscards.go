package main

type PlayersCards struct {
	hand   Bitcard
	tricks Bitcard
}

func NewPlayersCards() *PlayersCards {
	p := new(PlayersCards)
	p.hand = *NewBitcard(false)
	p.tricks = *NewBitcard(false)
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
