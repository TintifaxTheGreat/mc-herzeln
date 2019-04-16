package main

// an agent (player)
type Agent struct {
	state Gamestate
	pool *Pool
	cards *PlayersCards
}

// an interface the agent should implement
type AgentPlayer interface {
	Lead() uint
	Pass(lead uint) (uint, bool)
	Card() *PlayersCards
}