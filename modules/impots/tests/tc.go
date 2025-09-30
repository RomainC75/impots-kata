package tests

type TranchesTC struct {
	Revenu            float64
	ExpectedTaxeToPay float64
}

var Tcs []TranchesTC = []TranchesTC{
	{10_000, 0},
	{15_000, 500},
	{20_000, 1000},
	{30_000, 2800},
	{35_000, 4050},
	{40_000, 5300},
	{50_000, 7800},
	{60_000, 10_800},
}
