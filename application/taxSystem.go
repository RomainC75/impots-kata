package application

import (
	"impots/domain"

	"github.com/google/uuid"
)

type CalculateTaxRequest struct {
	userId uuid.UUID
}

type TaxSystem struct {
	users domain.Users
}

func NewTaxSystem(users domain.Users) *TaxSystem {
	return &TaxSystem{
		users: users,
	}
}

func (ts *TaxSystem) CalculateTax(ctReq CalculateTaxRequest) (float64, error) {
	foundUser, err := ts.users.GetUser(ctReq.userId)
	if err != nil {
		return 0, err
	}
	alreadypayedtax := foundUser.GetAlreadyPayedTaxe()

	tt := domain.NewTaxeTranches(foundUser).SetTranches().Calculate()

	return tt.GetTotalTaxe().Sub(alreadypayedtax).ToFloat(), nil
}
