package deal

type CardValue struct {
	player uint
	index  uint
}

func (cv *CardValue) value() uint {
	return cv.index % FIGURES
}

func (cv *CardValue) next() {
	cv.player += 1
	if cv.player == PLAYERS {
		cv.player = 0
	}
}

type Gamestate struct {
	tricksCount uint
	lead        CardValue
	current     CardValue
	high        CardValue
}

func (g *Gamestate) copy() *Gamestate {
	return &Gamestate{
		tricksCount: g.tricksCount,
		lead:        g.lead,
		current:     g.current,
		high:        g.high,
	}
}
