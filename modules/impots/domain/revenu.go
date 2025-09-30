package domain

type Revenu struct {
	Montant
}

func NewRevenu(value float64) Revenu {
	return Revenu{
		Montant: NewMontant(value),
	}
}
