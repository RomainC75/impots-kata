package bootstrap

import (
<<<<<<< HEAD
	"impots/internal/core"
	routing "impots/internal/core/router"
	"net/http"
)

func Bootstrap() {
	mux := &http.ServeMux{}
	routing.SetRoutes(mux)
	core.Launch(mux)
=======
	"impots/internal/server"
	"impots/internal/server/routing"
)

func Bootstrap() {
	mux := routing.Routing()

	server.Serve(mux)
>>>>>>> re
}
