package linear_test

import (
	"testing"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/transformations/linear"
	"github.com/es-debug/backend-academy-2024-go-template/pkg"
	"github.com/stretchr/testify/assert"
)

func TestSinusoidalTransformation(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name        string
		multipliers *linear.Coefficients
		x, y        float64
		expx, expy  float64
	}

	testCases := []testCase{
		{
			name:        "identity transformation",
			multipliers: &linear.Coefficients{A: 1, B: 0, C: 0, D: 0, E: 1, F: 0},
			x:           1.0, y: 1.0,
			expx: 1.0, expy: 1.0,
		},
		{
			name:        "regular transformation",
			multipliers: &linear.Coefficients{A: 1, B: 2, C: 3, D: 4, E: 5, F: 6},
			x:           1.0, y: 1.0,
			expx: 6.0, expy: 15.0,
		},
	}

	for _, tc := range testCases {
		tc := tc

		transformation := linear.New(0)
		transformation.Multipliers = tc.multipliers

		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()

			newX, newY := transformation.Transform(tc.x, tc.y)

			assert.Equal(tt, pkg.TruncateFloat(newX, 5), tc.expx)
			assert.Equal(tt, pkg.TruncateFloat(newY, 5), tc.expy)
		})
	}
}
