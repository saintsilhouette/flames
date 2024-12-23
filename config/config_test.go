package config_test

import (
	"testing"

	"github.com/es-debug/backend-academy-2024-go-template/config"
	"github.com/stretchr/testify/assert"
)

func TestApplicationInit(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name       string
		height     uint
		width      uint
		samples    uint
		iterations uint
		goroutines uint
		symmetry   uint
		directory  string
		err        error
	}

	testCases := []testCase{
		{
			name:  "width overflow test",
			width: 7681,
			err:   config.WidthValueOverflow,
		},
		{
			name:   "height overflow test",
			width:  7680,
			height: 4321,
			err:    config.HeightValueOverflow,
		},
		{
			name:     "symmetry overflow test",
			width:    1920,
			height:   1080,
			symmetry: 7,
			err:      config.SymmetryValueOverflow,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			_, err := config.New(tc.width, tc.height, tc.samples, tc.iterations, tc.goroutines, tc.symmetry, tc.directory)

			a := assert.New(tt)
			if tc.err != nil {
				a.EqualError(err, tc.err.Error())

				return
			}
		})
	}
}

type IOAdapter interface {
	Output(string)
}
