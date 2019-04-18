package deal

type AgentMonteCarlo Agent

func NewAgentMonteCarlo(p *Pool) *AgentMonteCarlo {
	return &AgentMonteCarlo{
		cards: NewPlayersCards(),
	}
}

func (a *AgentMonteCarlo) Lead(pool *Pool, state *Gamestate) uint {
	// if there is only one card left, play this card
	if state.tricksCount == INHAND {
		return a.cards.hand.next(0)
	}

	tPool := &Pool{}
	tState := state

	aBuddy := NewAgentRandom(tPool)
	bBuddy := NewAgentRandom(tPool)
	cBuddy := NewAgentRandom(tPool)
	dBuddy := NewAgentRandom(tPool)

	cardsOfOtherPlayers := *pool.NotDropped &^ *a.cards.hand

	foo := uint(0) // TODO FIXME
	for {
		// distribute remaining cards to other players
		foo = cardsOfOtherPlayers.drawRandom()
		tPool.NotDropped.unset(foo)
		bBuddy.cards.hand.set(foo)

		foo = cardsOfOtherPlayers.drawRandom()
		tPool.NotDropped.unset(foo)
		cBuddy.cards.hand.set(foo)

		foo = cardsOfOtherPlayers.drawRandom()
		tPool.NotDropped.unset(foo)
		dBuddy.cards.hand.set(foo)

		index := uint(0)
		for i := state.tricksCount; i < INHAND; i++ {
			index = a.cards.hand.next(index)

			// update game state
			tState.tricksCount--
			tState.currentPlayer += 1
			if tState.currentPlayer == PLAYERS {
				tState.currentPlayer = 0
			}

			playout := NewDeal(tPool, tState, [PLAYERS]AgentPlayer{aBuddy, bBuddy, cBuddy, dBuddy,})
			playout.Play()
		}


		//TODO Fixme
		break
	}
	// TODO Fixme
	return 7
}

func (a *AgentMonteCarlo) Pass(pool *Pool, state *Gamestate, lead uint) (uint, bool) {
	legalCards, followedSuit := a.cards.hand.legalCards(lead, true)
	return legalCards.drawRandom(), followedSuit //TODO FIXME
}

func (a *AgentMonteCarlo) Card() *PlayersCards {
	return a.cards
}
