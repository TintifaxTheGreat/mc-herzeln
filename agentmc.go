package main

type AgentMonteCarlo Agent

func NewAgentMonteCarlo(p *Pool) *AgentMonteCarlo {
	return &AgentMonteCarlo{
		pool:  p,
		cards: NewPlayersCards(),
	}
}

func (a *AgentMonteCarlo) Lead() uint {
	return a.cards.hand.DrawRandom() //TODO FIXME
}

func (a *AgentMonteCarlo) Pass(lead uint) (uint, bool) {
	legalCards, followedSuit := a.cards.hand.LegalCards(lead, true)
	return legalCards.DrawRandom(), followedSuit //TODO FIXME
}

func (a *AgentMonteCarlo) Card() *PlayersCards {
	return a.cards
}
