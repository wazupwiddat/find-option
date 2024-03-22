package strategies

import (
	"sort"

	"github.com/piquette/finance-go"
	"github.com/wazupwiddat/find-option/model"
)

type AnnualizeOutOfTheMoneyCalls struct {
	quote  *finance.Quote
	values []StrategyValue
}

func (s *AnnualizeOutOfTheMoneyCalls) Run(ideas []model.OptionIdea, q *finance.Quote) {
	s.quote = q
	fIdeas := model.Filter(ideas, model.OutOfTheMoney, model.IsCall)

	s.values = []StrategyValue{}
	for _, i := range fIdeas {
		val := StrategyValue{
			Value: ((((i.Bid + i.Strike) / (i.Strike)) - 1) * 100) / float64(i.DaysToExpiration) * 365,
			Idea:  i,
		}
		s.values = append(s.values, val)
	}

	sort.Sort(ByStrategyValue(s.values))
}

func (s AnnualizeOutOfTheMoneyCalls) PrintOutput() {
	printHeader("% Annualized OTM Call")
	printValues(s.values, "VALUE")
}

func init() {
	Strategies = append(Strategies, &AnnualizeOutOfTheMoneyCalls{})
}
