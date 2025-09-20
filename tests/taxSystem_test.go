package domain_test

import (
	"impots/application"
	"impots/domain"
	infra "impots/infrastructure"
	"testing"

	"github.com/google/uuid"
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
			userUuid := uuid.MustParse("45c971a4-5aeb-40e8-ba51-0f6698e92528")
			inMemoryPayments := infra.InMemoryPayments{}
			revenu, _ := domain.NewRevenu(tc.paySlip)
			payment := domain.NewPayment(userUuid, revenu)
			inMemoryPayments.ExpectedPayement = *payment

			ts := application.NewTaxSystem(&inMemoryPayments)
			tax, err := ts.CalculateTax(application.CalculateTaxRequest{})

			assert.NoError(t, err)

			assert.Equal(t, tc.expected, tax)
		}
	})

}
