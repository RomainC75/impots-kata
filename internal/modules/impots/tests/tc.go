package tests

import (
	reduction_domain "impots/internal/modules/impots/domain/reduction"
)

type TranchesTC struct {
	Revenu            float64
	AlreadyPayed      float64
	Reductions        []reduction_domain.ReductionParameters
	ExpectedTaxeToPay float64
	ExpectedTaxeBase  float64
}

var (
	tenPercent    = reduction_domain.ReductionParameters{"PERCENT", 10}
	twentyPercent = reduction_domain.ReductionParameters{"PERCENT", 20}
	minus200      = reduction_domain.ReductionParameters{"FIXE", 200}
	minus1000     = reduction_domain.ReductionParameters{"FIXE", 1000}
)

var Tcs []TranchesTC = []TranchesTC{
	{10_000, 0, []reduction_domain.ReductionParameters{}, 0, 0},
	{15_000, 0, []reduction_domain.ReductionParameters{}, 500, 5_000},
	{20_000, 0, []reduction_domain.ReductionParameters{}, 1000, 10_000},
	{30_000, 0, []reduction_domain.ReductionParameters{}, 2800, 20_000},
	{35_000, 0, []reduction_domain.ReductionParameters{}, 4050, 25_000},
	{40_000, 0, []reduction_domain.ReductionParameters{}, 5300, 30_000},
	{50_000, 0, []reduction_domain.ReductionParameters{}, 7800, 40_000},
	{60_000, 0, []reduction_domain.ReductionParameters{}, 10_800, 50_000},
	// prepayed
	{10_000, 300, []reduction_domain.ReductionParameters{}, 0, 0},
	{15_000, 200, []reduction_domain.ReductionParameters{}, 300, 5_000},
	{20_000, 600, []reduction_domain.ReductionParameters{}, 400, 10_000},
	{30_000, 1000, []reduction_domain.ReductionParameters{}, 1800, 20_000},
	// prepayed + reductions
	{10_000, 300, []reduction_domain.ReductionParameters{tenPercent, minus200}, 0, 0},
	{20_000, 600, []reduction_domain.ReductionParameters{tenPercent, minus200}, 160, 10_000},
	{20_000, 600, []reduction_domain.ReductionParameters{tenPercent, twentyPercent, minus200}, 120, 10_000},

	// edge cases
	{15_000, 0, []reduction_domain.ReductionParameters{minus1000, tenPercent}, 0, 5_000},
	// {15_000, 0, []reduction_domain.ReductionParameters{tenPercent, twentyPercent, tenPercent, twentyPercent}, 0, 5_000},
}
