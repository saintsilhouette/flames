package application

import (
	"image"
	"image/color"
	"math"
	"sort"
	"sync"

	"github.com/es-debug/backend-academy-2024-go-template/config"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/transformations/linear"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/transformations/nonlinear/disk"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/transformations/nonlinear/heart"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/transformations/nonlinear/polar"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/transformations/nonlinear/sinusoidal"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/transformations/nonlinear/spherical"
	"github.com/es-debug/backend-academy-2024-go-template/pkg"
)

// fillTransformations creates linear and nonlinear
// transformations as well as specifies the weight
// of each transformation.
func (a *Application) fillTransformations() {
	linearTransformations := make([]*linear.Linear, 0, config.NumberOfTransformations)

	for range config.NumberOfTransformations {
		weight := pkg.GetRandomInt(config.WeightLower, config.WeightUpper)
		linearTransformations = append(linearTransformations, linear.New(weight))
	}

	sort.Slice(linearTransformations, func(i, j int) bool {
		return linearTransformations[i].Weight() < linearTransformations[j].Weight()
	})

	nonlinearTransformations := make([]Transformation, 0, config.NumberOfTransformations)

	nonlinearTransformations = append(nonlinearTransformations,
		sinusoidal.New(pkg.GetRandomInt(config.WeightLower, config.WeightUpper)),
		spherical.New(pkg.GetRandomInt(config.WeightLower, config.WeightUpper)),
		polar.New(pkg.GetRandomInt(config.WeightLower, config.WeightUpper)),
		heart.New(pkg.GetRandomInt(config.WeightLower, config.WeightUpper)),
		disk.New(pkg.GetRandomInt(config.WeightLower, config.WeightUpper)))

	sort.Slice(nonlinearTransformations, func(i, j int) bool {
		return nonlinearTransformations[i].Weight() < nonlinearTransformations[j].Weight()
	})

	a.Linear = linearTransformations
	a.Nonlinear = nonlinearTransformations
}

// selectLinearTransformation randomly selects the linear
// transformation from the generated slice.
func (a *Application) selectLinearTransformation() *linear.Linear {
	totalWeight := 0
	cumulativeWeights := make([]int, 0, config.NumberOfTransformations)

	for _, linear := range a.Linear {
		totalWeight += linear.Weight()
		cumulativeWeights = append(cumulativeWeights, linear.Weight())
	}

	r := pkg.GetRandomInt(0, totalWeight)

	index := sort.SearchInts(cumulativeWeights, r)

	return a.Linear[min(index, len(a.Linear)-1)]
}

// SelectNonlinearTransformation randomly selects the nonlinear
// transformation from the generated slice.
func (a *Application) SelectNonlinearTransformation() Transformation {
	totalWeight := 0
	cumulativeWeights := make([]int, 0, config.NumberOfTransformations)

	for _, nonlinear := range a.Nonlinear {
		totalWeight += nonlinear.Weight()
		cumulativeWeights = append(cumulativeWeights, nonlinear.Weight())
	}

	r := pkg.GetRandomInt(0, totalWeight)

	index := sort.SearchInts(cumulativeWeights, r)

	return a.Nonlinear[min(index, len(a.Nonlinear)-1)]
}

// constructImage creates image based on 2D slice of pixels.
func (a *Application) constructImage() *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, int(a.Config.Width), int(a.Config.Height))) //nolint

	for i := range a.Config.Width {
		for j := range a.Config.Height {
			if a.Pixels[i][j].Hits > 0 {
				col := color.RGBA{
					R: pkg.ClampToUint8(a.Pixels[i][j].Color.Red),
					G: pkg.ClampToUint8(a.Pixels[i][j].Color.Green),
					B: pkg.ClampToUint8(a.Pixels[i][j].Color.Blue),
					A: 255,
				}
				img.Set(int(i), int(j), col) //nolint
			}
		}
	}

	return img
}

// scaleCoordinates scales transformed newX and newY
// with respect to the image size.
func (a *Application) scaleCoordinates(newX, newY float64) (x, y float64) {
	x = float64(a.Config.Width) - math.Trunc(((config.MaximaX-newX)/ //nolint
		(config.MaximaX-config.MinimaX))*float64(a.Config.Width))
	y = float64(a.Config.Height) - math.Trunc(((config.MaximaY-newY)/ //nolint
		(config.MaximaY-config.MinimaY))*float64(a.Config.Height))

	return x, y
}

// rotate performs rotatation of the selected pixel
// using rotation matrix.
func (a *Application) rotate(x, y, theta float64) (rotX, rotY int) {
	rotX = int(x*math.Cos(theta) - y*math.Sin(theta))
	rotY = int(x*math.Sin(theta) + y*math.Cos(theta))

	return rotX, rotY
}

// insideBounds checks whether transformed newX and newY
// satisfies the necessary bounds.
func (a *Application) insideBounds(newX, newY float64) bool {
	return config.MinimaX <= newX && newX <= config.MaximaX &&
		config.MinimaY <= newY && newY <= config.MaximaY
}

// insideImage checks whether the obtained point inside
// the image or not.
func (a *Application) insideImage(x, y int) bool {
	return 0 <= x && x < int(a.Config.Width) && 0 <= y && y < int(a.Config.Height) //nolint
}

// correction performs brightness correction on
// the number of hits per pixel.
func (a *Application) correction() {
	maxima, gamma := 0.0, config.GammaCoefficient

	for i := range a.Config.Width {
		for j := range a.Config.Height {
			if a.Pixels[i][j].Hits != 0 {
				maxima = max(maxima, a.Pixels[i][j].LogCorrection())
			}
		}
	}

	for i := range a.Config.Width {
		for j := range a.Config.Height {
			a.Pixels[i][j].GammaCorrection(gamma, maxima)
		}
	}
}

// processCell implements logic of coloring the
// provided point with checking all necessary condition.
func (a *Application) processCell(newX, newY float64) {
	mu := &sync.Mutex{}

	for i := -20; i < int(a.Config.Iterations); i++ { //nolint
		linearTransformation := a.selectLinearTransformation()
		newX, newY = linearTransformation.Transform(newX, newY)

		nonlinearTransformation := a.SelectNonlinearTransformation()
		newX, newY = nonlinearTransformation.Transform(newX, newY)

		var theta float64

		for range a.Config.Symmetry {
			if i >= 0 && a.insideBounds(newX, newY) {
				x, y := a.scaleCoordinates(newX, newY)

				rotX, rotY := a.rotate(x, y, theta)
				theta += config.Circle / float64(a.Config.Symmetry)

				if a.insideImage(rotX, rotY) {
					red := linearTransformation.Color.Red
					green := linearTransformation.Color.Green
					blue := linearTransformation.Color.Blue

					mu.Lock()

					if a.Pixels[rotX][rotY].Hits == 0 {
						a.Pixels[rotX][rotY].Coloring(red, green, blue)
					} else {
						a.Pixels[rotX][rotY].Recoloring(red, green, blue)
					}

					a.Pixels[rotX][rotY].Hits++

					mu.Unlock()
				}
			}
		}
	}
}
