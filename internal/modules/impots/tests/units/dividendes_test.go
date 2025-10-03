package units

import (
	dividendes "impots/internal/modules/impots/domain/dividende"
	money_domain "impots/internal/modules/impots/domain/money"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func DividendeTestDriver() {

}

func TestDividendes(t *testing.T) {
	tc := []struct {
		Dividendes []dividendes.Dividende
		Expected   float64
	}{
		{
			Dividendes: []dividendes.Dividende{
				dividendes.NewDividende(uuid.MustParse("76188e2f-d101-4b25-8259-f4bbce8907a4"), money_domain.NewRevenu(2000)),
				dividendes.NewDividende(uuid.MustParse("654b0460-641a-4064-810e-15e391a58c27"), money_domain.NewRevenu(6000)),
			},
			Expected: 3600,
		},
	}
	t.Run("dividendes", func(t *testing.T) {
		dividendes := dividendes.NewDividendes(tc[0].Dividendes)
		totalTaxe := dividendes.CalculateTotalTaxes()
		assert.Equal(t, tc[0].Expected, totalTaxe.ToFloat())
	})
}
