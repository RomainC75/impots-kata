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

func (t Taxe) MultiplyByValue(value float64) Taxe {
	return NewTaxe(t.value * value)
}

func TaxeBaseMontantFromRevenu(revenu Revenu) Montant {
	taxableThreshold := NewMontant(10_000)
	base := revenu.Sub(taxableThreshold)
	if base.IsNegative() {
		return NewMontant(0)
	}
	return base
}
