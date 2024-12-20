package linear

import (
	"github.com/es-debug/backend-academy-2024-go-template/config"
	"github.com/es-debug/backend-academy-2024-go-template/pkg"
)

// Coefficients stores the generated coefficients
// for the linear affine transformation.
type Coefficients struct {
	A, B, C, D, E, F float64
}

// NewCoefficients instantiates a new Coefficients
// entity according to the following constraints:
// a^2 + d^2 < 1
// b^2 + e^2 < 1
// a^2 + b^2 + d^2 + e^2 < 1 + (ae - bd)^2
// a, b, c, d, e, f belong to the interval [-1, 1].
func NewCoefficients() *Coefficients {
	var a, b, d, e float64

	for {
		a = pkg.GetRandomFloat64(config.CoefficientLower, config.CoefficientUpper)
		b = pkg.GetRandomFloat64(config.CoefficientLower, config.CoefficientUpper)
		d = pkg.GetRandomFloat64(config.CoefficientLower, config.CoefficientUpper)
		e = pkg.GetRandomFloat64(config.CoefficientLower, config.CoefficientUpper)

		if a*a+d*d < 1 && b*b+e*e < 1 &&
			a*a+b*b+d*d+e*e < 1+(a*e-b*d)*(a*e-b*d) {
			break
		}
	}

	c := pkg.GetRandomFloat64(config.CoefficientLower, config.CoefficientUpper)
	f := pkg.GetRandomFloat64(config.CoefficientLower, config.CoefficientUpper)

	return &Coefficients{A: a, B: b, C: c,
		D: d, E: e, F: f}
}
