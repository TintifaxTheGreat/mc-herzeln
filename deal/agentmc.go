package deal

import "fmt"

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

	rIndex := uint(0) // TODO FIXME add channel stuff
	for {
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

		fmt.Println("............")
		for player := uint(0); player < PLAYERS; player++ {
			Info(buddies[player].cards.Show(false))
		}
		fmt.Println("............")

		index := uint(0)
		for i := state.tricksCount; i < INHAND; i++ {
			index = a.cards.hand.next(index)

			// update game state
			state.current.next()

			PlayersCopy := [PLAYERS]AgentPlayer{}
			for player := uint(0); player < PLAYERS; player++ {
				PlayersCopy[player] = buddies[player].copy()
			}

			fmt.Println("mmmmmmmmmmmm")
			for player := uint(0); player < PLAYERS; player++ {
				Info(PlayersCopy[player].Card().Show(false))
			}
			fmt.Println("mmmmmmmmmmmm")

			playout := NewDeal(pool.copy(), state.copy(), PlayersCopy)
			playout.Play()
			acc[index] += playout.playerOutcome(state.current.player)
			if acc[maxKey] < acc[index] {
				maxKey = index
			}
		}

		//TODO Fixme
		break
	}
	return maxKey
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
