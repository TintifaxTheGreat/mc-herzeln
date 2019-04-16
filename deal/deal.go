package main

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

func NewGame(pool *Pool, agents [PLAYERS] AgentPlayer) *Deal {
	return &Deal{
		cardpool: pool,
		players:  agents,
		state:    &Gamestate{},
	}
}

func (g *Deal) Play() [PLAYERS] int {
	g.dealCards()
	g.play()
	return g.outcome()
}

func (g *Deal) dealCards() {
	info(g.cardpool.notDropped.toString())
	var index uint
	for player := uint(0); player < PLAYERS; player++ {
		for i := uint(0); i < INHAND; i++ {
			index = g.cardpool.notDropped.drawRandom()
			g.cardpool.notDropped.unset(index)
			g.players[player].Card().hand.set(index)
		}
	}
}

func (g *Deal) play() {
	// TODO Fixme
	for ; g.state.tricksCount < INHAND; {
		info("Stich " + strconv.Itoa(int(1+g.state.tricksCount)))

		g.state.currentPlayer = g.state.leadPlayer
		var play, lead uint
		highest := CardValue{0, 0}
		var followedSuit bool

		info("Ausspiel Spieler " + strconv.Itoa(int(1+g.state.currentPlayer)))

		for i := uint(0); i < PLAYERS; i++ {
			info(g.players[i].Card().Show(i == g.state.leadPlayer))
		}

		for i := uint(0); i < PLAYERS; i++ {

			if i == 0 {
				// lead
				lead = g.players[g.state.currentPlayer].Lead()
				play = lead
				highest = CardValue{g.state.currentPlayer, value(lead)}
			} else {
				// pass
				play, followedSuit = g.players[g.state.currentPlayer].Pass(lead)

				if followedSuit && (value(play) < highest.value) {
					highest = CardValue{g.state.currentPlayer, value(play)}
				}
			}
			g.players[g.state.currentPlayer].Card().hand.unset(play)
			g.cardpool.onTable.set(play)
			info(g.cardpool.onTable.toString())

			g.state.currentPlayer += 1
			if g.state.currentPlayer == PLAYERS {
				g.state.currentPlayer = 0
			}
		}

		*g.players[highest.player].Card().tricks |= *g.cardpool.onTable
		*g.cardpool.dropped |= *g.cardpool.onTable
		*g.cardpool.onTable = 0

		g.state.tricksCount++
		g.state.leadPlayer = highest.player
		info("Trick won by player " + strconv.Itoa(int(1+g.state.leadPlayer)))
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

func main() {

	for j := 0; j < 1; j++ {
		// create cardpool
		cardpool := NewCardpool()
		//gamestate := &Gamestate{}

		// create agents
		agents := [PLAYERS] AgentPlayer{
			NewAgentRandom(cardpool),
			NewAgentRandom(cardpool),
			NewAgentRandom(cardpool),
			NewAgentRandom(cardpool),
		}

		myGame := NewGame(cardpool, agents)
		result := myGame.Play()
		fmt.Println(result)
	}
}
