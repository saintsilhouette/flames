package main

import (
	"flag"
	"log/slog"
	"os"

	"github.com/es-debug/backend-academy-2024-go-template/config"
	"github.com/es-debug/backend-academy-2024-go-template/internal/application"
	"github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure/adapter"
)

func main() {
	width := flag.Uint("w", 1920, "image width (<= 7680)")
	height := flag.Uint("h", 1080, "image height (<= 4320)")
	samples := flag.Uint("n", 32768, "number of event loop iterations")
	iterations := flag.Uint("i", 128, "number of iterations per point")
	goroutines := flag.Uint("g", 32, "maximum number of goroutines")
	symmetry := flag.Uint("s", 1, "x-way rotational symmetry, where x is your value (x <= 6)")
	directory := flag.String("d", "images", "the directory to store generated images")

	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
	io := adapter.New(os.Stdout, logger)

	cfg, err := config.New(*width, *height, *samples, *iterations, *goroutines, *symmetry, *directory)
	if err != nil {
		io.Output(err.Error())
		os.Exit(1)
	}

	app := application.New(cfg)

	if err = app.Render(); err != nil {
		io.Output(err.Error())
		os.Exit(1)
	}
}
