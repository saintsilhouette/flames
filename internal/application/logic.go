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

// FillTransformations creates linear and nonlinear
// transformations as well as specifies the weight
// of each transformation.
func (a *Application) FillTransformations() {
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

// SelectLinearTransformation randomly selects the linear
// transformation from the generated slice.
func (a *Application) SelectLinearTransformation() *linear.Linear {
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

// ConstructImage creates image based on 2D slice of pixels.
func (a *Application) ConstructImage() *image.RGBA {
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

// ScaleCoordinates scales transformed newX and newY
// with respect to the image size.
func (a *Application) ScaleCoordinates(newX, newY float64) (x, y int) {
	x = int(a.Config.Width) - int(math.Trunc(((config.MaximaX-newX)/ //nolint
		(config.MaximaX-config.MinimaX))*float64(a.Config.Width)))
	y = int(a.Config.Height) - int(math.Trunc(((config.MaximaY-newY)/ //nolint
		(config.MaximaY-config.MinimaY))*float64(a.Config.Height)))

	return
}

// InsideBounds checks whether transformed newX and newY
// satisfies the necessary bounds.
func (a *Application) InsideBounds(newX, newY float64) bool {
	return config.MinimaX <= newX && newX <= config.MaximaX &&
		config.MinimaY <= newY && newY <= config.MaximaY
}

// InsideImage checks whether the obtained point inside
// the image or not.
func (a *Application) InsideImage(x, y int) bool {
	return 0 <= x && x < int(a.Config.Width) && 0 <= y && y < int(a.Config.Height) //nolint
}

// Coloring colors cell is there were no hits by
// taking color from a corresponding transformation.
func (a *Application) Coloring(x, y int, linearTransformation *linear.Linear) {
	a.Pixels[x][y].Color.Red = linearTransformation.Color.Red
	a.Pixels[x][y].Color.Green = linearTransformation.Color.Green
	a.Pixels[x][y].Color.Blue = linearTransformation.Color.Blue
}

// Correction performs brightness correction on
// the number of hits per pixel.
func (a *Application) Correction() {
	maxima, gamma := 0.0, 2.2

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

// ProcessCell implements logic of coloring the
// provided point with checking all necessary condition.
func (a *Application) ProcessCell(newX, newY float64) {
	mu := &sync.Mutex{}

	for i := -20; i < int(a.Config.Iterations); i++ { //nolint
		linearTransformation := a.SelectLinearTransformation()
		newX, newY = linearTransformation.Transform(newX, newY)

		nonlinearTransformation := a.SelectNonlinearTransformation()
		newX, newY = nonlinearTransformation.Transform(newX, newY)

		if i >= 0 && a.InsideBounds(newX, newY) {
			x, y := a.ScaleCoordinates(newX, newY)
			if a.InsideImage(x, y) {
				mu.Lock()

				red := linearTransformation.Color.Red
				green := linearTransformation.Color.Green
				blue := linearTransformation.Color.Blue

				if a.Pixels[x][y].Hits == 0 {
					a.Pixels[x][y].Coloring(red, green, blue)
				} else {
					a.Pixels[x][y].Recoloring(red, green, blue)
				}

				a.Pixels[x][y].Hits++

				mu.Unlock()
			}
		}
	}
}
