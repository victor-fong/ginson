package gods
 
import (
    "testing"
    "github.com/stretchr/testify/assert"
    "time"
)
 
func TestStub(t *testing.T) {
    assert.True(t, true, "This is good. Canary test passing")
}

func TestActions(t *testing.T) {
	assert.NotEqual(t, BUY, SELL)
}

func TestChoice(t *testing.T) {
	symbol := "abcd"
	action := BUY
	amount := 5000
	
	choice1 := Choice{symbol: symbol, action: action, amount: amount}
	assert.Equal(t, choice1.symbol, symbol)
	assert.Equal(t, choice1.action, action)
	assert.Equal(t, choice1.amount, amount)
}

func TestTrack(t *testing.T){
	symbol := "abcd"
	var open float32 = 15.23
	var low float32 = 14.39
	var high float32 = 15.85
	var end float32 = 15.80
	var volume uint64 = 3939140000
	var date time.Time = time.Now()
	
	track1 := Track{
		date: date,
		symbol: symbol,
		open: open,
		low: low,
		high: high,
		end: end,
		volume: volume,
	}
	
	assert.Equal(t, track1.date, date)
	assert.Equal(t, track1.symbol, symbol)
	assert.Equal(t, track1.open, open)
	assert.Equal(t, track1.low, low)
	assert.Equal(t, track1.high, high)
	assert.Equal(t, track1.end, end)
	assert.Equal(t, track1.volume, volume)
}

func TestFoolName(t *testing.T){
	fool := Fool{}
	assert.Equal(t, fool.name(), "Fool ~_~")
}

func TestFoolDecide(t *testing.T){
	fool := Fool{}
	empty_array := []Choice{}
	result := fool.decide(nil)
	assert.Equal(t, result, empty_array)
	assert.Equal(t, len(result), 0)
	
	track1 := &Track{symbol: "abcd", open: 15.01}
	tracks := []*Track {track1}
	result = fool.decide(tracks)
	assert.Equal(t, result, empty_array)
	assert.Equal(t, len(result), 0)
	
	track2 := &Track{symbol: "edads", open: 15.01}
	tracks = []*Track {track1, track2}
	result = fool.decide(tracks)
	assert.Equal(t, result, empty_array)
	assert.Equal(t, len(result), 0)
}
