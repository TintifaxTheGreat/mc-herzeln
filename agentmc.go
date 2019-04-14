package main

import (
	"math/bits"
)

type AgentMonteCarlo Agent

func NewAgentMonteCarlo(p *Pool) *AgentMonteCarlo {
	return &AgentMonteCarlo{
		pool:  p,
		cards: NewPlayersCards(),
	}
}

func (a *AgentMonteCarlo) Lead() uint {
	countHand := bits.OnesCount64(uint64(*a.cards.hand))
	if countHand == 1 {
		return a.cards.hand.Next(0)
	}
	//tPool := &Pool{}
	//tCards := &PlayersCards{}
	index := uint(0)
	for i := uint(0); i < 1; i++ {
		for i := 0; i < countHand; i++ {
			index = a.cards.hand.Next(index)
			//tPool = a.pool.copy()
			//tCards = a.cards.copy()

			//assign cards from pool to the other players

			//playout game

			// TODO add stuff here
		}
	}
	return 7
}

func (a *AgentMonteCarlo) Pass(lead uint) (uint, bool) {
	legalCards, followedSuit := a.cards.hand.LegalCards(lead, true)
	return legalCards.DrawRandom(), followedSuit //TODO FIXME
}

func (a *AgentMonteCarlo) Card() *PlayersCards {
	return a.cards
}
