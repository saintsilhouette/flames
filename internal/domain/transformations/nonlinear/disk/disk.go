package disk

import "math"

// Disk struct stores the cumulative weight of the
// transformation that affects directly the contribution
// of specific nonlinear transformation.
type Disk struct {
	weight int
}

// New instantiates a new Disk entity.
func New(weight int) *Disk {
	return &Disk{weight: weight}
}

// Transform computes the coordinates of the next point
// to be chosen.
func (d *Disk) Transform(x, y float64) (newX, newY float64) {
	r := math.Sqrt(x*x + y*y)
	theta := math.Atan2(x, y)

	return theta / math.Pi * math.Sin(math.Pi*r), theta / math.Pi * math.Cos(math.Pi*r)
}

// Weight returns the assigned weight.
func (d *Disk) Weight() int {
	return d.weight
}
