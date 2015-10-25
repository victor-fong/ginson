package gods

import (
	"fmt"
	"net/http"
	"time"
	"io/ioutil"
)

func NewDate(month int, day int, year int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func Fetch(symbol string, date time.Time) string{
	var url string = getFetchLink(symbol, date)
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	return string(result)
}

func getFetchLink(symbol string, currentDate time.Time) string {
//	var currentDate time.Time = time.Now()
	var result string = fmt.Sprintf( 
		"http://real-chart.finance.yahoo.com/table.csv?s=%s&a=00&b=1&c=1900&d=%02d&e=%02d&f=%d&g=d&ignore=.csv",
		symbol, (currentDate.Month() - 1), currentDate.Day(), currentDate.Year());
	return result 
}