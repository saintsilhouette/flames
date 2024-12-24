package application_test

import (
	"testing"

	"github.com/es-debug/backend-academy-2024-go-template/config"
	"github.com/es-debug/backend-academy-2024-go-template/internal/application"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/transformations/linear"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/transformations/nonlinear/heart"
)

func benchmarkRender(b *testing.B) {
	cfg, _ := config.New(1920, 1080, 2048, 32, uint(b.N), 0, config.Directory) //nolint
	app := application.New(cfg)

	app.Linear = []*linear.Linear{
		linear.New(0), linear.New(0), linear.New(0), linear.New(0), linear.New(0),
	}

	app.Nonlinear = []application.Transformation{
		heart.New(0), heart.New(0), heart.New(0), heart.New(0), heart.New(0),
	}

	app.SyntheticRender()
}

func BenchmarkRender1(b *testing.B) {
	benchmarkRender(b)
}

func BenchmarkRender8(b *testing.B) {
	benchmarkRender(b)
}

func BenchmarkRender16(b *testing.B) {
	benchmarkRender(b)
}

func BenchmarkRender32(b *testing.B) {
	benchmarkRender(b)
}

func BenchmarkRender64(b *testing.B) {
	benchmarkRender(b)
}

func BenchmarkRender128(b *testing.B) {
	benchmarkRender(b)
}

func BenchmarkRender256(b *testing.B) {
	benchmarkRender(b)
}

func BenchmarkRender512(b *testing.B) {
	benchmarkRender(b)
}

func BenchmarkRender1024(b *testing.B) {
	benchmarkRender(b)
}

func BenchmarkRender2048(b *testing.B) {
	benchmarkRender(b)
}
