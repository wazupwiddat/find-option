# Find an option

This applicaion fetches Option Quotes from Yahoo finance and displays a few strategies.

Leverages https://github.com/piquette/finance-go

## Build
    make all

## Run
    % ./bin/foption -help       
      Usage of ./bin/foption:
         -ticker string
    	      This is the ticker to the underlying security you want to find option ideas. (default "AAPL")
         -weeksout int
    	      How many weeks out do you want to look.  Default: 5 (default 5)

## Strategies
    type Strategy interface {
	    Run(ideas []model.OptionIdea, q *finance.Quote)
	    PrintOutput()
    }

  The application collects Option quotes for a particular stock for a period of time.  These are then added to a collection of <b><i>OptionIdea</i></b>.  The strategies iterate over the <b><i>OptionIdea</i></b> collection to produce a <b><i>StrategyValue</i></b> which is then sorted and printed to console.

## List of strategies

  * AnnualizeOutOfTheMoneyCalls
  * AnnualizeOutOfTheMoneyPuts
  * ThreePercentPremiumMaxAssignCalls
  * ThreePercentPremiumMaxAssignPuts

I use the following to decide which option I want to <i>WRITE/SELL</i> when wheeling.
