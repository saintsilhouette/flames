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
		height     string
		width      string
		samples    string
		iterations string
		goroutines string
		err        error
	}

	testCases := []testCase{
		{
			name:  "ascii width test",
			width: "a",
			err:   config.InvalidWidthError,
		},
		{
			name:  "negative width test",
			width: "-1",
			err:   config.NegativeWidthError,
		},
		{
			name:   "ascii height test",
			width:  "1",
			height: "a",
			err:    config.InvalidHeightError,
		},
		{
			name:   "negative height test",
			width:  "1",
			height: "-1",
			err:    config.NegativeHeightError,
		},
		{
			name:    "ascii samples test",
			width:   "1",
			height:  "1",
			samples: "a",
			err:     config.InvalidSamplesError,
		},
		{
			name:    "negative samples test",
			width:   "1",
			height:  "1",
			samples: "-1",
			err:     config.NegativeSamplesError,
		},
		{
			name:       "ascii iterations test",
			width:      "1",
			height:     "1",
			samples:    "1",
			iterations: "a",
			err:        config.InvalidIterationsError,
		},
		{
			name:       "negative iterations test",
			width:      "1",
			height:     "1",
			samples:    "1",
			iterations: "-1",
			err:        config.NegativeIterationsError,
		},
		{
			name:       "ascii goroutines test",
			width:      "1",
			height:     "1",
			samples:    "1",
			iterations: "1",
			goroutines: "a",
			err:        config.InvalidGoroutinesError,
		},
		{
			name:       "negative goroutines test",
			width:      "1",
			height:     "1",
			samples:    "1",
			iterations: "1",
			goroutines: "-1",
			err:        config.NegativeGoroutinesError,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			_, err := config.New(tc.width, tc.height, tc.samples, tc.iterations, tc.goroutines)

			a := assert.New(tt)
			if tc.err != nil {
				a.EqualError(err, tc.err.Error())

				return
			}
		})
	}
}

type IOAdapter interface {
	Input() string
	Output(string)
}
