package gods

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "math/rand"
    "time"
//    "fmt"
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