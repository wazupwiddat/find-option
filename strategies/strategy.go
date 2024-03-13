package strategies

import (
	"github.com/fatih/color"
	"github.com/piquette/finance-go"
	"github.com/wazupwiddat/find-option/model"
)

type Strategy interface {
	Run(ideas []model.OptionIdea, q *finance.Quote)
	PrintOutput()
}

var Strategies []Strategy

type StrategyValue struct {
	Value float64
	Idea  model.OptionIdea
}

type ByStrategyValue []StrategyValue

func (a ByStrategyValue) Len() int { return len(a) }
func (a ByStrategyValue) Less(i, j int) bool {
	return a[i].Value > a[j].Value
}
func (a ByStrategyValue) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func printHeader(header string) {
	c := color.New(color.FgGreen).Add(color.Bold)
	c.Println(" ")
	c.Println(" ")
	c.Println(header)
}

func printValues(values []StrategyValue, valueHeader string) {
	cnt := 0
	h := color.New(color.FgGreen).Add(color.Bold).Add(color.Underline)
	h.Printf("%30s %7s %4s %9s %9s\n", "contract", "bid", "days", valueHeader, "strike")

	for _, value := range values {
		if cnt > 10 {
			break
		}

		cnt++
		printValue(value)

	}
}

func printValue(value StrategyValue) {
	c := color.New(color.FgGreen)

	c.Printf("%30s|%7.2f|%4d|%9.2f|%9.2f\n\r",
		value.Idea.Contract, value.Idea.Bid, value.Idea.DaysToExpiration, value.Value,
		value.Idea.Strike)
}
