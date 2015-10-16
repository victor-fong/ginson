package gods

import (
	"time"
	"code.google.com/p/rog-go/reverse" 
	"os"
	"io"
	"strings"
	"strconv"
)

type DataProvider interface {
	dataChannel(startDate time.Time) <- chan *Track
}

type TrackByDate []*Track

func (a TrackByDate) Len() int           { return len(a) }
func (a TrackByDate) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a TrackByDate) Less(i, j int) bool { return a[i].date.Before(a[j].date) }

type InMemoryDataProvider struct {
	tracks []*Track
}

func (dp InMemoryDataProvider) dataChannel() <- chan *Track {
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

func createReader(file_path string) io.ReadSeeker{
	reader, err := os.Open(file_path)
	if err != nil {
		panic(err)
	}
	return reader
}

func parseFloat32(input string) float32 {
	result, err := strconv.ParseFloat(input, 32)
	if err != nil {
		panic(err)
	}
	return float32(result)
}

func getSymbolFromPath(input string) string {
	var elements []string = strings.Split(input, "/")
	var last_element string = elements[len(elements)-1]
	return strings.Split(last_element, ".")[0]
}

func parseUint64(input string) uint64 {
	result, err := strconv.ParseUint(input, 10, 64)
	if err != nil {
		panic(err)
	}
	return result
}

func parseTime(input string) time.Time {
	result, err := time.Parse("2006-01-02", input)
	if err != nil{
		panic(err)
	}
	return result
}

func (dp CSVDataProvider) dataChannel() <- chan *Track {
	result := make(chan *Track, 10)
	
	var reader io.ReadSeeker = createReader(dp.file_path)
	var scanner *reverse.Scanner = reverse.NewScanner(reader)
	
	var symbol string = getSymbolFromPath(dp.file_path)
	
	go func(){
		for scanner.Scan() {
			var text string = scanner.Text()
			
			var split_text []string = strings.Split(text, ",")
			if(split_text[0] != "Date") {
				var track *Track = &Track{
					symbol: symbol,
					date: parseTime(split_text[0]),
					open: parseFloat32(split_text[1]),
					high: parseFloat32(split_text[2]),
					low: parseFloat32(split_text[3]),
					end: parseFloat32(split_text[4]),
					volume: parseUint64(split_text[5]),
				}
				
				result <- track
			}
		}
		close(result)
	}()
	 
	
	return result
}

func getTrackByDate(channel <-chan *Track, date time.Time) *Track {
	var track *Track = <-channel
	return track
}
