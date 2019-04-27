package deal

import "math/bits"

type goal func(players AllPlayers) [PLAYERS] int

// avoid hearts in the tricks
var GoalNoHearts = goal(func(players AllPlayers) [PLAYERS] int {
	var points [PLAYERS] int
	for player := uint(0); player < PLAYERS; player++ {
		// count all hearts in the player's tricks
		points[player] -= bits.OnesCount64(uint64(ALLCOLORS[0] & *players[player].Card().tricks))
	}
	return points
})

// avoid making tricks
var GoalNoTricks = goal(func(players AllPlayers) [PLAYERS] int {
	var points [PLAYERS] int
	for player := uint(0); player < PLAYERS; player++ {
		// count all the players tricks
		points[player] -= bits.OnesCount64(uint64(*players[player].Card().tricks)) / 4
	}
	return points
})

// avoid queens in the tricks
var GoalNoQueens = goal(func(players AllPlayers) [PLAYERS] int {
	var points [PLAYERS] int
	for player := uint(0); player < PLAYERS; player++ {
		// count all queens in the player's tricks
		points[player] -= 2 * bits.OnesCount64(uint64(ALLFIGURES[2] & *players[player].Card().tricks))
	}
	return points
})

/*
// avoid the king of hearts, and making the last trick
var GoalNoGepetto = goal(func(players AllPlayers) [PLAYERS] int {
	var points [PLAYERS] int
	for player := uint(0); player < PLAYERS; player++ {
		// count all hearts in the player's tricks
		points[player] -= bits.OnesCount64(uint64(ALLCOLORS[0] & *players[player].Card().tricks))
	}
	return points
})
*/