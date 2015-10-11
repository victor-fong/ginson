package gods

import (
	"time"
)

type DataProvider interface {
	dataChannel(startDate time.Time) <- chan *Track
}

type InMemoryDataProvider struct {
	tracks []*Track
}

func (dp InMemoryDataProvider) dataChannel(startDate time.Time) <- chan *Track {
	result := make(chan *Track, 10)
	
	go func(){
		for _, track := range dp.tracks{
			result <- track
		}
		close(result)
	}()
	
	return result
}

type CSVDataProvider struct {
	file_path string
}





