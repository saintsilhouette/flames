package application

import (
	"image"
	"image/color"
	"math"
	"sort"

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
	linearTransformations := make([]*linear.Linear, 0, 5)

	for range config.NumberOfTransformations {
		weight := pkg.GetRandomInt(config.WeightLower, config.WeightUpper)
		linearTransformations = append(linearTransformations, linear.New(weight))
	}

	sort.Slice(linearTransformations, func(i, j int) bool {
		return linearTransformations[i].Weight() < linearTransformations[j].Weight()
	})

	nonlinearTransformations := make([]Transformation, 0, 5)

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
	cumulativeWeights := make([]int, config.NumberOfTransformations)

	for i, linear := range a.Linear {
		totalWeight += linear.Weight()
		cumulativeWeights[i] = linear.Weight()
	}

	r := pkg.GetRandomInt(0, totalWeight)

	index := sort.SearchInts(cumulativeWeights, r)

	if index == config.NumberOfTransformations {
		return a.Linear[index-1]
	}

	return a.Linear[index]
}

// SelectNonlinearTransformation randomly selects the nonlinear
// transformation from the generated slice.
func (a *Application) SelectNonlinearTransformation() Transformation {
	totalWeight := 0
	cumulativeWeights := make([]int, config.NumberOfTransformations)

	for i, nonlinear := range a.Nonlinear {
		totalWeight += nonlinear.Weight()
		cumulativeWeights[i] = nonlinear.Weight()
	}

	r := pkg.GetRandomInt(0, totalWeight)

	index := sort.SearchInts(cumulativeWeights, r)

	if index == config.NumberOfTransformations {
		return a.Nonlinear[index-1]
	}

	return a.Nonlinear[index]
}

func (a *Application) ConstructImage() *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, a.Config.Width, a.Config.Height))

	for i := range a.Config.Width {
		for j := range a.Config.Height {
			if a.Pixels[i][j].Hits > 0 {
				col := color.RGBA{
					R: pkg.ClampToUint8(a.Pixels[i][j].Color.Red),
					G: pkg.ClampToUint8(a.Pixels[i][j].Color.Green),
					B: pkg.ClampToUint8(a.Pixels[i][j].Color.Blue),
					A: 255,
				}
				img.Set(i, j, col)
			}
		}
	}

	return img
}

// ScaleCoordinates scales transformed newX and newY
// with respect to the image size.
func (a *Application) ScaleCoordinates(newX, newY float64) (x, y int) {
	x = a.Config.Width - int(math.Trunc(((config.MaximaX-newX)/
		(config.MaximaX-config.MinimaX))*float64(a.Config.Width)))
	y = a.Config.Height - int(math.Trunc(((config.MaximaY-newY)/
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
	return 0 <= x && x < a.Config.Width && 0 <= y && y < a.Config.Height
}

// Coloring colors cell is there were no hits by
// taking color from a corresponding transformation.
func (a *Application) Coloring(x, y int, linearTransformation *linear.Linear) {
	a.Pixels[x][y].Color.Red = linearTransformation.Color.Red
	a.Pixels[x][y].Color.Green = linearTransformation.Color.Green
	a.Pixels[x][y].Color.Blue = linearTransformation.Color.Blue
}

// Recoloring recolors cell if there where hits by
// combining cell's and transformation's colors.
func (a *Application) Recoloring(x, y int, linearTransformation *linear.Linear) {
	a.Pixels[x][y].Color.Red = (a.Pixels[x][y].Color.Red +
		linearTransformation.Color.Red) / 2
	a.Pixels[x][y].Color.Green = (a.Pixels[x][y].Color.Green +
		linearTransformation.Color.Green) / 2
	a.Pixels[x][y].Color.Blue = (a.Pixels[x][y].Color.Blue +
		linearTransformation.Color.Blue) / 2
}

// Correction performs brightness correction on
// the number of hits per pixel.
func (a *Application) Correction() {
	maxima, gamma := 0.0, 2.2

	for i := range a.Config.Width {
		for j := range a.Config.Height {
			if a.Pixels[i][j].Hits != 0 {
				a.Pixels[i][j].Normal = math.Log10(float64(a.Pixels[i][j].Hits))
				maxima = max(maxima, a.Pixels[i][j].Normal)
			}
		}
	}

	for i := range a.Config.Width {
		for j := range a.Config.Height {
			a.Pixels[i][j].Normal /= maxima
			a.CorrectCell(i, j, gamma)
		}
	}
}

// CorrectCell performs color correction with
// respecet to the gamma coefficient.
func (a *Application) CorrectCell(i, j int, gamma float64) {
	a.Pixels[i][j].Color.Red = int(float64(a.Pixels[i][j].Color.Red) *
		math.Pow(a.Pixels[i][j].Normal, 1.0/gamma))
	a.Pixels[i][j].Color.Green = int(float64(a.Pixels[i][j].Color.Green) *
		math.Pow(a.Pixels[i][j].Normal, 1.0/gamma))
	a.Pixels[i][j].Color.Blue = int(float64(a.Pixels[i][j].Color.Blue) *
		math.Pow(a.Pixels[i][j].Normal, 1.0/gamma))
}

// ProcessCell implements logic of coloring the
// provided point with checking all necessary
// condition.
func (a *Application) ProcessCell(newX, newY float64) {
	for i := -20; i < a.Config.Iterations; i++ {
		linearTransformation := a.SelectLinearTransformation()
		newX, newY = linearTransformation.Transform(newX, newY)

		nonlinearTransformation := a.SelectNonlinearTransformation()
		newX, newY = nonlinearTransformation.Transform(newX, newY)

		if i >= 0 && a.InsideBounds(newX, newY) {
			x, y := a.ScaleCoordinates(newX, newY)
			if a.InsideImage(x, y) {
				a.Mu.Lock()

				if a.Pixels[x][y].Hits == 0 {
					a.Coloring(x, y, linearTransformation)
				} else {
					a.Recoloring(x, y, linearTransformation)
				}

				a.Pixels[x][y].Hits++

				a.Mu.Unlock()
			}
		}
	}
}
