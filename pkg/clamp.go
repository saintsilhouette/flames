package pkg

import "math"

// ClampToUint8 clamps an integer to the range [0, 255].
func ClampToUint8(value int) uint8 {
	return uint8(math.Max(0, math.Min(float64(value), 255)))
}
