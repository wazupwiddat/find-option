package model

type OptionIdea struct {
	Call               bool
	Contract           string
	Strike             float64
	Bid                float64
	InTheMoney         bool
	DaysToExpiration   int
	ReturnOnInvestment float64
	ReturnIfFlat       float64
	ReturnIfAssigned   float64
}
type ByReturnIfFlat []OptionIdea

func (a ByReturnIfFlat) Len() int           { return len(a) }
func (a ByReturnIfFlat) Less(i, j int) bool { return a[i].ReturnIfFlat > a[j].ReturnIfFlat }
func (a ByReturnIfFlat) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type ByReturnIfAssigned []OptionIdea

func (a ByReturnIfAssigned) Len() int           { return len(a) }
func (a ByReturnIfAssigned) Less(i, j int) bool { return a[i].ReturnIfAssigned > a[j].ReturnIfAssigned }
func (a ByReturnIfAssigned) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type OptionIdeaFilterCondition func(idea OptionIdea) bool

func Filter(ideas []OptionIdea, cond ...OptionIdeaFilterCondition) []OptionIdea {
	result := []OptionIdea{}
	for _, idea := range ideas {
		if checkCondition(cond, idea) {
			result = append(result, idea)
		}
	}
	return result
}

func InTheMoney(idea OptionIdea) bool {
	return idea.InTheMoney
}

func OutOfTheMoney(idea OptionIdea) bool {
	return !idea.InTheMoney
}

func IsCall(idea OptionIdea) bool {
	return idea.Call
}

func IsPut(idea OptionIdea) bool {
	return !idea.Call
}

func checkCondition(cond []OptionIdeaFilterCondition, idea OptionIdea) bool {
	for _, c := range cond {
		if !c(idea) {
			return false
		}
	}
	return true
}
