package applications

import (
	"impots/modules/impots/domain"

	"github.com/google/uuid"
)

type CalculateImpotsServiceRequest struct {
	Payslip float64
	User    uuid.UUID
}

type TaxSystem struct {
	usersRepo domain.Users
}

func NewTaxSystem(usersRepo domain.Users) *TaxSystem {
	return &TaxSystem{
		usersRepo: usersRepo,
	}
}

func (cis *TaxSystem) CalculateTax(cisRequest CalculateImpotsServiceRequest) (float64, error) {
	foundUser, err := cis.usersRepo.GetUser(cisRequest.User)
	if err != nil {
		return 0, err
	}

	taxeCalculator := domain.NewTaxCalculator()
	return taxeCalculator.CalculateTaxeToPay(foundUser, cisRequest.Payslip).ToFloat(), nil
}
