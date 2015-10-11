package gods

//import "fmt"
import (
	"time"
)

const (
	BUY = 1 << iota  // a == 1 (iota has been reset)
    SELL = 1 << iota  // b == 2
)

type Choice struct {
	symbol string
	action int
	amount int
}

type Track struct {
	date time.Time
	symbol string
	open float32
	high float32
	low float32
	end float32
}

type God interface {
	decide(openings []Track) []Choice
	name() string  
}

type Fool struct {
	
}

func (fool Fool) name() string{
	return "Fool ~_~"
}

func (fool Fool) decide(openings []Track) []Choice{
	return []Choice{}
}

//func (fool Fool) decide(openings []OpenMarket) Choice{
//	action := rand.Int()
//	amount := rand.Int()
//	return Choice{action: action, amount: amount}
//}





