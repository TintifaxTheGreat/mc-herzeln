package deal

import "math/bits"

type constraintLead func(hand *bitmap, tricksCount uint) *bitmap

// no constraints
var ConstraintLeadAny = constraintLead(func(hand *bitmap, tricksCount uint) *bitmap {
	return hand
})

// at the first trick, lead of hearts is not allowed
var ConstraintLeadNoHearts = constraintLead(func(hand *bitmap, tricksCount uint) *bitmap {
	if tricksCount != 0 {
		return hand
	}
	result := *hand &^ ALLCOLORS[0]
	if 0 == bits.OnesCount64(uint64(result)) {
		return hand
	}
	return &result
})

type constraintPass func(hand *bitmap, tricksCount uint, leadCard uint) *bitmap

// no constraints
var ConstraintPassAny = constraintPass(func(hand *bitmap, tricksCount uint, leadCard uint) *bitmap {
	return hand
})

// follow suit rule
var ConstraintPassFollowSuit = constraintPass(func(hand *bitmap, tricksCount uint, leadCard uint) *bitmap {
	result := *hand & ALLCOLORS[uint(leadCard/FIGURES)]
	if 0 == bits.OnesCount64(uint64(result)) {
		return hand
	}
	return &result
})
