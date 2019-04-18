package deal

import (
	"math/bits"
	"math/rand"
	"strconv"
	"time"
)

type Deal struct {
	players  [PLAYERS] AgentPlayer
	cardpool *Pool
	state    *Gamestate
}

var ALLCOLORS [COLORS]bitmap
var ALLFIGURES [FIGURES]bitmap
var CARDSTRINGS [COLORS * FIGURES]string

func NewDeal(pool *Pool, state *Gamestate, agents [PLAYERS] AgentPlayer) *Deal {
	gamestate := state
	return &Deal{
		cardpool: pool,
		players:  agents,
		state:    gamestate,
	}
}

func (g *Deal) Play() [PLAYERS] int {
	g.play()
	return g.outcome()
}

func (g *Deal) DealCards() {
	Info(g.cardpool.NotDropped.ToString())
	var index uint
	for player := uint(0); player < PLAYERS; player++ {
		for i := uint(0); i < INHAND; i++ {
			index = g.cardpool.NotDropped.drawRandom()
			g.cardpool.NotDropped.unset(index)
			g.players[player].Card().hand.set(index)
		}
	}
}

func (g *Deal) play() {
	var followedSuit bool
	for ; g.state.tricksCount < INHAND; {
		Info("Stich " + strconv.Itoa(int(1+g.state.tricksCount)))
		for i := uint(0); i < PLAYERS; i++ {
			Info(g.players[i].Card().Show(false))
		}

		g.state.high = CardValue{}
		for i := uint(0); i < PLAYERS; i++ {

			if g.state.current.player == g.state.lead.player {
				// lead
				g.state.lead.index = g.players[g.state.current.player].Lead(g.cardpool, g.state)
				g.state.current.index = g.state.lead.index
				g.state.high = CardValue{
					index:  g.state.lead.index,
					player: g.state.current.player,
				}
			} else {
				// pass
				g.state.current.index, followedSuit = g.players[g.state.current.player].Pass(g.cardpool, g.state, g.state.lead.index)

				if followedSuit && (g.state.current.value() < g.state.high.value()) {
					g.state.high = g.state.current
				}
			}
			g.players[g.state.current.player].Card().hand.unset(g.state.current.index)
			g.cardpool.OnTable.set(g.state.current.index)
			Info(g.cardpool.OnTable.ToString())

			g.state.current.next()
		}

		*g.players[g.state.high.player].Card().tricks |= *g.cardpool.OnTable
		*g.cardpool.Dropped |= *g.cardpool.OnTable
		*g.cardpool.OnTable = 0

		g.state.tricksCount++
		g.state.lead.player = g.state.high.player
		g.state.current.player = g.state.high.player
		Info("Trick won by player " + strconv.Itoa(int(1+g.state.lead.player)))
	}
}

// the outcome of the game
func (g *Deal) outcome() [PLAYERS] int {
	var points [PLAYERS] int
	for player := uint(0); player < PLAYERS; player++ {
		// count all hearts in the player's tricks
		points[player] -= bits.OnesCount64(uint64(ALLCOLORS[0] & *g.players[player].Card().tricks))
	}
	return points
}

// the outcome of the game from the perspective of a player
func (g *Deal) playerOutcome(player uint) int {
	result := 0
	outcome := g.outcome()
	for i := uint(0); i < PLAYERS; i++ {
		if i == player {
			result += outcome[i]
			continue
		}
		result -= outcome[i]
	}
	return result
}

func init() {
	// create helpers
	helper := new(Helper)
	CARDSTRINGS = helper.Cardstrings()
	ALLCOLORS = helper.AllColors()
	ALLFIGURES = helper.AllFigures()

	// misc
	rand.Seed(time.Now().UnixNano())
}
