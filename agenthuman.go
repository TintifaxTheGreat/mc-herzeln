package main

import "fmt"

type AgentHuman Agent

func NewAgentHuman(p *Pool) *AgentHuman {
	return &AgentHuman{
		pool:  p,
		cards: NewPlayersCards(),
	}
}

func (a *AgentHuman) Lead() uint {
	return a.readInput()
}

func (a *AgentHuman) Pass(lead uint) (uint, bool) {
	legalCards, followedSuit := a.cards.hand.LegalCards(lead, true)
	fmt.Print("--->LEGAL ")
	info(legalCards.ToString())
	index := uint(0)
	for {
		index = a.readInput()
		if legalCards.IsSet(index) {
			return index, followedSuit
		}
	}
}

func (a *AgentHuman) Card() *PlayersCards {
	return a.cards
}

func (a *AgentHuman) readInput() uint {
	var input string
	for {
		fmt.Print("IHRE EINGABE: ")
		fmt.Scanln(&input)
		index, found := IndexOfCard(input)
		if found{
			return index
		}
	}
}
