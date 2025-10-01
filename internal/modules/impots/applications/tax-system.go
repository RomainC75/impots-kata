package applications

import (
	"impots/internal/modules/impots/domain"
	money_domain "impots/internal/modules/impots/domain/money"
	reduction_domain "impots/internal/modules/impots/domain/reduction"
	taxe_domain "impots/internal/modules/impots/domain/taxe"
	user_domain "impots/internal/modules/impots/domain/users"

	"github.com/google/uuid"
)

type CalculateImpotsServiceRequest struct {
	RevenuSalarie money_domain.Revenu
	User          uuid.UUID
	Reductions    []reduction_domain.ReductionParameters
}

type CalculateImpotsServiceResponse struct {
	TaxeBase          float64
	AlreadyPayedTaxes float64
	ToBePayedTaxes    float64
}

type TaxSystem struct {
	usersRepo user_domain.Users
}

func NewTaxSystem(usersRepo user_domain.Users) *TaxSystem {
	return &TaxSystem{
		usersRepo: usersRepo,
	}
}

func (cis *TaxSystem) CalculateTax(cisRequest CalculateImpotsServiceRequest) (CalculateImpotsServiceResponse, error) {
	foundUser, err := cis.usersRepo.GetUser(cisRequest.User)
	if err != nil {
		return CalculateImpotsServiceResponse{}, err
	}
	reductionHandler, err := reduction_domain.NewReductionsHandler(cisRequest.RevenuSalarie, cisRequest.Reductions)
	if err != nil {
		return CalculateImpotsServiceResponse{}, err
	}
	taxeCalculator := domain.NewTaxCalculator(foundUser.Getprepaid(), cisRequest.RevenuSalarie, reductionHandler)
	totalTaxe := taxeCalculator.CalculateTaxeToPay()

	return CalculateImpotsServiceResponse{
		TaxeBase:          taxe_domain.TaxeBaseMontantFromRevenu(cisRequest.RevenuSalarie).ToFloat(),
		AlreadyPayedTaxes: foundUser.Getprepaid().ToFloat(),
		ToBePayedTaxes:    totalTaxe.ToFloat(),
	}, nil
}
