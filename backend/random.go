package main

import (
	"crypto/rand"
	"math/big"
)

func secureRandomIntn(n int) int {
	value, err := rand.Int(rand.Reader, big.NewInt(int64(n)))
	if err != nil {
		return 0
	}

	return int(value.Int64())
}
