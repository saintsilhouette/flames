package pixel

import (
	"math"

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

// Coloring colors cell is there were no hits by
// taking color from a corresponding transformation.
func (p *Pixel) Coloring(red, green, blue int) {
	p.Color.Red = red
	p.Color.Green = green
	p.Color.Blue = blue
}

// Recoloring recolors cell if there where hits by
// combining cell's and transformation's colors.
func (p *Pixel) Recoloring(red, green, blue int) {
	p.Color.Red = (p.Color.Red + red) / 2
	p.Color.Green = (p.Color.Green + green) / 2
	p.Color.Blue = (p.Color.Blue + blue) / 2
}

// LogCorrection performs log correction over the pixel
// and returns the resulting value.
func (p *Pixel) LogCorrection() float64 {
	p.Normal = math.Log10(float64(p.Hits))
	return p.Normal
}

// GammaCorrection performs gamma correction over the pixel.
func (p *Pixel) GammaCorrection(gamma, maxima float64) {
	p.Normal /= maxima
	p.Color.Red = int(float64(p.Color.Red) * math.Pow(p.Normal, 1.0/gamma))
	p.Color.Green = int(float64(p.Color.Green) * math.Pow(p.Normal, 1.0/gamma))
	p.Color.Blue = int(float64(p.Color.Blue) * math.Pow(p.Normal, 1.0/gamma))
}
