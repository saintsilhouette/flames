package linear

import "github.com/es-debug/backend-academy-2024-go-template/internal/domain/rgb"

// Linear struct stores the cumulative weight of the
// transformation that affects directly the contribution
// of specific nonlinear transformation and initial rgb value.
type Linear struct {
	weight int

	Multipliers *Coefficients
	Color       *rgb.RGB
}

// New instantiates a new Linear entity.
func New(weight int) *Linear {
	return &Linear{
		weight:      weight,
		Multipliers: NewCoefficients(),
		Color:       rgb.New(),
	}
}

// Transform computes the coordinates of the next point
// to be chosen.
func (l *Linear) Transform(x, y float64) (newX, newY float64) {
	newX = l.Multipliers.A*x + l.Multipliers.B*y + l.Multipliers.C
	newY = l.Multipliers.D*x + l.Multipliers.E*y + l.Multipliers.F

	return
}

// Weight returns the assigned weight.
func (l *Linear) Weight() int {
	return l.weight
}
