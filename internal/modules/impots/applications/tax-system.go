package applications

import (
	"impots/internal/modules/impots/domain"
	"impots/internal/modules/impots/domain/entrepreneur"
	money_domain "impots/internal/modules/impots/domain/money"
	reduction_domain "impots/internal/modules/impots/domain/reduction"
	taxe_domain "impots/internal/modules/impots/domain/taxe"
	user_domain "impots/internal/modules/impots/domain/users"

	"github.com/google/uuid"
)

type CalculateImpotsServiceRequest struct {
	RevenuSalarie       money_domain.Revenu
	RevenusByEntreprise []entrepreneur.RevenuByEntreprise
	User                uuid.UUID
	Reductions          []reduction_domain.ReductionParameters
}

type CalculateImpotsServiceResponse struct {
	TaxeBase          float64
	AlreadyPayedTaxes float64
	ToBePayedTaxes    float64
}

type TaxSystem struct {
	usersRepo         user_domain.Users
	entrepreneursRepo entrepreneur.Entrepreneurs
	timeProvider      domain.Time
}

func NewTaxSystem(usersRepo user_domain.Users, entrepreneursRepo entrepreneur.Entrepreneurs, timeProvider domain.Time) *TaxSystem {
	return &TaxSystem{
		usersRepo:         usersRepo,
		entrepreneursRepo: entrepreneursRepo,
		timeProvider:      timeProvider,
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

	entrepreneur, _ := cis.entrepreneursRepo.GetEntrepreneur(foundUser.GetID())

	taxeCalculator := domain.NewTaxCalculator(foundUser.GetprepaidTaxe(), cisRequest.RevenuSalarie, reductionHandler, cisRequest.RevenusByEntreprise, entrepreneur)
	totalTaxe, err := taxeCalculator.CalculateTaxeToPay(cis.timeProvider.Now())
	if err != nil {
		return CalculateImpotsServiceResponse{}, err
	}

	return CalculateImpotsServiceResponse{
		TaxeBase:          taxe_domain.TaxeBaseMontantFromRevenu(cisRequest.RevenuSalarie).ToFloat(),
		AlreadyPayedTaxes: foundUser.GetprepaidTaxe().ToFloat(),
		ToBePayedTaxes:    totalTaxe.ToFloat(),
	}, nil
}
