package main

type AgentRandom Agent

func NewAgentRandom(p *Pool) *AgentRandom {
	return &AgentRandom{
		pool:  p,
		cards: NewPlayersCards(),
	}
}

func (a *AgentRandom) Lead() uint {
	return a.cards.hand.DrawRandom()
}

func (a *AgentRandom) Pass(lead uint) (uint, bool) {
	legalCards, followedSuit := a.cards.hand.LegalCards(lead, true)
	return legalCards.DrawRandom(), followedSuit
}

func (a *AgentRandom) Card() *PlayersCards {
	return a.cards
}
