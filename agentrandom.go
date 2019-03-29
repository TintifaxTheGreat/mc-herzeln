package main

type AgentRandom Agent

func NewAgentRandom(p *Pool) *AgentRandom {
	a := new(AgentRandom)
	a.pool = p
	a.cards = NewPlayersCards()
	return a
}

func (a *AgentRandom) Lead() int {
	return a.cards.hand.DrawRandom()
}

func (a *AgentRandom) Pass() int {
	return 7
}

func (a *AgentRandom) Card() *PlayersCards {
	return a.cards
}
