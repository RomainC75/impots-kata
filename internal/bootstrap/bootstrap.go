package bootstrap

import (
	"impots/internal/core"
	routing "impots/internal/core/router"
	"net/http"
)

func Bootstrap() {
	mux := &http.ServeMux{}
	routing.SetRoutes(mux)
	core.Launch(mux)
}
