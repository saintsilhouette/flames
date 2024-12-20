package rgb

import (
	"github.com/es-debug/backend-academy-2024-go-template/config"
	"github.com/es-debug/backend-academy-2024-go-template/pkg"
)

// RGB reflects the cell's color by using rgb notation.
type RGB struct {
	Red, Green, Blue int
}

// New instantiates a new Color entity.
func New() *RGB {
	return &RGB{
		Red:   pkg.GetRandomInt(config.ColorLower, config.ColorUpper),
		Green: pkg.GetRandomInt(config.ColorLower, config.ColorUpper),
		Blue:  pkg.GetRandomInt(config.ColorLower, config.ColorUpper),
	}
}
