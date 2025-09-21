package core

import (
	"fmt"
	"net/http"
)

func Launch(mux *http.ServeMux) {

	server := http.Server{
		Addr:    ":3000",
		Handler: mux,
	}
	fmt.Println("-> launching server on port 3000")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("-> ERROR server ", err.Error())
	}
}
