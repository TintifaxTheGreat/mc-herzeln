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
	Pass(pool *Pool, state *Gamestate, lead uint) uint
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
	// move card from current player's hand onto the table
	ap[s.current.player].Card().hand.unset(s.current.index)
	p.OnTable.set(s.current.index)
	Info("table", p.OnTable.ToString())

	// proceed to the next player in the round
	s.current.player += 1
	if s.current.player == PLAYERS {
		s.current.player = 0
	}

	// increase play count
	s.playCount++
	if s.playCount == PLAYERS {
		// proceed to the next trick
		Info("trick end", "Trick won by player "+strconv.Itoa(int(1+s.high.player)))

		// move the cards from the table into the trick winner's tricks
		// also update the list of the already dropped cards
		*ap[s.high.player].Card().tricks |= *p.OnTable
		*p.Dropped |= *p.OnTable
		*p.OnTable = 0

		// the lead of the next play moves to the winner of the trick
		s.lead.player = s.high.player
		s.current.player = s.high.player

		// set play cound to zero again; increase tricks count
		s.playCount = 0
		s.high = CardValue{}
		s.tricksCount++
	}
}
