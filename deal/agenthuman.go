package deal

import "fmt"

type AgentHuman Agent

func NewAgentHuman(p *Pool) *AgentHuman {
	return &AgentHuman{
		cards: NewPlayersCards(),
	}
}

func (a *AgentHuman) Lead(pool *Pool, state *Gamestate,) uint {
	return a.readInput()
}

func (a *AgentHuman) Pass(pool *Pool, state *Gamestate, lead uint) (uint, bool) {
	legalCards, followedSuit := a.cards.hand.legalCards(lead, true)
	fmt.Print("--->LEGAL ")
	Info(legalCards.ToString())
	index := uint(0)
	for {
		index = a.readInput()
		if legalCards.isSet(index) {
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
		if found {
			return index
		}
	}
}
