package deal

type constraintFirstLead func(hand *bitmap) *bitmap

// no constraints
var ConstraintLeadAny = constraintFirstLead(func (hand *bitmap) *bitmap {
return hand
})

// at the first trick, lead of hearts is not allowed
var ConstraintLeadNoHearts = constraintFirstLead(func (hand *bitmap) *bitmap {
result := *hand & ALLCOLORS[0]
return &result
})

type constraintPassAll func (hand *bitmap, leadCard uint) *bitmap

// no constraints
var ConstraintPassAny = constraintPassAll(func (hand *bitmap, leadCard uint) *bitmap {
return hand
})

// follow suit rule
var ConstraintPassFollowSuit = constraintPassAll(func (hand *bitmap, leadCard uint) *bitmap {
result := *hand & ALLCOLORS[0]
return &result
})
