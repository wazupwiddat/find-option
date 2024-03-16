package main

import (
	"flag"
	"fmt"
	"strings"
	"time"

	"github.com/piquette/finance-go"
	"github.com/piquette/finance-go/datetime"
	"github.com/piquette/finance-go/options"
	"github.com/piquette/finance-go/quote"

	"github.com/wazupwiddat/find-option/model"
	"github.com/wazupwiddat/find-option/strategies"
)

func main() {
	ticker := flag.String("ticker", "AAPL", "This is the ticker to the underlying security you want to find option ideas.")
	weeksOut := flag.Int("weeksout", 5, "How many weeks out do you want to look.  Default: 5")
	flag.Parse()

	// current quote
	q, err := quote.Get(*ticker)
	if err != nil {
		panic("Unable to retrieve quote for " + *ticker)
	}
	// collect Fridays
	fridays := nextSoFridays(*weeksOut)

	// collect Ideas from Straddles
	ideas := []model.OptionIdea{}
	for _, friday := range fridays {
		dt := datetime.New(&friday)
		formattedDate := fmt.Sprintf("%02d-%02d-%d", friday.Month(), friday.Day(), friday.Year())
		fmt.Println("Collecting Straddles for", *ticker, formattedDate)
		iter := options.GetStraddleP(&options.Params{
			UnderlyingSymbol: strings.ToUpper(*ticker),
			Expiration:       dt,
		})
		if iter.Count() == 0 {
			fmt.Println("No options available for", formattedDate)
			thursday := friday.AddDate(0, 0, -1)
			dt = datetime.New(&thursday)
			formattedDate = fmt.Sprintf("%02d-%02d-%d", thursday.Month(), thursday.Day(), thursday.Year())
			fmt.Println("There must be a holiday that day, trying Thursday", formattedDate)
			iter = options.GetStraddleP(&options.Params{
				UnderlyingSymbol: strings.ToUpper(*ticker),
				Expiration:       dt,
			})
			if iter.Count() == 0 {
				fmt.Println("No options available for", thursday)
				fmt.Println("Must be something wrong??? SKIPPING")
				continue
			}
		}
		ideas = append(ideas, collectIdeas(iter)...)
	}

	// Run through all the strategies
	runStrategies(ideas, q)
}

func runStrategies(ideas []model.OptionIdea, q *finance.Quote) {
	for _, strat := range strategies.Strategies {
		strat.Run(ideas, q)
		strat.PrintOutput()
	}
}

func collectIdeas(iter *options.StraddleIter) []model.OptionIdea {
	ideas := []model.OptionIdea{}
	for iter.Next() {
		straddle := *iter.Straddle()
		if straddle.Call == nil || straddle.Put == nil {
			continue
		}
		days := daysToExpiration(straddle)

		callIdea := model.OptionIdea{
			Call:             true,
			Strike:           straddle.Strike,
			Contract:         straddle.Call.Symbol,
			Bid:              straddle.Call.Bid,
			InTheMoney:       straddle.Call.InTheMoney,
			DaysToExpiration: days,
		}
		ideas = append(ideas, callIdea)
		putIdea := model.OptionIdea{
			Call:             false,
			Strike:           straddle.Strike,
			Contract:         straddle.Put.Symbol,
			Bid:              straddle.Put.Bid,
			InTheMoney:       straddle.Put.InTheMoney,
			DaysToExpiration: days,
		}
		ideas = append(ideas, putIdea)
	}
	if err := iter.Err(); err != nil {
		fmt.Println(err)
	}
	return ideas
}

func daysToExpiration(straddle finance.Straddle) int {
	today := zeroOutTime(time.Now().UTC())
	exp := time.Unix(int64(straddle.Call.Expiration), 0).UTC()
	diff := exp.Sub(today)
	return int(diff.Hours() / 24)
}

func nextSoFridays(weeks int) []time.Time {
	var fridays []time.Time

	// Get today's date
	today := zeroOutTime(time.Now().UTC())

	// Find the first Friday
	for i := 0; i < 7; i++ {
		if today.Weekday() == time.Friday {
			break
		}
		today = today.AddDate(0, 0, 1)
	}

	// Collect the next so many Fridays
	thisYear := time.Now().Year()
	for i := 0; i < weeks; i++ {
		if today.Year() == thisYear || today.Year() > thisYear { // Ensure it's within the current year
			fridays = append(fridays, today)
		}
		today = today.AddDate(0, 0, 7) // Move to the next week
	}

	return fridays
}

func zeroOutTime(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}
