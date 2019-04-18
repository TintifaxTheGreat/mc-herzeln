package deal

type AgentMonteCarlo Agent

func NewAgentMonteCarlo(p *Pool) *AgentMonteCarlo {
	return &AgentMonteCarlo{
		pool:  p,
		cards: NewPlayersCards(),
	}
}

func (a *AgentMonteCarlo) Lead() uint {
	tricks := a.state.tricksCount
	if tricks == INHAND {
		return a.cards.hand.next(0)
	}
	/*
	tPool := &Pool{}
	tCards := &PlayersCards{}
	index := uint(0)



	// TODO infinite loop
	for i := uint(0); i < 1; i++ {
		for i := unit(tricks); i <  INHAND; i++ {
			index = a.cards.hand.next(index)
			tPool = a.pool.copy()
			tCards = a.cards.copy()

			for

			//assign cards from pool to the other players

			//playout game

			// TODO add stuff here
		}
	}
	*/
	return 7
}

func (a *AgentMonteCarlo) Pass(lead uint) (uint, bool) {
	legalCards, followedSuit := a.cards.hand.legalCards(lead, true)
	return legalCards.drawRandom(), followedSuit //TODO FIXME
}

func (a *AgentMonteCarlo) Card() *PlayersCards {
	return a.cards
}

func (a *AgentMonteCarlo) State() *Gamestate {
	return a.state
}

func (a *AgentMonteCarlo) SetState(gamestate *Gamestate) {
	a.state = gamestate
}

