package gods

import (
	"time"
//	"code.google.com/p/rog-go/reverse" 
	"os"
	"io"
)

type DataProvider interface {
	dataChannel(startDate time.Time) <- chan *Track
}

type InMemoryDataProvider struct {
	tracks []*Track
}

func (dp InMemoryDataProvider) dataChannel(startDate time.Time) <- chan *Track {
	result := make(chan *Track, 50)
	
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

func createReader(file_path string) io.Reader{
	reader, err := os.Open(file_path)
	if err != nil {
		panic(err)
	}
	return reader
}

func (dp CSVDataProvider) dataChannel(startDate time.Time) <- chan *Track {
	result := make(chan *Track, 10)
	
	return result
//	reverse.NewScanner(r io.ReadSeeker) *Scanner
}





