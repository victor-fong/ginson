package gods

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "math/rand"
    "time"
)

func randomTracks() []*Track{
	tracks := make([]*Track, 100)
	for i := 0; i< 100; i++ {
		tracks[i] = &Track{
			symbol: RandString(4),
			open: rand.Float32(),
			high: rand.Float32(),
			low: rand.Float32(),
			end: rand.Float32(),
		}
	}
	
	return tracks
}

func TestInMemoryDataProvider_dataChannel(t *testing.T) {
	tracks := randomTracks()
	dp := InMemoryDataProvider{tracks: tracks}
	assert.NotEqual(t, dp, nil) 
	
	dataChannel := dp.dataChannel(time.Time{})
	assert.NotEqual(t, dataChannel, nil)
}

func TestInMemoryDataProvider_iteration(t *testing.T) {
	tracks := randomTracks()
	dp := InMemoryDataProvider{tracks: tracks}
	
	i := 0
	for track := range dp.dataChannel(time.Time{}){
		assert.Equal(t, track.symbol, tracks[i].symbol)
		assert.Equal(t, track.open, tracks[i].open)
		assert.Equal(t, track.high, tracks[i].high)
		assert.Equal(t, track.low, tracks[i].low)
		assert.Equal(t, track.end, tracks[i].end)
		i++
	}
}

func TestCreateReader(t *testing.T){
	reader := createReader("../../test_data/test.csv")
	assert.NotEqual(t, reader, nil)	
}

func TestGetSymbolFromPath(t *testing.T){
	var symbol string = getSymbolFromPath("../../test_data/test.csv")
	assert.Equal(t, symbol, "test")
}

func TestParseTime(t *testing.T){
	var result time.Time = parseTime("1995-01-02")
	assert.Equal(t, result.Year(), 1995)
	assert.Equal(t, result.Month(), time.January)
	assert.Equal(t, result.Day(), 2)
	
	result = parseTime("2046-06-30")
	assert.Equal(t, result.Year(), 2046)
	assert.Equal(t, result.Month(), time.June)
	assert.Equal(t, result.Day(), 30)
	
	
}

func TestCSVDataProvider(t *testing.T){
	dp := CSVDataProvider{file_path: "../../test_data/test.csv"}
	
	tracks := make([]*Track, 2)
	tracks[0] = &Track{
		symbol: "test",
		date: parseTime("2015-10-08"),
		open: 1994.01001,
		low: 1987.530029,
		high: 2016.50,
		end: 2013.430054,
		volume: 3939140000,
	}
	
	tracks[1] = &Track{
		symbol: "test",
		date: parseTime("2015-10-09"),
		open: 2013.72998,
		high: 2020.130005,
		low: 2007.609985,
		end: 2014.890015,
		volume: 3706900000,
	}
	
	i := 0
	for track := range dp.dataChannel(time.Time{}){
		assert.Equal(t, track.date, tracks[i].date)
		assert.Equal(t, track.symbol, tracks[i].symbol)
		assert.Equal(t, track.open, tracks[i].open)
		assert.Equal(t, track.high, tracks[i].high)
		assert.Equal(t, track.low, tracks[i].low)
		assert.Equal(t, track.end, tracks[i].end)
		assert.Equal(t, track.volume, tracks[i].volume)
		i++
	}
}
