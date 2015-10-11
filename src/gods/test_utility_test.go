package gods

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestRandString(t *testing.T) {
	randomString := RandString(5)
	assert.Equal(t, len(randomString), 5)
	
	randomString2 := RandString(5)
	assert.Equal(t, len(randomString2), 5)
	
	assert.NotEqual(t, randomString, randomString2)
}
