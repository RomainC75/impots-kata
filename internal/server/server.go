package server

import (
	"log/slog"
	"net/http"
)

func Serve(mux *http.ServeMux) {
	server := http.Server{
		Addr:    ":3000",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		slog.Error("server error : ", err.Error())
	}
}
