package config

import "strconv"

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
)

// Config stores all necessary image properties
// and information that affects on generation process.
type Config struct {
	Width      int
	Height     int
	Samples    int
	Iterations int
	Goroutines int
	Directory  string
}

// New instantiates a new Config entity.
func New(width, height, samples, iterations, goroutines, directory string) (*Config, error) {
	numericWidth, err := strconv.Atoi(width)
	if err != nil {
		return nil, InvalidWidthError
	}

	if numericWidth < 0 {
		return nil, NegativeWidthError
	}

	numericHeight, err := strconv.Atoi(height)
	if err != nil {
		return nil, InvalidHeightError
	}

	if numericHeight < 0 {
		return nil, NegativeHeightError
	}

	numericSamples, err := strconv.Atoi(samples)
	if err != nil {
		return nil, InvalidSamplesError
	}

	if numericSamples < 0 {
		return nil, NegativeSamplesError
	}

	numericIterations, err := strconv.Atoi(iterations)
	if err != nil {
		return nil, InvalidIterationsError
	}

	if numericIterations < 0 {
		return nil, NegativeIterationsError
	}

	numericGoroutines, err := strconv.Atoi(goroutines)
	if err != nil {
		return nil, InvalidGoroutinesError
	}

	if numericGoroutines < 0 {
		return nil, NegativeGoroutinesError
	}

	return &Config{
		Width:      numericWidth,
		Height:     numericHeight,
		Samples:    numericSamples,
		Iterations: numericIterations,
		Goroutines: numericGoroutines,
		Directory:  directory,
	}, nil
}
