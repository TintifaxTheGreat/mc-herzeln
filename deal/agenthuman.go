package deal

import "fmt"

type AgentHuman Agent

func NewAgentHuman() *AgentHuman {
	return &AgentHuman{
		cards: NewPlayersCards(),
	}
}

func (a *AgentHuman) Play(_ *Pool, state *Gamestate, isLead bool, lead uint) uint {
	fmt.Println("")
	Info("your cards", a.cards.hand.ToString())
	legalCards := new(bitmap)
	if isLead {
		legalCards = state.constraintFirstLead(a.cards.hand, state.tricksCount)
	} else {
		legalCards = state.constraintPassAll(a.cards.hand, state.tricksCount, lead)
	}
	//Info("legal",legalCards.ToString())
	fmt.Println(showCardSlice(state.sTable))
	index := uint(0)
	for {
		index = a.readInput()
		if legalCards.isSet(index) {
			return index
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

// deep copy
func (a *AgentHuman) Copy() AgentPlayer {
	ar := NewAgentHuman()
	ar.cards = a.cards.copy()
	return ar
}