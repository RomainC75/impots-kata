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

func (t Taxe) Add(other Taxe) Taxe {
	otherMontant := other.ToMontant()
	return t.Montant.Add(otherMontant).ToTaxe()
}

func (t Taxe) Sub(other Taxe) Taxe {
	otherMontant := other.Montant
	result := t.Montant.Sub(otherMontant)
	if result.IsNegative() {
		return NewTaxe(0)
	}
	return result.ToTaxe()
}
