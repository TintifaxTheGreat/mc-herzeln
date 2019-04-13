package main

import (
	"math/bits"
	"math/rand"
	"strconv"
	"time"
)

type Game struct {
	players  [PLAYERS]AgentPlayer
	cardpool *Pool
}

var ALLCOLORS [COLORS]Bitmap
var CARDSTRINGS [COLORS * FIGURES]string

func NewGame() *Game {
	g := new(Game)

	// create the cardpool
	g.cardpool = NewCardpool()

	// TODO change this
	//g.players[0] = NewAgentHuman(g.cardpool)
	g.players[0] = NewAgentRandom(g.cardpool)
	g.players[1] = NewAgentRandom(g.cardpool)
	g.players[2] = NewAgentRandom(g.cardpool)
	g.players[3] = NewAgentRandom(g.cardpool)

	return g
}

func (g *Game) Start() {
	g.dealCards()
	g.Play()
	//result := g.Outcome()
	g.Outcome()
	//fmt.Println(result)
}

func (g *Game) dealCards() {
	var index uint
	for player := uint(0); player < PLAYERS; player++ {
		for i := uint(0); i < INHAND; i++ {
			index = g.cardpool.notDropped.DrawRandom()
			g.cardpool.notDropped.Unset(index)
			g.players[player].Card().hand.Set(index)
		}
	}
}

func (g *Game) Play() {
	leadPlayer := uint(0)
	for trick := uint(0); trick < INHAND; trick++ {
		info("Stich " + strconv.Itoa(int(1+trick)))

		var currentPlayer = leadPlayer
		var play, lead uint
		highest := CardValue{0, 0}
		var followedSuit bool

		info("Ausspiel Spieler " + strconv.Itoa(int(1+currentPlayer)))

		for i := uint(0); i < PLAYERS; i++ {
			info(g.players[i].Card().Show(i == leadPlayer))
		}

		for i := uint(0); i < PLAYERS; i++ {

			if i == 0 {
				// lead
				lead = g.players[currentPlayer].Lead()
				play = lead
				highest = CardValue{currentPlayer, value(lead)}
			} else {
				// pass
				play, followedSuit = g.players[currentPlayer].Pass(lead)

				if followedSuit && (value(play) < highest.value) {
					highest = CardValue{currentPlayer, value(play)}
				}
			}
			g.players[currentPlayer].Card().hand.Unset(play)
			g.cardpool.onTable.Set(play)
			info(g.cardpool.onTable.ToString())

			currentPlayer += 1
			if currentPlayer == PLAYERS {
				currentPlayer = 0
			}
		}

		*g.players[highest.player].Card().tricks |= *g.cardpool.onTable
		*g.cardpool.dropped |= *g.cardpool.onTable
		*g.cardpool.onTable = 0

		leadPlayer = highest.player
		info("Trick won by player " + strconv.Itoa(int(1+leadPlayer)))
	}
}

func (g *Game) Outcome() [PLAYERS] int {
	var points [PLAYERS] int
	for player := uint(0); player < PLAYERS; player++ {
		// count all hearts in the player's tricks
		points[player] -= bits.OnesCount64(uint64(ALLCOLORS[0] & *g.players[player].Card().tricks))
	}
	return points
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// create helpers
	helper := new(Helper)
	ALLCOLORS = helper.AllColors()
	CARDSTRINGS = helper.Cardstrings()

	for j := 0; j < 10000; j++ {
		myGame := NewGame()
		myGame.Start()
	}
}
