package deal

type CardValue struct {
	player uint
	index  uint
}

func (cv *CardValue) value() uint {
	return cv.index % FIGURES
}

type Gamestate struct {
	tricksCount uint
	playCount   uint
	lead        CardValue
	current     CardValue
	high        CardValue
}

func (g *Gamestate) copy() *Gamestate {
	return &Gamestate{
		tricksCount: g.tricksCount,
		playCount:   g.playCount,
		lead:        g.lead,
		current:     g.current,
		high:        g.high,
	}
}
