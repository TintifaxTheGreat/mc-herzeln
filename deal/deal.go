package deal

import (
	"math/rand"
	"strconv"
	"time"
)

type Deal struct {
	players             AllPlayers
	cardpool            *Pool
	state               *Gamestate
	constraintFirstLead constraintFirstLead
	constraintPassAll   constraintPassAll
	goal                goal
}

var ALLCOLORS [COLORS]bitmap
var ALLFIGURES [FIGURES]bitmap
var CARDSTRINGS [COLORS * FIGURES]string

func NewDeal(
	pool *Pool,
	state *Gamestate,
	agents AllPlayers,
	cfl constraintFirstLead,
	cpa constraintPassAll,
	goal goal,
) *Deal {
	return &Deal{
		cardpool:            pool,
		players:             agents,
		state:               state,
		constraintFirstLead: cfl,
		constraintPassAll:   cpa,
		goal:                goal,
	}
}

func (g *Deal) Play() [PLAYERS] int {
	g.play()
	return g.outcome()
}

func (g *Deal) DealCards() {
	Info("dealing cards", g.cardpool.NotDropped.ToString())
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
		if g.state.playCount == 0 {
			for i := uint(0); i < PLAYERS; i++ {
				Info("trick "+strconv.Itoa(int(1+g.state.tricksCount)), g.players[i].Card().Show(false))
			}
			// lead
			// TODO lead hearts in first trick is not allowed
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
		g.players.update(g.state, g.cardpool)
	}
}

// the outcome of the game
func (g *Deal) outcome() [PLAYERS] int {
	return g.goal(g.players)
}

// the outcome of the game from the perspective of a player
func (g *Deal) playerOutcome(player uint) int {
	result := 0
	outcome := g.outcome()
	for i := uint(0); i < PLAYERS; i++ {
		if i == player {
			result += outcome[i]
		} else {
			result -= outcome[i]
		}
	}
	//fmt.Println(player,result)
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
