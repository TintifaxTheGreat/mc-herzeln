package deal

import "math/bits"

type goal func(players AllPlayers) [PLAYERS] int

var GoalHearts = goal(func(players AllPlayers) [PLAYERS] int {
	var points [PLAYERS] int
	for player := uint(0); player < PLAYERS; player++ {
		// count all hearts in the player's tricks
		points[player] -= bits.OnesCount64(uint64(ALLCOLORS[0] & *players[player].Card().tricks))
	}
	return points
})
