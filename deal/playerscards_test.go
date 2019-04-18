package deal

import (
	"testing"
)

func TestCopy(t *testing.T) {
	cards := NewPlayersCards()
	cards.hand.set(24)
	cards.hand.set(36)
	cards.hand.set(21)

	cards.tricks.set(16)
	cards.tricks.set(47)

	cardsCopy := cards.copy()

	if *cards.hand != *cardsCopy.hand {
		t.Errorf("TestCopy was incorrect, expected copy to be equal.")
	}

	if *cards.tricks != *cardsCopy.tricks {
		t.Errorf("TestCopy was incorrect, expected copy to be equal.")
	}
}

