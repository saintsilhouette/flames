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
func New(cfg *config.Config) (*Application, error) {
	pixels := make([][]*pixel.Pixel, cfg.Width)
	for i := range cfg.Width {
		pixels[i] = make([]*pixel.Pixel, cfg.Height)
		for j := range cfg.Height {
			pixels[i][j] = pixel.New()
		}
	}

	manager, err := fileroutine.New()
	if err != nil {
		return nil, err
	}

	return &Application{
		Pixels:  pixels,
		Config:  cfg,
		Manager: manager,
	}, nil
}

// Render creates the actual fractal flames image.
func (a *Application) Render() error {
	a.FillTransformations()

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
			a.ProcessCell(newX, newY)
		}()
	}

	wg.Wait()

	a.Correction()

	if err := a.Manager.CreateImageFile(a.ConstructImage()); err != nil {
		return err
	}

	return nil
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
