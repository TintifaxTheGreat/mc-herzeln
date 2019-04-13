package main

import (
	"math/rand"
	"strconv"
	"time"
)

type Game struct {
	players  [PLAYERS]AgentPlayer
	cardpool *Pool
}

var ALLCOLORS [COLORS]Bitmap

func NewGame() *Game {
	rand.Seed(time.Now().UnixNano())
	g := new(Game)

	// create the cardpool
	g.cardpool = NewCardpool()

	// TODO change this
	g.players[0] = NewAgentHuman(g.cardpool)
	//g.players[0] = NewAgentRandom(g.cardpool)
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

		info("Ausspiel Spieler " + strconv.Itoa(int(1+ currentPlayer)))

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

func main() {
	for j := 0; j < 1; j++ {
		myGame := NewGame()
		myGame.Start()
		myGame.Play()
	}
}
