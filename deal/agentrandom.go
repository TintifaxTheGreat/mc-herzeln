package deal

type AgentRandom Agent

func NewAgentRandom() *AgentRandom {
	return &AgentRandom{
		cards: NewPlayersCards(),
	}
}

func (a *AgentRandom) Lead(_ *Pool, _ *Gamestate) uint {
	return a.cards.hand.drawRandom()
}

func (a *AgentRandom) Pass(_ *Pool, _ *Gamestate, lead uint) (uint, bool) {
	legalCards, followedSuit := a.cards.hand.legalCards(lead, true)
	return legalCards.drawRandom(), followedSuit
}

func (a *AgentRandom) Card() *PlayersCards {
	return a.cards
}

// deep copy
func (a *AgentRandom) copy() *AgentRandom {
	ar := NewAgentRandom()
	ar.cards = a.cards.copy()
	return ar
}

