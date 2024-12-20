package heart

import "math"

// Heart struct stores the cumulative weight of the
// transformation that affects directly the contribution
// of specific nonlinear transformation.
type Heart struct {
	weight int
}

// New instantiates a new Heart entity.
func New(weight int) *Heart {
	return &Heart{weight: weight}
}

// Transform computes the coordinates of the next point
// to be chosen.
func (h *Heart) Transform(x, y float64) (newX, newY float64) {
	newX = math.Sqrt(x*x+y*y) * math.Sin(math.Sqrt(x*x+y*y)*math.Atan2(x, y))
	newY = -math.Sqrt(x*x+y*y) * math.Cos(math.Sqrt(x*x+y*y)*math.Atan2(x, y))

	return newX, newY
}

// Weight returns the assigned weight.
func (h *Heart) Weight() int {
	return h.weight
}
