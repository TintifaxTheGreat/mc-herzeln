package deal

type AgentRandom Agent

func NewAgentRandom(p *Pool) *AgentRandom {
	return &AgentRandom{
		pool:  p,
		cards: NewPlayersCards(),
	}
}

func (a *AgentRandom) Lead() uint {
	return a.cards.hand.drawRandom()
}

func (a *AgentRandom) Pass(lead uint) (uint, bool) {
	legalCards, followedSuit := a.cards.hand.legalCards(lead, true)
	return legalCards.drawRandom(), followedSuit
}

func (a *AgentRandom) Card() *PlayersCards {
	return a.cards
}
