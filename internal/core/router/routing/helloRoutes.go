package routing

import (
	"fmt"
	routes "impots/internal/core/router/controllers"
	"net/http"
)

func HelloRoutes() http.Handler {
	helloMux := http.NewServeMux()
	helloMux.HandleFunc("GET /", routes.HelloBasicRoute)
	helloMux.HandleFunc("GET /query/", routes.HelloRouteQuery)
	helloMux.HandleFunc("POST /param/{id}", routes.HelloRouteParam)

	helloMux.Handle("POST /param-mid/{id}", (&Mid{}).Add(routes.HelloBasicRoute))

	return helloMux
}

type Mid struct {
	fn func(w http.ResponseWriter, r *http.Request)
}

func (m *Mid) Add(fn func(w http.ResponseWriter, r *http.Request)) *Mid {
	m.fn = fn
	return m
}

func (m *Mid) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("--> Go")
	m.fn(w, r)
}
