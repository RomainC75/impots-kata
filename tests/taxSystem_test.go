package domain_test

import (
	"impots/application"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TaxSystemTestCases struct {
	paySlip  float64
	expected float64
}

var tcs = []TaxSystemTestCases{
	{9_000, 0},
	{11_000, 100},
	{21_000, 1180},
	{31_000, 2050},
	{40_000, 4300},
}

func TestTaxSystem(t *testing.T) {
	t.Run("calculate tax", func(t *testing.T) {
		for _, tc := range tcs {
			ts := application.NewTaxSystem()
			tax := ts.CalculateTax(tc.paySlip)

			assert.Equal(t, tc.expected, tax)
		}
	})
}
