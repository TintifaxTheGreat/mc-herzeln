package main

type Agent struct {
	state Gamestate
	pool *Pool
	cards *PlayersCards
}

type AgentPlayer interface {
	Lead() int
	Pass(lead int) (int, bool)
	Card() *PlayersCards
}