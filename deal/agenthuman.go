package deal

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

func (a *AgentHuman) State() *Gamestate {
	return a.state
}

func (a *AgentHuman) SetState(gamestate *Gamestate) {
	a.state = gamestate
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
