package strategies

import (
	"sort"

	"github.com/piquette/finance-go"
	"github.com/wazupwiddat/find-option/model"
)

type ThreePercentPremiumMaxAssignPuts struct {
	quote  *finance.Quote
	values []StrategyValue
}

func (s *ThreePercentPremiumMaxAssignPuts) Run(ideas []model.OptionIdea, q *finance.Quote) {
	s.quote = q
	fIdeas := model.Filter(ideas, model.OutOfTheMoney, model.IsPut, func(idea model.OptionIdea) bool {
		premPercent := (idea.Bid / idea.Strike) * 100
		return premPercent > 3
	})
	for _, i := range fIdeas {
		val := StrategyValue{
			Value: ((q.Ask / (i.Strike - i.Bid)) - 1) * 100,
			Idea:  i,
		}
		s.values = append(s.values, val)
	}
	sort.Sort(ByStrategyValue(s.values))
}

func (s ThreePercentPremiumMaxAssignPuts) PrintOutput() {
	printHeader("3% Premiums Max % Assigned Puts")
	printValues(s.values, "VALUE")
}

func init() {
	Strategies = append(Strategies, &ThreePercentPremiumMaxAssignPuts{})
}
