package application

import (
	"impots/domain"

	"github.com/google/uuid"
)

type CalculateTaxRequest struct {
	UserId uuid.UUID
}

type TaxSystem struct {
	payments domain.Payments
}

func NewTaxSystem(payments domain.Payments) *TaxSystem {
	return &TaxSystem{
		payments: payments,
	}
}

func (ts *TaxSystem) CalculateTax(ctReq CalculateTaxRequest) (float64, error) {
	foundPayment, err := ts.payments.ForUser(ctReq.UserId)
	if err != nil {
		return 0, err
	}
	alreadypayedtax := foundPayment.GetAlreadyPayedTaxe()

	tt := domain.NewTaxeTranches(&foundPayment).SetTranches().Calculate()

	return tt.GetTotalTaxe().Sub(alreadypayedtax).ToFloat(), nil
}
