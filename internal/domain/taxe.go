package domain

type Taxe struct {
	percent float64
}

func NewTaxe(percent float64) Taxe {
	return Taxe{
		percent: percent,
	}
}

func (t Taxe) ApplyToMontant(montant Montant) Montant {
	return NewMontant(montant.valeur * t.percent / 100)
}
