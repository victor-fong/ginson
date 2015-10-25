package gods

//import "fmt"
import (
	"time"
)

type Track struct {
	date time.Time
	symbol string
	open float32
	high float32
	low float32
	end float32
	volume uint64
}

type God interface {
	decide(openings []*Track) []Choice
	name() string  
}

type Fool struct {
	
}

func (fool Fool) name() string{
	return "Fool ~_~"
}

func (fool Fool) decide(openings []*Track) []Choice{
	return []Choice{}
}

//func (fool Fool) decide(openings []OpenMarket) Choice{
//	action := rand.Int()
//	amount := rand.Int()
//	return Choice{action: action, amount: amount}
//}





