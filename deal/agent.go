package deal

// an agent (player)
type Agent struct {
	name string
	cards *PlayersCards
}

// an interface the agent should implement
type AgentPlayer interface {
	Lead(pool *Pool, state *Gamestate) uint
	Pass(pool *Pool, state *Gamestate, lead uint) (uint, bool)
	Card() *PlayersCards
}
