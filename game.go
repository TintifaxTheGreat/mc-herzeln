package main

import (
	"math/rand"
	"strconv"
	"time"
)

type Game struct {
	players      [PLAYERS]AgentPlayer
	cardpool     *Pool
}

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

		// lead
		play := g.players[currentPlayer].Lead()
		//leadPlay := play
		highest.player = currentPlayer
		highest.value = value(play)

		// pass

		for i := 1; i < PLAYERS; i++ {
			play := g.players[currentPlayer].Pass()

			// TODO Change this later
			//if followed_suit and play[1] < highest[1]:
			if

			/*
			                    play, followed_suit = self.myAgent[cur_player](cur_player, False, lead_play)

			                    # the highest card is the one with the smallest(!) index
			                    if followed_suit and play[1] < highest[1]:
			                        highest = cur_player, play[1]
			 */


		}

	}
}

func main() {
	for j := 0; j < 1; j++ {
		myGame := NewGame()
		myGame.Start()
		myGame.Play()
	}
}
