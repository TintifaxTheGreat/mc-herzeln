package main

type AgentHuman Agent

func NewAgentHuman(p *Pool) *AgentHuman{
	return &AgentHuman{
		pool:  p,
		cards: NewPlayersCards(),
	}
}

func (a *AgentHuman) Lead() int {
	return 7
}

func (a *AgentHuman) Pass(lead int) (int, bool) {
	return 7, false
}

func (a *AgentHuman) Card() *PlayersCards {
	return a.cards
}