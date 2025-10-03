package dividendes

import (
	"fmt"
	taxe_domain "impots/internal/modules/impots/domain/taxe"
)

type Dividendes struct {
	dividendes []Dividende
}

func NewDividendes(dividendes []Dividende) Dividendes {
	return Dividendes{
		dividendes: dividendes,
	}
}

func (d Dividendes) CalculateTotalTaxes() taxe_domain.Taxe {
	totalTaxe := taxe_domain.NewTaxe(0)
	for _, dividende := range d.dividendes {
		taxe := dividende.GetTaxe()
		fmt.Println(dividende, taxe)
		totalTaxe = taxe.Add(taxe)
	}
	return totalTaxe
}
