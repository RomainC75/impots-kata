package applications

import (
	"impots/modules/impots/domain"

	"github.com/google/uuid"
)

type CalculateImpotsServiceRequest struct {
	Payslip domain.Revenu
	User    uuid.UUID
}

type CalculateImpotsServiceResponse struct {
	TaxeBase          float64
	AlreadyPayedTaxes float64
	ToBePayedTaxes    float64
}

type TaxSystem struct {
	usersRepo domain.Users
}

func NewTaxSystem(usersRepo domain.Users) *TaxSystem {
	return &TaxSystem{
		usersRepo: usersRepo,
	}
}

func (cis *TaxSystem) CalculateTax(cisRequest CalculateImpotsServiceRequest) (CalculateImpotsServiceResponse, error) {
	foundUser, err := cis.usersRepo.GetUser(cisRequest.User)
	if err != nil {
		return CalculateImpotsServiceResponse{}, err
	}

	taxeCalculator := domain.NewTaxCalculator()
	return CalculateImpotsServiceResponse{
		TaxeBase:          domain.TaxeBaseMontantFromRevenu(cisRequest.Payslip).ToFloat(),
		AlreadyPayedTaxes: foundUser.GetPayedTaxe().ToFloat(),
		ToBePayedTaxes:    taxeCalculator.CalculateTaxeToPay(foundUser, cisRequest.Payslip).ToFloat(),
	}, nil
}
