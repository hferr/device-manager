package main

import (
	"log"
	"net/http"

	"github/hferr/device-manager/internal/protocols/httpjson"
)

func main() {
	s := &http.Server{
		Addr:    ":8080",
		Handler: httpjson.NewRouter(),
	}

	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}
