package reduction_domain

import (
	"errors"
	money_d "impots/internal/modules/impots/domain/money"
	taxe_d "impots/internal/modules/impots/domain/taxe"
)

var (
	ErrInvalidReductionType         = errors.New("invalid reduction type (FIXE, PERCENT)")
	MaxReduction            float64 = 1_271
)

type ReductionParameters struct {
	RType          string
	Value          float64
	ApplicableFrom float64
}

type ReductionsHandler struct {
	reductions []Reduction
}

func NewReductionsHandler(revenu money_d.Revenu, reductionParams []ReductionParameters) (ReductionsHandler, error) {
	filterReductionParams := filterLowestPercents(revenu, reductionParams)
	reductions, err := reductionsFactory(filterReductionParams)
	if err != nil {
		return ReductionsHandler{}, err
	}
	return ReductionsHandler{
		reductions: reductions,
	}, nil
}

func (rh ReductionsHandler) ApplyReductions(revenu money_d.Revenu, taxe taxe_d.Taxe) taxe_d.Taxe {
	buffTaxe := taxe

	for _, r := range rh.reductions {
		buffTaxe = r.Apply(revenu, buffTaxe)

	}
	limitedReductionTaxe := limitReductionTaxe(taxe, buffTaxe)
	return limitedReductionTaxe
}

func limitReductionTaxe(preReductionTaxe taxe_d.Taxe, buffTaxe taxe_d.Taxe) taxe_d.Taxe {
	totalTaxeReduction := preReductionTaxe.Sub(buffTaxe)
	if totalTaxeReduction.IsMoreThan(money_d.NewMontant(MaxReduction)) {
		return preReductionTaxe.Sub(taxe_d.NewTaxe(MaxReduction))
	}
	return buffTaxe
}

func filterLowestPercents(revenu money_d.Revenu, reductionParameters []ReductionParameters) []ReductionParameters {
	biggestPercentIndex := findBiggestPercent(reductionParameters)
	if biggestPercentIndex == -1 {
		return reductionParameters
	}
	filteredReductions := make([]ReductionParameters, 0, len(reductionParameters)-1)
	for i, rp := range reductionParameters {
		if revenu.IsLessThan(money_d.NewMontant(rp.ApplicableFrom)) {
			continue
		} else if rp.RType == PERCENT_TAX_REDUCTION_TYPE && i != biggestPercentIndex {
			continue
		} else if i == biggestPercentIndex {
			filteredReductions = append([]ReductionParameters{rp}, filteredReductions...)
		} else {
			filteredReductions = append(filteredReductions, rp)
		}
	}
	return filteredReductions
}

func findBiggestPercent(reductionParams []ReductionParameters) int {
	biggestPercentIndex := -1
	var biggestPercentRate float64 = 0
	for i, rp := range reductionParams {

		if rp.RType == PERCENT_TAX_REDUCTION_TYPE && rp.Value > float64(biggestPercentRate) {
			biggestPercentIndex = i
			biggestPercentRate = rp.Value
		}
	}
	return biggestPercentIndex
}

func reductionsFactory(reductionParams []ReductionParameters) ([]Reduction, error) {
	reductions := make([]Reduction, 0, len(reductionParams))
	for _, r := range reductionParams {

		if r.RType == FIXE_REDUCTION_TYPE {

			fixedReduction, err := NewFixedReduction(r.Value, money_d.NewRevenu(r.ApplicableFrom))
			if err != nil {
				return []Reduction{}, err
			}
			reductions = append(reductions, fixedReduction)
		} else if r.RType == PERCENT_TAX_REDUCTION_TYPE {

			percentReduction, err := NewPercentTaxReduction(r.Value, money_d.NewRevenu(r.ApplicableFrom))
			if err != nil {
				return []Reduction{}, err
			}
			reductions = append(reductions, percentReduction)
		} else {

			return []Reduction{}, ErrInvalidReductionType
		}
	}
	return reductions, nil
}
