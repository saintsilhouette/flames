package application

import (
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

	Mu    *sync.Mutex
	Wg    *sync.WaitGroup
	Guard chan struct{}

	Config  *config.Config
	Manager *fileroutine.ImagesManager
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
		Mu:      &sync.Mutex{},
		Wg:      &sync.WaitGroup{},
		Guard:   make(chan struct{}, cfg.Goroutines),
		Config:  cfg,
		Manager: manager,
	}, nil
}

// Render creates the actual fractal flames image.
func (a *Application) Render() error {
	a.FillTransformations()

	for range a.Config.Samples {
		newX := pkg.GetRandomFloat64(config.MinimaX, config.MaximaX)
		newY := pkg.GetRandomFloat64(config.MinimaY, config.MaximaY)

		a.Wg.Add(1)
		a.Guard <- struct{}{}

		go func(x, y float64) {
			defer a.Wg.Done()
			defer func() { <-a.Guard }()
			a.ProcessCell(x, y)
		}(newX, newY)
	}

	a.Correction()

	err := a.Manager.CreateImageFile(a.ConstructImage())
	if err != nil {
		return err
	}

	return nil
}

// Transformation provides the general interface for
// the interaction with transformations.
type Transformation interface {
	Transform(float64, float64) (float64, float64)
	Weight() int
}

// IOAdapter provides the interface to work with i/o.
type IOAdapter interface {
	Input() string
	Output(string)
}
