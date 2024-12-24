package config

import "math"

const (
	MinimaX = -1.6 // Lower bound for the x-coordinate of the initial point.
	MaximaX = 1.6  // Upper bound for the x-coordinate of the initial point.
	MinimaY = -1   // Lower bound for the y-coordinate of the initial point.
	MaximaY = 1    // Upper bound for the y-coordinate of the initial point.

	FloatPrecision = 1_000_000 // Controls accuracy of the generated floats.

	CoefficientLower = -1 // Lower bound for the coefficient of the linear affine transformation.
	CoefficientUpper = 1  // Upper bound for the coefficient of the linear affine transformation.

	NumberOfTransformations = 5 // Number of transformations.

	WeightLower = 1  // Lower bound for the weight of the transformation.
	WeightUpper = 15 // Upper bound for the weight of the transformation.

	ColorLower = 0   // Lower bound for the rgb value.
	ColorUpper = 255 // Upper bound for the rgb value.

	Directory = "images" // Directory to store rendered images.

	Delta = 0.00001 // Delta to compare coordinates.

	WidthUpperBound    = 7680 // Upper bound for the image width.
	HeightUpperBound   = 4320 // Upper bound for the image height.
	SymmetryUpperBound = 6    // Upper bound for the symmetry option.

	GammaCoefficient = 2.2 // Gamma correction coefficient.

	NumberOfCores = 16 // Number of cores in my system for the benchmark purpose.

	Circle = 2 * math.Pi // Circle in radians.
)

// Config stores all necessary image properties
// and information that affects on generation process.
type Config struct {
	Width      uint
	Height     uint
	Samples    uint
	Iterations uint
	Goroutines uint
	Symmetry   uint
	Directory  string
}

// New instantiates a new Config entity.
func New(width, height, samples, iterations, goroutines, symmetry uint, directory string) (*Config, error) {
	if width > WidthUpperBound {
		return nil, WidthValueOverflow
	}

	if height > HeightUpperBound {
		return nil, HeightValueOverflow
	}

	if symmetry > SymmetryUpperBound {
		return nil, SymmetryValueOverflow
	}

	return &Config{
		Width:      width,
		Height:     height,
		Samples:    samples,
		Iterations: iterations,
		Goroutines: goroutines,
		Symmetry:   symmetry,
		Directory:  directory,
	}, nil
}
