package deal

import "strconv"

// an agent (player)
type Agent struct {
	name  string
	cards *PlayersCards
}

// an interface the agent should implement
type AgentPlayer interface {
	Lead(pool *Pool, state *Gamestate) uint
	Pass(pool *Pool, state *Gamestate, lead uint) (uint, bool)
	Card() *PlayersCards
	Copy() AgentPlayer
}

type AllPlayers [PLAYERS] AgentPlayer

func (ap *AllPlayers) copy() *AllPlayers {
	result := AllPlayers{}
	for i := uint(0); i < PLAYERS; i++ {
		result[i] = ap[i].Copy()
	}
	return &result
}

func (ap *AllPlayers) update(s *Gamestate, p *Pool) {
	// TODO beautify
	ap[s.current.player].Card().hand.unset(s.current.index)
	p.OnTable.set(s.current.index)
	Info("table", p.OnTable.ToString())
	s.current.player += 1
	if s.current.player == PLAYERS {
		s.current.player = 0
	}
	s.playCount++
	if s.playCount == PLAYERS {
		// proceed to the next trick
		Info("trick end", "Trick won by player "+strconv.Itoa(int(1+s.high.player)))
		*ap[s.high.player].Card().tricks |= *p.OnTable
		*p.Dropped |= *p.OnTable
		*p.OnTable = 0
		s.lead.player = s.high.player
		s.current.player = s.high.player
		s.playCount = 0
		s.high = CardValue{}
		s.tricksCount++
	}
}
