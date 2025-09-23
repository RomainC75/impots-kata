package router

import (
	"impots/internal/core/router/routing"
	"net/http"
)

func SetRoutes(mux *http.ServeMux) {
	mux.Handle("/hello", routing.HelloRoutes())
	mux.Handle("/taxes", routing.TaxeRoutes())
}
