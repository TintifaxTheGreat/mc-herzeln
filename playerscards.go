package main

type PlayersCards struct {
	c [COLORS]Bitcard
}

func NewPlayersCards() *PlayersCards {
	p := new(PlayersCards)
	for color := 0; color < COLORS; color++ {
		p.c[color] = *NewBitcard(false)
	}
	return p
}
