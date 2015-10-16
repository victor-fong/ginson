package gods

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "math/rand"
    "time"
    "sort"
    "fmt"
)

func randomTracks(num int) []*Track{
	tracks := make([]*Track, num)
	
	for i := 0; i< num; i++ {
		var year int = rand.Int() % 115 + 1900
		var month int = rand.Int() % 12 + 1
		var day int = rand.Int() % 29 + 1
		
		var str_date = fmt.Sprintf("%d-%02d-%02d", year, month, day)
		
		var date = parseTime(str_date)
		
		tracks[i] = &Track{
			date: date,
			symbol: RandString(4),
			open: rand.Float32(),
			high: rand.Float32(),
			low: rand.Float32(),
			end: rand.Float32(),
		}
	}
	
	sort.Sort(TrackByDate(tracks))
	
	return tracks
}

func TestSortTrack(t *testing.T) {
	var tracks []*Track = randomTracks(50)
	
	for i:=0; i<len(tracks); i++ {
		if i != 0 {
			assert.True(t, tracks[i-1].date.Before(tracks[i].date))
		}
	}
}

func TestInMemoryDataProvider_dataChannel(t *testing.T) {
	tracks := randomTracks(100)
	dp := InMemoryDataProvider{tracks: tracks}
	assert.NotEqual(t, dp, nil) 
	
	dataChannel := dp.dataChannel()
	assert.NotEqual(t, dataChannel, nil)
}

func TestInMemoryDataProvider_iteration(t *testing.T) {
	tracks := randomTracks(100)
	dp := InMemoryDataProvider{tracks: tracks}
	
	i := 0
	for track := range dp.dataChannel(){
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
	for track := range dp.dataChannel(){
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
