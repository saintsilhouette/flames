package polar

import "math"

// Polar struct stores the cumulative weight of the
// transformation that affects directly the contribution
// of specific nonlinear transformation.
type Polar struct {
	weight int
}

// New instantiates a new Polar entity.
func New(weight int) *Polar {
	return &Polar{weight: weight}
}

// Transform computes the coordinates of the next point
// to be chosen.
func (p *Polar) Transform(x, y float64) (newX, newY float64) {
	newX = math.Atan(y/x) / math.Pi
	newY = math.Sqrt(x*x+y*y) - 1.0

	return
}

// Weight returns the assigned weight.
func (p *Polar) Weight() int {
	return p.weight
}
