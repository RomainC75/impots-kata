package applications

import (
	"impots/modules/impots/domain"

	"github.com/google/uuid"
)

type CalculateImpotsServiceRequest struct {
	Payslip    domain.Revenu
	User       uuid.UUID
	Reductions []domain.ReductionParameters
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
	reductionHandler, err := domain.NewReductionsHandler(cisRequest.Reductions)
	if err != nil {
		return CalculateImpotsServiceResponse{}, err
	}
	taxeCalculator := domain.NewTaxCalculator(foundUser.Getprepaid(), cisRequest.Payslip, reductionHandler)
	totalTaxe := taxeCalculator.CalculateTaxeToPay()

	return CalculateImpotsServiceResponse{
		TaxeBase:          domain.TaxeBaseMontantFromRevenu(cisRequest.Payslip).ToFloat(),
		AlreadyPayedTaxes: foundUser.Getprepaid().ToFloat(),
		ToBePayedTaxes:    totalTaxe.ToFloat(),
	}, nil
}
