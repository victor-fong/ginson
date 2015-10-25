package gods

import (
	"time"
	"fmt"
)

const (
	BUY_LONG = 1 << iota
    SELL_LONG = 1 << iota
    SELL_SHORT = 1 << iota
    BUY_SHORT = 1 << iota
)

type Choice struct {
	symbol string
	action int
	amount int
	price float64
}

type Long struct {
	Symbol string
	Spend float64
	Amount int
}

type Short struct {
	Symbol string
	Spend float64
	Amount int
}

type Portfolio struct {
	Cash float64
	Total float64
	Longs map[string]Long
	Shorts map[string]Short
	Transactions []History
}

type History struct {
	Date time.Time
	Action int
	Symbol string
	Amount int
}

// TODO: move to config.yml
const InitCash float64 = 10000.0
var InitDate time.Time = time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
var EndDate time.Time = time.Now()

func InitPortfolio() Portfolio{
	var result Portfolio = Portfolio{
		Cash: InitCash,
		Total: InitCash,
		Longs: make(map[string]Long),
		Shorts: make(map[string]Short),
	}
	
	return result
}

func getTrackByDate(dataProviders []DataProviderByDate, date time.Time) []*Track {
	var result []*Track = []*Track {}
	
	for _, dp := range dataProviders {
		var track *Track = dp.getTrackByDate(date)
		if track != nil {
			result = append(result, track)
		}
	}
	
	return result
}

func applyBuyLongChoice(portfolio Portfolio, choice Choice) Portfolio{
	var spend float64 = float64(choice.amount) * choice.price
	portfolio.Cash = portfolio.Cash - spend
	if portfolio.Cash < 0 {
		panic("Running out of cash!")
	}
	
	if longTransaction, element_exist := portfolio.Longs[choice.symbol]; element_exist {
		longTransaction.Amount += choice.amount
		longTransaction.Spend += spend
	} else {

		var longTransaction Long = Long{
			Symbol: choice.symbol,
			Spend: spend,
			Amount: choice.amount,
		}
		
		portfolio.Longs[choice.symbol] = longTransaction		
	}
	
	return portfolio
}

func applySellLongChoice(portfolio Portfolio, choice Choice) Portfolio{
	var gain float64 = float64(choice.amount) * choice.price
	portfolio.Cash = portfolio.Cash + gain
	
	if longTransaction, element_exist := portfolio.Longs[choice.symbol]; !element_exist {
		var message string = fmt.Sprintf("Attempting to sell long on stock %s that is not in portfolio!", choice.symbol)
		panic(message)
	} else {
		
		if longTransaction.Amount < choice.amount {
			var message string = fmt.Sprintf("Attempting to sell %d shares of stock %s but only own %d shares", 
				choice.symbol, choice.amount, longTransaction.Amount)
			panic(message)	
		} else if longTransaction.Amount == choice.amount {
			delete(portfolio.Longs, choice.symbol)
		} else {
			longTransaction.Amount = longTransaction.Amount - choice.amount
			portfolio.Longs[choice.symbol] = longTransaction
		}
	}
	
	return portfolio
}

func addHistory(portfolio Portfolio, choice Choice, date time.Time) Portfolio {
	var history History = History{
		Date: date,
		Action: choice.action,
		Symbol: choice.symbol,
		Amount: choice.amount,
	}
	portfolio.Transactions = append(portfolio.Transactions, history)
	return portfolio
}

func applyChoice(portfolio Portfolio, choices []Choice, date time.Time) Portfolio{
	for _, choice := range choices {
		switch choice.action {
		case BUY_LONG:
			portfolio = applyBuyLongChoice(portfolio, choice)
		case SELL_LONG:
			portfolio = applySellLongChoice(portfolio, choice)
		}
		
		portfolio = addHistory(portfolio, choice, date) 
	}
	
	return portfolio
}

func battle(gods []God, dataProviders []DataProviderByDate) map[God]Portfolio{
	var portfolios map[God]Portfolio = make(map[God]Portfolio)
	for _, god := range gods {
		portfolios[god] = InitPortfolio()
	}
	
	var currentDate time.Time = InitDate
	for currentDate.Before(EndDate) {
		currentDate = currentDate.AddDate(0, 0, 1)
		var tracks []*Track = getTrackByDate(dataProviders, currentDate)
		
		if len(tracks) > 0 {
			for _, god := range gods {
				var choices []Choice = god.decide(tracks)
				portfolios[god] = applyChoice(portfolios[god], choices, currentDate)
				
			}
		}
	} 
	
	return portfolios
}

