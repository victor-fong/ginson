package gods
 
import (
    "testing"
    "github.com/stretchr/testify/assert"
    "fmt"
)
 
func TestStub(t *testing.T) {
    assert.True(t, true, "This is good. Canary test passing")
}

func TestActions(t *testing.T) {
	fmt.Printf("Testing")
	assert.NotEqual(t, BUY, SELL)
	assert.NotEqual(t, SELL, DO_NOTHING)
}

//func TestChoice(){
//	Choice choice = Choice{action: 
//}