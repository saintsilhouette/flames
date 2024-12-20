package sinusoidal

import "math"

// Sinusoidal struct stores the cumulative weight of the
// transformation that affects directly the contribution
// of specific nonlinear transformation.
type Sinusoidal struct {
	weight int
}

// New instantiates a new Sinusoidal entity.
func New(weight int) *Sinusoidal {
	return &Sinusoidal{weight: weight}
}

// Transform computes the coordinates of the next point
// to be chosen.
func (s *Sinusoidal) Transform(x, y float64) (newX, newY float64) {
	newX = math.Sin(x)
	newY = math.Sin(y)

	return
}

// Weight returns the assigned weight.
func (s *Sinusoidal) Weight() int {
	return s.weight
}
