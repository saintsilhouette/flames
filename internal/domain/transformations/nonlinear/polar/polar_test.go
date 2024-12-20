package polar_test

import (
	"testing"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/transformations/nonlinear/polar"
	"github.com/es-debug/backend-academy-2024-go-template/pkg"
	"github.com/stretchr/testify/assert"
)

func TestPolarTransformation(t *testing.T) {
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
			expx: 0.25, expy: 0.41421,
		},
		{
			name: "transformation with coefficients with different signs",
			x:    1.0, y: -1.0,
			expx: -0.25, expy: 0.41421,
		},
	}

	for _, tc := range testCases {
		tc := tc

		transformation := polar.New(0)

		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()

			newX, newY := transformation.Transform(tc.x, tc.y)

			assert.Equal(tt, tc.expx, pkg.TruncateFloat(newX, 5))
			assert.Equal(tt, tc.expy, pkg.TruncateFloat(newY, 5))
		})
	}
}
