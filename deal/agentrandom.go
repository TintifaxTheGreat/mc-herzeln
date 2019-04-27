package deal

type AgentRandom Agent

func NewAgentRandom() *AgentRandom {
	return &AgentRandom{
		cards: NewPlayersCards(),
	}
}

func (a *AgentRandom) Play(_ *Pool, state *Gamestate, isLead bool, lead uint) uint {
	legalCards := new(bitmap)
	if isLead {
		legalCards = state.constraintFirstLead(a.cards.hand, state.tricksCount)
	} else {
		legalCards = state.constraintPassAll(a.cards.hand, state.tricksCount, lead)
	}
	return legalCards.drawRandom()
}

func (a *AgentRandom) Card() *PlayersCards {
	return a.cards
}

// deep copy
func (a *AgentRandom) Copy() AgentPlayer {
	ar := NewAgentRandom()
	ar.cards = a.cards.copy()
	return ar
}
