package gods

//import "fmt"

const (
	BUY = 1 << iota  // a == 1 (iota has been reset)
    SELL = 1 << iota  // b == 2
	DO_NOTHING = 1 << iota  // c == 4
)
//
//type Choice struct {
//	action uint8
//	amount uint16
//}

//type god interface {
//	decide(openings []OpenMarket) Choice
//	name() string  
//}
//
//type Fool interface {
//	
//}

//func (fool Fool) decide(openings []OpenMarket) Choice{
//	action := rand.Int()
//	amount := rand.Int()
//	return Choice{action: action, amount: amount}
//}





