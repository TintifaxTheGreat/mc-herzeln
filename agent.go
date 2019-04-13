package main

type Agent struct {
	state Gamestate
	pool *Pool
	cards *PlayersCards
}

type AgentPlayer interface {
	Lead() uint
	Pass(lead uint) (uint, bool)
	Card() *PlayersCards
}