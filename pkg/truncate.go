package pkg

import (
	"math"
)

// TruncateFloat truncates the floating point number to the desired precision.
func TruncateFloat(f float64, precision int) float64 {
	factor := math.Pow(10, float64(precision))
	return math.Trunc(f*factor) / factor
}
