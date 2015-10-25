package gods

import (

)

type Long struct {
	symbol string
	entry float64
	amount int64
}

type Short struct {
	symbol string
	entry float64
	amount int64
}

type Portfolio struct {
	Cash float64
	Total float64
	Longs []Long
	Shorts []Short
}

const(
	InitCash float64 = 10000.0
)

func InitPortfolio() Portfolio{
	var result Portfolio = Portfolio{
		Cash: InitCash,
		Total: InitCash,
		Longs: make([]Long, 0),
		Shorts: make([]Short, 0),
	}
	
	return result
}

func battle(gods []God, dataProviders []DataProvider){
	var portfolios map[God]Portfolio = make(map[God]Portfolio)
	for _, god := range gods {
		portfolios[god] = InitPortfolio()
	}
	
	
	
}

