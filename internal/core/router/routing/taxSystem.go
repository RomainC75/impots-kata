package routing

import (
	"impots/internal/application"
	routes "impots/internal/core/router/controllers"
	"impots/internal/domain"
	infra "impots/internal/infrastructure"
	"net/http"

	"github.com/google/uuid"
)

func setupPayments() domain.Payments {
	userUuid := uuid.MustParse("45c971a4-5aeb-40e8-ba51-0f6698e92528")
	payments := infra.NewInMemoryPayments()
	revenu, _ := domain.NewRevenu(21000)
	payment := domain.NewPayment(userUuid, revenu)
	payment.AddPayedTaxe(domain.NewMontant(0))
	payments.ExpectedPayement = *payment

	return payments
}

func TaxeRoutes() http.Handler {
	taxesRoutes := http.NewServeMux()

	taxSystem := application.NewTaxSystem(setupPayments())
	ctrl := routes.NewTaxCtrl(taxSystem)

	taxesRoutes.HandleFunc("POST /taxes", ctrl.TaxCalculation)

	return taxesRoutes
}
