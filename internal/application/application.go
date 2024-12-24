package application

import (
	"image"
	"sync"

	"github.com/es-debug/backend-academy-2024-go-template/config"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/pixel"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/transformations/linear"
	"github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure/fileroutine"
	"github.com/es-debug/backend-academy-2024-go-template/pkg"
)

// Application represents the main program entity.
type Application struct {
	Linear    []*linear.Linear
	Nonlinear []Transformation
	Pixels    [][]*pixel.Pixel

	Config  *config.Config
	Manager Manager
}

// New instantities a new Application entity.
func New(cfg *config.Config) *Application {
	pixels := make([][]*pixel.Pixel, cfg.Width)
	for i := range cfg.Width {
		pixels[i] = make([]*pixel.Pixel, cfg.Height)
		for j := range cfg.Height {
			pixels[i][j] = pixel.New()
		}
	}

	manager := fileroutine.New(cfg.Directory)

	return &Application{
		Pixels:  pixels,
		Config:  cfg,
		Manager: manager,
	}
}

// Render creates the actual fractal flames image.
func (a *Application) Render() error {
	a.fillTransformations()

	wg := &sync.WaitGroup{}
	guard := make(chan struct{}, a.Config.Goroutines)

	for range a.Config.Samples {
		newX := pkg.GetRandomFloat64(config.MinimaX, config.MaximaX)
		newY := pkg.GetRandomFloat64(config.MinimaY, config.MaximaY)

		wg.Add(1)
		guard <- struct{}{}

		go func() {
			defer wg.Done()
			defer func() { <-guard }()
			a.processCell(newX, newY)
		}()
	}

	wg.Wait()

	a.correction()

	if err := a.Manager.CreateImageFile(a.constructImage()); err != nil {
		return err
	}

	return nil
}

// SyntheticRender implemented in order to measure the
// efficiency of concurrent rendering process.
func (a *Application) SyntheticRender() {
	a.fillTransformations()

	wg := &sync.WaitGroup{}
	guard := make(chan struct{}, a.Config.Goroutines)

	for range a.Config.Samples {
		newX := pkg.GetRandomFloat64(config.MinimaX, config.MaximaX)
		newY := pkg.GetRandomFloat64(config.MinimaY, config.MaximaY)

		wg.Add(1)
		guard <- struct{}{}

		go func() {
			defer wg.Done()
			defer func() { <-guard }()
			a.processCell(newX, newY)
		}()
	}

	wg.Wait()

	a.correction()
}

// Manager provides the interface for saving images.
type Manager interface {
	CreateImageFile(*image.RGBA) error
}

// Transformation provides the interface for transformations.
type Transformation interface {
	Transform(float64, float64) (float64, float64)
	Weight() int
}

// IOAdapter provides the interface to work with i/o.
type IOAdapter interface {
	Output(string)
}
