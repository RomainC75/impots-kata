package application

import (
	"impots/internal/domain"

	"github.com/google/uuid"
)

type CalculateTaxRequest struct {
	UserId       uuid.UUID                      `json:"user_id"`
	TaxReduction []domain.TaxReductionBasicInfo `json:"tax_reductions"`
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

	taxeReduction, err := domain.NewTaxeReductions(ctReq.TaxReduction)
	if err != nil {
		return CalculateTaxResponse{}, nil
	}

	return CalculateTaxResponse{
		TaxableBase:     foundPayment.GetTaxableBase().ToFloat(),
		Alreadypayedtax: foundPayment.GetAlreadyPayedTaxe().ToFloat(),
		ToBePayed:       tt.GetTotalTaxe(taxeReduction).Sub(alreadypayedtax).ToFloat(),
	}, nil
}
