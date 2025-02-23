package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandomOnwer(t *testing.T) {
	str := RandomOwner()
	require.Len(t, str, 6)
}

func TestRandomMoney(t *testing.T) {
	amount := RandomMoney()
	require.True(t, amount >= 0 && amount <= 1000)

}

func TestRandomCurrency(t *testing.T) {
	curr := RandomCurrency()
	valid_curr := []string{"USD", "EUR", "CAD"}
	require.Contains(t, valid_curr, curr)
}
