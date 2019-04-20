package deal

import (
	"context"
	"fmt"
	"time"
)

type AgentMonteCarlo Agent

func NewAgentMonteCarlo() *AgentMonteCarlo {
	return &AgentMonteCarlo{
		cards: NewPlayersCards(),
	}
}

func (a *AgentMonteCarlo) Lead(pool *Pool, state *Gamestate) uint {
	// if there is only one card left, play this card
	if state.tricksCount == INHAND {
		return a.cards.hand.next(0)
	}

	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	return a.playouts(ctx, pool, state)
}

func (a *AgentMonteCarlo) playouts(ctx context.Context, pool *Pool, state *Gamestate) uint {
	acc := make(map[uint]int)
	maxKey := uint(0)

	buddies := [PLAYERS] *AgentRandom{}
	for player := uint(0); player < PLAYERS; player++ {
		buddies[player] = NewAgentRandom()
	}
	buddies[state.current.player].cards.tricks = a.cards.tricks
	buddies[state.current.player].cards.hand = a.cards.hand

	// calculate the hidden cards
	hiddenCards := newBitmap(true)
	*hiddenCards &^= *a.cards.hand
	*hiddenCards &^= *pool.Dropped
	*hiddenCards &^= *pool.OnTable

	rIndex := uint(0)
	count := uint(0)
	for {
		select {
		case <-ctx.Done():
			fmt.Print("Games played: ")
			fmt.Println(count)
			return maxKey
		default:
			// distribute remaining cards to other players
			for i := state.tricksCount; i < INHAND; i++ {
				for player := uint(0); player < PLAYERS; player++ {
					if player != state.current.player {
						rIndex = hiddenCards.drawRandom()
						hiddenCards.unset(rIndex)
						buddies[player].cards.hand.set(rIndex)
					}
				}
			}


			for player := uint(0); player < PLAYERS; player++ {
				Info(buddies[player].cards.Show(false))
			}


			index := uint(0)
			for i := state.tricksCount; i < INHAND; i++ {
				index = a.cards.hand.next(index)

				// update game state
				state.current.next()

				PlayersCopy := [PLAYERS]AgentPlayer{}
				for player := uint(0); player < PLAYERS; player++ {
					PlayersCopy[player] = buddies[player].copy()
				}


				for player := uint(0); player < PLAYERS; player++ {
					Info(PlayersCopy[player].Card().Show(false))
				}


				playout := NewDeal(pool.copy(), state.copy(), PlayersCopy)
				playout.Play()
				count++
				acc[index] += playout.playerOutcome(state.current.player)
				if acc[maxKey] < acc[index] {
					maxKey = index
				}
			}
		}
	}
}

func (a *AgentMonteCarlo) Pass(pool *Pool, state *Gamestate, lead uint) (uint, bool) {
	legalCards, followedSuit := a.cards.hand.legalCards(lead, true)
	return legalCards.drawRandom(), followedSuit //TODO FIXME
}

func (a *AgentMonteCarlo) Card() *PlayersCards {
	return a.cards
}

// deep copy
func (a *AgentMonteCarlo) copy() *AgentMonteCarlo {
	ar := NewAgentMonteCarlo()
	ar.cards = a.cards.copy()
	return ar
}
