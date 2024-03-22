package strategies

import (
	"sort"

	"github.com/piquette/finance-go"
	"github.com/wazupwiddat/find-option/model"
)

type AnnualizeOutOfTheMoneyPuts struct {
	quote  *finance.Quote
	values []StrategyValue
}

func (s *AnnualizeOutOfTheMoneyPuts) Run(ideas []model.OptionIdea, q *finance.Quote) {
	s.quote = q
	fIdeas := model.Filter(ideas, model.OutOfTheMoney, model.IsPut)

	s.values = []StrategyValue{}
	for _, i := range fIdeas {
		val := StrategyValue{
			Value: (((i.Strike / (i.Strike - i.Bid)) - 1) * 100) / float64(i.DaysToExpiration) * 365,
			Idea:  i,
		}
		s.values = append(s.values, val)
	}

	sort.Sort(ByStrategyValue(s.values))
}

func (s AnnualizeOutOfTheMoneyPuts) PrintOutput() {
	printHeader("% Annualized OTM Puts")
	printValues(s.values, "VALUE")
}

func init() {
	Strategies = append(Strategies, &AnnualizeOutOfTheMoneyPuts{})
}
