package main

type AgentRandom Agent

func NewAgentRandom(p *Pool) *AgentRandom {
	return &AgentRandom{
		pool:  p,
		cards: NewPlayersCards(),
	}
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
