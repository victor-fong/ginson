package gods

import (
	"time"
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

// TODO: move to config.yml
const InitCash float64 = 10000.0
var InitDate time.Time = time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
var EndDate time.Time = time.Now()

func InitPortfolio() Portfolio{
	var result Portfolio = Portfolio{
		Cash: InitCash,
		Total: InitCash,
		Longs: make([]Long, 0),
		Shorts: make([]Short, 0),
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

func applyChoice(portfolio Portfolio, choices []Choice) {
	
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
		
		for _, god := range gods {
			var choices []Choice = god.decide(tracks)
			applyChoice(portfolios[god], choices)
		}
	} 
	
	return portfolios
}

