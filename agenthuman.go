package main

type AgentHuman Agent

func NewAgentHuman(p *Pool) *AgentHuman{
	a := new(AgentHuman)
	a.pool = p
	a.cards = NewPlayersCards()
	return a
}

func (a *AgentHuman) Lead() int {
	return 7
}

func (a *AgentHuman) Pass() int {
	return 7
}

func (a *AgentHuman) Card() *PlayersCards {
	return a.cards
}
