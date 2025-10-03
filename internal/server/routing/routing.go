package routing

import "net/http"

func Routing() *http.ServeMux {
	mux := &http.ServeMux{}
	return mux
}
