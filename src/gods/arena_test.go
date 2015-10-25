package gods

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestInitPortfolio(t *testing.T) {
	var portfolio Portfolio = InitPortfolio()
	
	assert.Equal(t, InitCash, portfolio.Cash)
	assert.Equal(t, InitCash, portfolio.Total)
	assert.Equal(t, 0, len(portfolio.Longs))
	assert.Equal(t, 0, len(portfolio.Shorts))
}
