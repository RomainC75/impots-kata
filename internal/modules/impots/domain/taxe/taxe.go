package taxe_domain

import money_domain "impots/internal/modules/impots/domain/money"

type Taxe struct {
	money_domain.Montant
}

func NewTaxe(value float64) Taxe {
	return Taxe{
		Montant: money_domain.NewMontant(value),
	}
}

func NewTaxeFromMontant(montant money_domain.Montant) Taxe {
	return Taxe{
		Montant: montant,
	}
}

func (t Taxe) ToMontant() money_domain.Montant {
	return t.Montant
}

func (t Taxe) Add(other Taxe) Taxe {
	otherMontant := other.ToMontant()
	return NewTaxeFromMontant(t.Montant.Add(otherMontant))
}

func (t Taxe) Sub(other Taxe) Taxe {
	otherMontant := other.Montant
	result := t.Montant.Sub(otherMontant)
	if result.IsNegative() {
		return NewTaxe(0)
	}
	return NewTaxeFromMontant(result)
}

func (t Taxe) MultiplyByValue(value float64) Taxe {
	return NewTaxe(t.ToFloat() * value)
}

func TaxeBaseMontantFromRevenu(revenu money_domain.Revenu) money_domain.Montant {
	taxableThreshold := money_domain.NewMontant(10_000)
	base := revenu.Sub(taxableThreshold)
	if base.IsNegative() {
		return money_domain.NewMontant(0)
	}
	return base
}

func (t Taxe) Round2Decimals() Taxe {
	return NewTaxeFromMontant(t.Montant.Round2Decimals())
}
