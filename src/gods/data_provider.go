package gods

import (
	"time"
	"code.google.com/p/rog-go/reverse" 
	"os"
	"io"
	"fmt"
	"strings"
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

func createReader(file_path string) io.ReadSeeker{
	reader, err := os.Open(file_path)
	if err != nil {
		panic(err)
	}
	return reader
}

func (dp CSVDataProvider) dataChannel(startDate time.Time) <- chan *Track {
	result := make(chan *Track, 10)
	
	var reader io.ReadSeeker = createReader(dp.file_path)
	var scanner *reverse.Scanner = reverse.NewScanner(reader)
	
	for scanner.Scan() {
		var text string = scanner.Text()
		fmt.Printf("Text: %s\n", text)
		
		var split_text []string = strings.Split(text, ",")
		for i:=0; i<len(split_text); i++ {
			fmt.Printf("Element[%i] = %s\n", i, split_text[i])
		}
		
//		result <-
	}
	close(result)
	
//	tracks []*Track = 
	
	return result
//	reverse.NewScanner(r io.ReadSeeker) *Scanner
}







