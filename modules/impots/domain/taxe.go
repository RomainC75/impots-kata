package domain

type Taxe struct {
	Montant
}

func NewTaxe(value float64) Taxe {
	return Taxe{
		Montant: NewMontant(value),
	}
}

func (t Taxe) ToMontant() Montant {
	return t.Montant
}
