package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HelloBasicRes struct {
	Name string `json:"name"`
}

func HelloBasicRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Println("--> REQUEST HELLO")
	name := "bob"

	res := HelloBasicRes{
		Name: name,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&res)

}

func HelloRouteParam(w http.ResponseWriter, r *http.Request) {
	param := r.PathValue("id")
	fmt.Println(param)
}

func HelloRouteQuery(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query().Get("message")
	fmt.Println(param)
}
