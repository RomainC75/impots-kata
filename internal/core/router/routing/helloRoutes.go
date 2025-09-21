package routing

import (
	routes "impots/internal/core/router/controllers"
	"net/http"
)

func HelloRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /hello", routes.HelloBasicRoute)
	mux.HandleFunc("GET /helloquery/", routes.HelloRouteQuery)
	mux.HandleFunc("POST /helloparam/{id}", routes.HelloRouteParam)
}
