package main

import (
	"github.com/yourbasic/bit"
	"math/rand"
	"strconv"
	"time"
)

type Game struct {
	players  [PLAYERS]AgentPlayer
	cardpool *Pool
}

var ALLCOLORS [COLORS]*bit.Set

func NewGame() *Game {
	rand.Seed(time.Now().UnixNano())
	g := new(Game)

	// create the cardpool
	g.cardpool = NewCardpool()

	// TODO change this
	//g.players[0] = NewAgentHuman(g.cardpool)
	g.players[0] = NewAgentRandom(g.cardpool)
	g.players[1] = NewAgentRandom(g.cardpool)
	g.players[2] = NewAgentRandom(g.cardpool)
	g.players[3] = NewAgentRandom(g.cardpool)

	// TODO create helpers
	ALLCOLORS = allcolors()

	// TODO add the agents stuff here

	return g
}

func (g *Game) Start() {
	g.dealCards()
}

func (g *Game) dealCards() {
	for player := 0; player < PLAYERS; player++ {
		for i := 0; i < INHAND; i++ {
			g.players[player].Card().hand.Set(g.cardpool.notDropped.DrawRandom())
		}
	}
}

func (g *Game) Play() {
	var leadPlayer int = 0
	for trick := 0; trick < INHAND; trick++ {
		info("Stich " + strconv.Itoa(trick+1))

		for player := 0; player < PLAYERS; player++ {
			info(g.players[player].Card().Show(player == leadPlayer))
		}

		var currentPlayer = leadPlayer
		highest := CardValue{-1, -999}

		info("Ausspiel Spieler " + strconv.Itoa(1 + currentPlayer))

		// lead
		lead := g.players[currentPlayer].Lead()
		highest = CardValue{currentPlayer, value(lead)}
		g.cardpool.notDropped.Unset(lead)
		g.cardpool.onTable.Set(lead)

		// pass
		for i := 1; i < PLAYERS; i++ {
			pass, followedSuit := g.players[currentPlayer].Pass(lead)

			if followedSuit && (value(lead) < highest.value) {
				highest = CardValue{currentPlayer, value(pass)}
			}
			g.cardpool.notDropped.Unset(lead)
			g.cardpool.onTable.Set(lead)
		}
		g.players[highest.player].Card().tricks.c.Or(&g.cardpool.onTable.c)
		g.cardpool.onTable = *NewBitcard(true) // TODO ???

		leadPlayer = highest.player
		info("Trick won by player " + strconv.Itoa(1+ leadPlayer))
	}
}

func main() {
	for j := 0; j < 1; j++ {
		myGame := NewGame()
		myGame.Start()
		myGame.Play()
	}
}
