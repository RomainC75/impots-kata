package functionnal

import (
	"impots/modules/impots/applications"
	"impots/modules/impots/domain"
	"impots/modules/impots/infrastructure"

	"github.com/google/uuid"
)

func taxSystemTestDriver(alreadyPayed float64) *applications.TaxSystem {
	user := domain.NewUser(uuid.MustParse("699117b4-7df1-4e14-95e4-5912a6564ef9"))
	user.PayTaxe(domain.NewTaxe(alreadyPayed))

	inMemoryUsers := infrastructure.NewInMemoryUsers()
	inMemoryUsers.ExpectedUser = *user

	return applications.NewTaxSystem(inMemoryUsers)
}
