package application

import (
	"impots/domain"

	"github.com/google/uuid"
)

type CalculateTaxRequest struct {
	UserId uuid.UUID
}

type CalculateTaxResponse struct {
	TaxableBase     float64
	Alreadypayedtax float64
	ToBePayed       float64
}

type TaxSystem struct {
	payments domain.Payments
}

func NewTaxSystem(payments domain.Payments) *TaxSystem {
	return &TaxSystem{
		payments: payments,
	}
}

func (ts *TaxSystem) CalculateTax(ctReq CalculateTaxRequest) (CalculateTaxResponse, error) {
	foundPayment, err := ts.payments.ForUser(ctReq.UserId)
	if err != nil {
		return CalculateTaxResponse{}, err
	}
	alreadypayedtax := foundPayment.GetAlreadyPayedTaxe()

	tt := domain.NewTaxeTranches(&foundPayment).SetTranches().Calculate()

	return CalculateTaxResponse{
		TaxableBase:     foundPayment.GetTaxableBase().ToFloat(),
		Alreadypayedtax: foundPayment.GetAlreadyPayedTaxe().ToFloat(),
		ToBePayed:       tt.GetTotalTaxe().Sub(alreadypayedtax).ToFloat(),
	}, nil
}
