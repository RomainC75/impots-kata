package bootstrap

import (
	"impots/internal/server"
	"impots/internal/server/routing"
)

func Bootstrap() {
	mux := routing.Routing()

	server.Serve(mux)
}
