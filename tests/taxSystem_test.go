package domain_test

import (
	"fmt"
	"impots/application"
	"impots/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TaxSystemTestCases struct {
	paySlip  float64
	expected float64
}

var tcs = []TaxSystemTestCases{
	{5000, 0},
	{12_000, 200},
	{21_000, 1180},
	{31_000, 3050},
	{54_000, 9000},
}

func TestTaxSystem(t *testing.T) {
	t.Run("calculate tax", func(t *testing.T) {
		for _, tc := range tcs {
			ts := application.NewTaxSystem()
			tax, err := ts.CalculateTax(tc.paySlip)

			assert.NoError(t, err)

			assert.Equal(t, tc.expected, tax)
		}
	})

	t.Run("expect an Error", func(t *testing.T) {
		ts := application.NewTaxSystem()
		_, err := ts.CalculateTax(-1)

		assert.Error(t, err, fmt.Errorf(domain.ERROR_NEGATIVE_VALEUR))
	})
}
