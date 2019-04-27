package deal

type AgentRandom Agent

func NewAgentRandom() *AgentRandom {
	return &AgentRandom{
		cards: NewPlayersCards(),
	}
}

func (a *AgentRandom) Lead(_ *Pool, state *Gamestate) uint {
	legalCards := state.constraintFirstLead(a.cards.hand, state.tricksCount)
	return legalCards.drawRandom()
}

func (a *AgentRandom) Pass(_ *Pool, state *Gamestate, lead uint) uint {
	legalCards := state.constraintPassAll(a.cards.hand, state.tricksCount, lead)
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
