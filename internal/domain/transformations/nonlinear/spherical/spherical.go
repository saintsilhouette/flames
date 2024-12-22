package spherical

// Spherical struct stores the cumulative weight of the
// transformation that affects directly the contribution
// of specific nonlinear transformation.
type Spherical struct {
	weight int
}

// New instantiates a new Spherical entity.
func New(weight int) *Spherical {
	return &Spherical{weight: weight}
}

// Transform computes the coordinates of the next point
// to be chosen.
func (s *Spherical) Transform(x, y float64) (newX, newY float64) {
	rsquare := x*x + y*y

	return x / rsquare, y / rsquare
}

// Weight returns the assigned weight.
func (s *Spherical) Weight() int {
	return s.weight
}
