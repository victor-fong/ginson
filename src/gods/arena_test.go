package gods

import (
	"testing"
	"github.com/stretchr/testify/assert"
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
