package deal

type CardValue struct {
	player uint
	index  uint
}

func (cv *CardValue) value() uint {
	return cv.index % FIGURES
}

type Gamestate struct {
	tricksCount         uint
	playCount           uint
	sTable              []uint
	lead                CardValue
	current             CardValue
	high                CardValue
	constraintFirstLead constraintLead
	constraintPassAll   constraintPass
	goal                goal
}

// factory for Gamestate
func NewGamestate(cfl constraintLead, cpa constraintPass, goal goal) *Gamestate {
	return &Gamestate{
		constraintFirstLead: cfl,
		constraintPassAll:   cpa,
		goal:                goal,
	}
}

func (s *Gamestate) copy() *Gamestate {
	return &Gamestate{
		tricksCount:         s.tricksCount,
		playCount:           s.playCount,
		lead:                s.lead,
		current:             s.current,
		high:                s.high,
		constraintFirstLead: s.constraintFirstLead,
		constraintPassAll:   s.constraintPassAll,
		goal:                s.goal,
	}
}
