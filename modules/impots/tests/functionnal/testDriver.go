package functionnal

import (
	"impots/modules/impots/applications"
	taxe_domain "impots/modules/impots/domain/taxe"
	user_domain "impots/modules/impots/domain/users"
	"impots/modules/impots/infrastructure"

	"github.com/google/uuid"
)

func taxSystemTestDriver(alreadyPayed float64) *applications.TaxSystem {
	user := user_domain.NewUser(uuid.MustParse("699117b4-7df1-4e14-95e4-5912a6564ef9"))
	user.PayTaxe(taxe_domain.NewTaxe(alreadyPayed))

	inMemoryUsers := infrastructure.NewInMemoryUsers()
	inMemoryUsers.ExpectedUser = *user

	return applications.NewTaxSystem(inMemoryUsers)
}
