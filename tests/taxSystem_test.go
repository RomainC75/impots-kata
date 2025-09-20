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
	expected application.CalculateTaxResponse
}

var tcs = []TaxSystemTestCases{
	{5000, application.CalculateTaxResponse{
		TaxableBase:     0,
		Alreadypayedtax: 0,
		ToBePayed:       0,
	}},
	{12_000, application.CalculateTaxResponse{
		TaxableBase:     2_000,
		Alreadypayedtax: 0,
		ToBePayed:       200,
	}},
	{21_000, application.CalculateTaxResponse{
		TaxableBase:     11_000,
		Alreadypayedtax: 0,
		ToBePayed:       1180,
	}},
	{31_000, application.CalculateTaxResponse{
		TaxableBase:     21_000,
		Alreadypayedtax: 0,
		ToBePayed:       3050,
	}},
	{54_000, application.CalculateTaxResponse{
		TaxableBase:     44_000,
		Alreadypayedtax: 0,
		ToBePayed:       9000,
	}},
}

func testDriver(paySlip float64) (*application.TaxSystem, uuid.UUID) {
	userUuid := uuid.MustParse("45c971a4-5aeb-40e8-ba51-0f6698e92528")
	inMemoryPayments := infra.InMemoryPayments{}
	revenu, _ := domain.NewRevenu(paySlip)
	payment := domain.NewPayment(userUuid, revenu)
	inMemoryPayments.ExpectedPayement = *payment

	ts := application.NewTaxSystem(&inMemoryPayments)
	return ts, userUuid

}

func TestTaxSystem(t *testing.T) {
	t.Run("calculate tax", func(t *testing.T) {
		for _, tc := range tcs {
			ts, userId := testDriver(tc.paySlip)
			tax, err := ts.CalculateTax(application.CalculateTaxRequest{
				UserId: userId,
			})

			assert.NoError(t, err)

			assert.Equal(t, tc.expected, tax)
		}
	})

}
