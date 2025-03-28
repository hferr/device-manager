package main

import (
	"fmt"
	"log"
	"net/http"

	"github/hferr/device-manager/config"
	"github/hferr/device-manager/internal/protocols/httpjson"
)

func main() {
	c := config.New()

	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", c.Server.Port),
		Handler:      httpjson.NewRouter(),
		ReadTimeout:  c.Server.TimeoutRead,
		WriteTimeout: c.Server.TimeoutWrite,
		IdleTimeout:  c.Server.TimeoutIdle,
	}

	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}
