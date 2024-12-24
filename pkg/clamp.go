package pkg

// ClampToUint8 clamps an integer to the range [0, 255].
func ClampToUint8(value int) uint8 {
	return uint8(max(0, min(value, 255))) //nolint
}
