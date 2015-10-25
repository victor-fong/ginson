package gods
 
import (
    "testing"
    "github.com/stretchr/testify/assert"
    "time"
)
 
//func TestFetch(t *testing.T) {
//    Fetch(
//}
 
 func TestGetFetchLink(t *testing.T) {
 	var date time.Time = NewDate(12, 31, 2015)
 	var result string = getFetchLink("AAPL", date)
 	
 	assert.Equal(t, 
 		"http://real-chart.finance.yahoo.com/table.csv?s=AAPL&a=00&b=1&c=1900&d=11&e=31&f=2015&g=d&ignore=.csv",
 		result)
 	
 	date = NewDate(1, 7, 2015)
 	result = getFetchLink("ABC", date)
 	assert.Equal(t, 
 		"http://real-chart.finance.yahoo.com/table.csv?s=ABC&a=00&b=1&c=1900&d=00&e=07&f=2015&g=d&ignore=.csv",
 		result)
 }
 
 func TestFetch(t *testing.T) {
 	var date time.Time = NewDate(1, 1, 2015)
	Fetch("AAPL", date)
 }