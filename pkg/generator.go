package pkg

import (
	"crypto/rand"
	"math/big"

	"github.com/es-debug/backend-academy-2024-go-template/config"
)

// GetRandomInt generates pseudorandom integer in the given range.
func GetRandomInt(left, right int) int {
	nBig, _ := rand.Int(rand.Reader, big.NewInt(int64(right-left+1)))
	n := nBig.Int64()

	return int(n) + left
}

// GetRandomFloat64 generates pseudorandom float in the given range.
func GetRandomFloat64(left, right float64) float64 {
	leftInt := int(left * config.FloatPrecision)
	rightInt := int(right * config.FloatPrecision)

	return float64(GetRandomInt(leftInt, rightInt)) / config.FloatPrecision
}
