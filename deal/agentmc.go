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

	// limit execution time of calculation
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	return a.playouts(ctx, pool, state)
}

func (a *AgentMonteCarlo) playouts(ctx context.Context, pool *Pool, state *Gamestate) uint {
	maxKey := uint(0)
	acc := make(map[uint]int)

	innerMaxKey := uint(0)
	innerAcc := make(map[uint]int)

	buddies := AllPlayers{}
	for player := uint(0); player < PLAYERS; player++ {
		buddies[player] = NewAgentRandom()
	}
	buddies[state.current.player].Card().tricks = a.cards.tricks
	buddies[state.current.player].Card().hand = a.cards.hand

	fmt.Println(a.cards.hand.ToString())

	// calculate the hidden cards
	hiddenCards := newBitmap(true)
	*hiddenCards &^= *a.cards.hand
	*hiddenCards &^= *pool.Dropped
	*hiddenCards &^= *pool.OnTable

	fmt.Println(hiddenCards.ToString())

	rIndex := uint(0)
	count := uint(0)
	for {
		select {
		case <-ctx.Done():
			maxKey = 999
			acc[maxKey] = -9999999 // TODO FIxme
			for key, value := range acc {
				if value > acc[maxKey] {
					maxKey = key
				}
			}
			fmt.Print("Games played: ")
			fmt.Println(count)
			fmt.Println(acc)
			fmt.Println(maxKey)
			return maxKey
		default:
			// copy everything
			tState := state.copy()
			tPool := pool.copy()
			tBuddies := buddies.copy()
			tHiddenCards := newBitmap(false)
			*tHiddenCards = *hiddenCards

			thisPlayer := tState.current.player

			// distribute remaining cards to other players
			for i := tState.tricksCount; i < INHAND; i++ {
				for player := uint(0); player < PLAYERS; player++ {
					if player != tState.current.player {
						rIndex = tHiddenCards.drawRandom()
						tHiddenCards.unset(rIndex)
						tBuddies[player].Card().hand.set(rIndex)
					}
				}
			}
			innerAcc = make(map[uint]int)
			index := uint(0)
			for i := tState.tricksCount; i < INHAND; i++ {

				// copy everything
				tState2 := tState.copy()
				tPool2 := tPool.copy()
				tBuddies2 := tBuddies.copy()

				// index = a.cards.hand.next(index)
				index = tBuddies2[tState2.current.player].Card().hand.next(index)

				tBuddies2[tState2.current.player].Card().hand.unset(index)
				tPool2.OnTable.set(index)

				tState2.lead.index = index
				tState2.current.index = tState2.lead.index
				tState2.high = CardValue{
					index:  tState2.lead.index,
					player: tState2.current.player,
				}

				tBuddies2.update(tState2, tPool2)

				playout := NewDeal(tPool2, tState2, *tBuddies2, ConstraintLeadAny, ConstraintPassAny, GoalNoQueens)
				playout.Play()

				count++
				innerAcc[index] += playout.playerOutcome(thisPlayer) //TODO Fixme
			}
			innerMaxKey = 999
			innerAcc[innerMaxKey] = -99999 // TODO FIxme
			for key, value := range innerAcc {
				if value > innerAcc[innerMaxKey] {
					innerMaxKey = key
				}
			}
			acc[innerMaxKey]++
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
func (a *AgentMonteCarlo) Copy() AgentPlayer {
	ar := NewAgentMonteCarlo()
	ar.cards = a.cards.copy()
	return ar
}
