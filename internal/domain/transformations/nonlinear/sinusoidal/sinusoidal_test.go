package sinusoidal_test

import (
	"testing"

	"github.com/es-debug/backend-academy-2024-go-template/config"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/transformations/nonlinear/sinusoidal"
	"github.com/stretchr/testify/assert"
)

func TestSinusoidalTransformation(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name       string
		x, y       float64
		expx, expy float64
	}

	testCases := []testCase{
		{
			name: "regular transformation",
			x:    1.0, y: 1.0,
			expx: 0.84147, expy: 0.84147,
		},
		{
			name: "transformation with coefficients with different signs",
			x:    1.0, y: -1.0,
			expx: 0.84147, expy: -0.84147,
		},
	}

	for _, tc := range testCases {
		tc := tc

		transformation := sinusoidal.New(0)

		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()

			newX, newY := transformation.Transform(tc.x, tc.y)

			assert.InDelta(tt, tc.expx, newX, config.Delta)
			assert.InDelta(tt, tc.expy, newY, config.Delta)
		})
	}
}
