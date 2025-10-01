package reduction_domain

import (
	"errors"
	taxe_domain "impots/internal/modules/impots/domain/taxe"
)

var ErrInvalidReductionType = errors.New("invalid reduction type (FIXE, PERCENT)")

type ReductionParameters struct {
	RType string
	Value float64
}

type ReductionsHandler struct {
	reductions []Reduction
}

func NewReductionsHandler(reductionParams []ReductionParameters) (ReductionsHandler, error) {
	filterReductionParams := filterLowestPercents(reductionParams)
	reductions, err := reductionsFactory(filterReductionParams)
	if err != nil {
		return ReductionsHandler{}, err
	}
	return ReductionsHandler{
		reductions: reductions,
	}, nil
}

func (rh ReductionsHandler) ApplyReductions(taxe taxe_domain.Taxe) taxe_domain.Taxe {
	for _, r := range rh.reductions {
		taxe = r.Apply(taxe)

	}
	return taxe
}

func filterLowestPercents(reductionParameters []ReductionParameters) []ReductionParameters {
	biggestPercentIndex := findBiggestPercent(reductionParameters)
	if biggestPercentIndex == -1 {
		return reductionParameters
	}
	filteredReductions := make([]ReductionParameters, 0, len(reductionParameters)-1)
	for i, rp := range reductionParameters {
		if rp.RType == PERCENT_TAX_REDUCTION_TYPE && i != biggestPercentIndex {
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

			fixedReduction, err := NewFixedReduction(r.Value)
			if err != nil {
				return []Reduction{}, err
			}
			reductions = append(reductions, fixedReduction)
		} else if r.RType == PERCENT_TAX_REDUCTION_TYPE {

			percentReduction, err := NewPercentTaxReduction(r.Value)
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
