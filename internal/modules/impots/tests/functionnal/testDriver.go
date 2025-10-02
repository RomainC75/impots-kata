package functionnal

import (
	"impots/internal/modules/impots/applications"
	"impots/internal/modules/impots/domain/entrepreneur"
	taxe_domain "impots/internal/modules/impots/domain/taxe"
	user_domain "impots/internal/modules/impots/domain/users"
	"impots/internal/modules/impots/infrastructure"
	"time"

	"github.com/google/uuid"
)

var (
	newCompanyId = uuid.MustParse("575c4d3d-0ed4-4192-9734-b3dbb9b80476")
	oldCompanyId = uuid.MustParse("60fc0f78-d6e7-4f21-a5e5-73670d0518ef")
)

func taxSystemTestDriver(alreadyPayed float64) *applications.TaxSystem {
	inMemoryTimeProvider := infrastructure.NewInMemoryTime(time.Date(2025, 6, 1, 0, 0, 0, 0, time.UTC))

	user := user_domain.NewUser(uuid.MustParse("699117b4-7df1-4e14-95e4-5912a6564ef9"))
	user.PayTaxe(taxe_domain.NewTaxe(alreadyPayed))

	inMemoryEntrepreneurs := infrastructure.NewInMemoryEntrepreneurs()
	entrepreneurId := uuid.MustParse("799117b4-7df1-4e14-95e4-5912a6564ef9")
	companies := []entrepreneur.Company{
		entrepreneur.Company{oldCompanyId, time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)},
		entrepreneur.Company{newCompanyId, time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)},
	}
	inMemoryEntrepreneurs.ExpectedEntrepreneur, _ = *entrepreneur.NewEntrepreneur(entrepreneurId, user.GetID(), companies)

	inMemoryUsers := infrastructure.NewInMemoryUsers()
	inMemoryUsers.ExpectedUser = *user

	return applications.NewTaxSystem(inMemoryUsers, inMemoryEntrepreneurs, inMemoryTimeProvider)
}
