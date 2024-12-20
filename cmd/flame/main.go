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
	width := flag.String("w", "1920", "image width")
	height := flag.String("h", "1080", "image height")
	samples := flag.String("s", "32768", "number of event loop iterations")
	iterations := flag.String("i", "512", "number of iterations per point")
	goroutines := flag.String("g", "32", "maximum number of goroutines")
	directory := flag.String("d", "images", "images directory name")

	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
	io := adapter.New(os.Stdout, os.Stdin, logger)

	cfg, err := config.New(*width, *height, *samples, *iterations, *goroutines, *directory)
	if err != nil {
		io.Output(err.Error())
		os.Exit(1)
	}

	app, err := application.New(cfg)
	if err != nil {
		io.Output(err.Error())
		os.Exit(1)
	}

	err = app.Render()
	if err != nil {
		io.Output(err.Error())
		os.Exit(1)
	}
}
