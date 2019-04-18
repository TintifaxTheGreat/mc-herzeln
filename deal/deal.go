package deal

import (
	"fmt"
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
	g.dealCards()
	g.play()
	return g.outcome()
}

func (g *Deal) dealCards() {
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
	for ; g.state.tricksCount < INHAND; {
		Info("Stich " + strconv.Itoa(int(1+g.state.tricksCount)))
		for i := uint(0); i < PLAYERS; i++ {
			Info(g.players[i].Card().Show(false))
		}

		//g.state.currentPlayer = g.state.leadPlayer
		//var play, lead uint
		//highest := CardValue{0, 0}
		g.state.currentHiPlayer = 0
		g.state.currentHiValue = 0
		var followedSuit bool

		for i := uint(0); i < PLAYERS; i++ {

			if g.state.currentPlayer == g.state.leadPlayer {
				// lead
				g.state.leadCard = g.players[g.state.currentPlayer].Lead(g.cardpool, g.state)
				g.state.currentCard = g.state.leadCard
				g.state.currentHiValue = value(g.state.leadCard)
				g.state.currentHiPlayer = g.state.currentPlayer
			} else {
				// pass
				g.state.currentCard, followedSuit = g.players[g.state.currentPlayer].Pass(g.cardpool, g.state, g.state.leadCard)

				if followedSuit && (value(g.state.currentCard) < g.state.currentHiValue) {
					fmt.Println("im there")
					g.state.currentHiValue = value(g.state.currentCard)
					g.state.currentHiPlayer = g.state.currentPlayer
				}
			}
			g.players[g.state.currentPlayer].Card().hand.unset(g.state.currentCard)
			g.cardpool.OnTable.set(g.state.currentCard)
			Info(g.cardpool.OnTable.ToString())

			g.state.currentPlayer += 1
			if g.state.currentPlayer == PLAYERS {
				g.state.currentPlayer = 0
			}
		}

		*g.players[g.state.currentHiPlayer].Card().tricks |= *g.cardpool.OnTable
		*g.cardpool.Dropped |= *g.cardpool.OnTable
		*g.cardpool.OnTable = 0

		g.state.tricksCount++
		g.state.leadPlayer = g.state.currentHiPlayer
		g.state.currentPlayer = g.state.currentHiPlayer
		Info("Trick won by player " + strconv.Itoa(int(1+g.state.leadPlayer)))
	}
}

func (g *Deal) outcome() [PLAYERS] int {
	var points [PLAYERS] int
	for player := uint(0); player < PLAYERS; player++ {
		// count all hearts in the player's tricks
		points[player] -= bits.OnesCount64(uint64(ALLCOLORS[0] & *g.players[player].Card().tricks))
	}
	return points
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

