package pixel

import (
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/rgb"
)

// Pixel is a main bulding unit.
type Pixel struct {
	Color  *rgb.RGB
	Hits   int
	Normal float64
}

// New instantiates a new Pixel entity.
func New() *Pixel {
	return &Pixel{
		Color: rgb.New(),
		Hits:  0,
	}
}
