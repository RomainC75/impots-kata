package routes

import (
	"encoding/json"
	"impots/internal/application"
	"net/http"
)

type TaxCtrl struct {
	taxSystemApp *application.TaxSystem
}

func NewTaxCtrl(taxSystemApp *application.TaxSystem) *TaxCtrl {
	return &TaxCtrl{
		taxSystemApp: taxSystemApp,
	}
}

func (tsc *TaxCtrl) TaxCalculation(w http.ResponseWriter, r *http.Request) {

	req := application.CalculateTaxRequest{}

	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		// send error
	}

	res, err := tsc.taxSystemApp.CalculateTax(req)
	if err != nil {
		// send error
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
