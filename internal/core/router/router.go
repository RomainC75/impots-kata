package router

import (
	"impots/internal/core/router/routing"
	"net/http"
)

func SetRoutes(mux *http.ServeMux) {
	routing.HelloRoutes(mux)
	routing.TaxeRoutes(mux)
}
