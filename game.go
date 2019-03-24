package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Agent uint8

const (
	AGENT_HUMAN Agent = iota + 1
	AGENT_RANDOM
	AGENT_MCTS
)

type Game struct {
	players      int
	colors       int
	figures      int
	cardsInHand  int
	agent        [PLAYERS]Agent
	cardpool     *Cardpool
	playerscards *PlayersCards
}

func NewGame() *Game {
	rand.Seed(time.Now().UnixNano())
	g := new(Game)
	g.cardsInHand = 8
	g.agent = [4]Agent{AGENT_HUMAN, AGENT_RANDOM, AGENT_RANDOM, AGENT_RANDOM}

	// create the cardpool
	g.cardpool = NewCardpool()

	// create the players cards
	g.playerscards = NewPlayersCards()

	// TODO create helpers

	// TODO add the agents stuff here

	return g
}

func (g *Game) Start() {
	g.dealCards()
}

func (g *Game) dealCards() {
	for player := 0; player < PLAYERS; player++ {
		for i := 0; i < INHAND; i++ {
			g.cardpool.notDropped.DrawRandom(&g.playerscards.c[player])
		}

	}
}

func main() {
	for j := 0; j < 1; j++ {
		myGame := NewGame()
		myGame.Start()
		fmt.Println(myGame.playerscards.c[0].ToString(CARDSTRINGS))
		fmt.Println(myGame.playerscards.c[1].ToString(CARDSTRINGS))
		fmt.Println(myGame.playerscards.c[2].ToString(CARDSTRINGS))
		fmt.Println(myGame.playerscards.c[3].ToString(CARDSTRINGS))
	}
}
