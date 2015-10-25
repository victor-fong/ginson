package gods

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"time"
)

func TestInitPortfolio(t *testing.T) {
	var portfolio Portfolio = InitPortfolio()
	
	assert.Equal(t, InitCash, portfolio.Cash)
	assert.Equal(t, InitCash, portfolio.Total)
	assert.Equal(t, 0, len(portfolio.Longs))
	assert.Equal(t, 0, len(portfolio.Shorts))
}

func createDataProviders(tracks []*Track) []DataProviderByDate {
	dp := InMemoryDataProvider{tracks: tracks}
	dataChannel := dp.dataChannel()
	
	var dataProvider DataProviderByDate = DataProviderByDate{
		data_channel: dataChannel,
	}
	
	var dataProviders []DataProviderByDate = []DataProviderByDate { dataProvider }
	return dataProviders
}

func testTrackByDate(t *testing.T) {
	tracks := randomTracks(100)
	var dataProviders []DataProviderByDate = createDataProviders(tracks)
	
	var results []*Track = getTrackByDate(dataProviders, tracks[79].date)
	assert.Equal(t, 1, len(results))
	assert.Equal(t, tracks[79].date, results[0].date)
	assert.Equal(t, tracks[79].open, results[0].open)
	
}

func TestBattle(t *testing.T) {
	tracks := randomTracks(100)
	var dataProviders []DataProviderByDate = createDataProviders(tracks)
	var fool Fool = Fool{}
	
	var gods []God = []God { fool }
	
	var portfolios map[God]Portfolio = battle(gods, dataProviders)
	
	assert.Equal(t, 1, len(portfolios))
	assert.Equal(t, InitCash, portfolios[fool].Total)
}

func TestApplyChoice_Buy(t *testing.T) {
	var portfolio Portfolio = InitPortfolio() 
	
	var choice Choice = Choice {
		action: BUY_LONG,
		symbol: "ABCD",
		price: 15.50,
		amount: 200,
	}
	
	portfolio = applyChoice(portfolio, []Choice{ choice }, time.Now())
	
	assert.Equal(t, (InitCash - 15.50 * 200), portfolio.Cash)
	assert.Equal(t, 1, len(portfolio.Longs))
	assert.Equal(t, 0, len(portfolio.Shorts))
	assert.Equal(t, 1, len(portfolio.Transactions))
}

func TestApplyChoice_SellLong_NotOwn(t *testing.T) {
	var portfolio Portfolio = InitPortfolio() 
	
	var choice Choice = Choice {
		action: SELL_LONG,
		symbol: "ABCD",
		price: 15.50,
		amount: 200,
	}
	
	assert.Panics(t, func(){
		applyChoice(portfolio, []Choice{ choice }, time.Now())
	}, "")
}

func TestApplyChoice_SellLong_NotOwnEnough(t *testing.T) {
	var portfolio Portfolio = InitPortfolio() 
	
	var buyChoice Choice = Choice {
		action: BUY_LONG,
		symbol: "ABCD",
		price: 15.50,
		amount: 100,
	}
	portfolio = applyChoice(portfolio, []Choice{ buyChoice }, time.Now())
	
	var sellChoice Choice = Choice {
		action: SELL_LONG,
		symbol: "ABCD",
		price: 15.50,
		amount: 200,
	}
	
	assert.Panics(t, func(){
		applyChoice(portfolio, []Choice{ sellChoice }, time.Now())
	}, "")
}

func TestApplyChoice_SellLong_All(t *testing.T) {
	var portfolio Portfolio = InitPortfolio() 
	
	var buyChoice Choice = Choice {
		action: BUY_LONG,
		symbol: "ABCD",
		price: 15.50,
		amount: 200,
	}
	portfolio = applyChoice(portfolio, []Choice{ buyChoice }, time.Now())
	
	assert.Equal(t, (InitCash - 15.50 * 200), portfolio.Cash)
	assert.Equal(t, 1, len(portfolio.Longs))
	assert.Equal(t, 0, len(portfolio.Shorts))
	assert.Equal(t, 1, len(portfolio.Transactions))
	
	var sellChoice Choice = Choice {
		action: SELL_LONG,
		symbol: "ABCD",
		price: 16.50,
		amount: 200,
	}
	
	portfolio = applyChoice(portfolio, []Choice{ sellChoice }, time.Now())
	
	assert.Equal(t, (InitCash + 200), portfolio.Cash)
	assert.Equal(t, 0, len(portfolio.Longs))
	assert.Equal(t, 0, len(portfolio.Shorts))
	assert.Equal(t, 2, len(portfolio.Transactions))
}

func TestApplyChoice_SellLong_Part(t *testing.T) {
	var portfolio Portfolio = InitPortfolio() 
	
	var buyChoice Choice = Choice {
		action: BUY_LONG,
		symbol: "ABCD",
		price: 15.50,
		amount: 200,
	}
	portfolio = applyChoice(portfolio, []Choice{ buyChoice }, time.Now())
	
	assert.Equal(t, (InitCash - 15.50 * 200), portfolio.Cash)
	assert.Equal(t, 1, len(portfolio.Longs))
	assert.Equal(t, 0, len(portfolio.Shorts))
	assert.Equal(t, 1, len(portfolio.Transactions))
	
	var sellChoice Choice = Choice {
		action: SELL_LONG,
		symbol: "ABCD",
		price: 16.50,
		amount: 100,
	}
	
	portfolio = applyChoice(portfolio, []Choice{ sellChoice }, time.Now())
	
	assert.Equal(t, (InitCash - 1450), portfolio.Cash)
	assert.Equal(t, 1, len(portfolio.Longs))
	assert.Equal(t, 100, portfolio.Longs["ABCD"].Amount)
	assert.Equal(t, 0, len(portfolio.Shorts))
	assert.Equal(t, 2, len(portfolio.Transactions))
	
	portfolio = applyChoice(portfolio, []Choice{ sellChoice }, time.Now())
	
	assert.Equal(t, (InitCash + 200), portfolio.Cash)
	assert.Equal(t, 0, len(portfolio.Longs))
	assert.Equal(t, 0, len(portfolio.Shorts))
	assert.Equal(t, 3, len(portfolio.Transactions))
}
