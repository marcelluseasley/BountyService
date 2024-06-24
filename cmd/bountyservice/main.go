package main

import (
	"net/http"

	"bounty-poc/internal/bountyservice/rest"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"bounty-poc/internal/database/postgres"
	log "github.com/sirupsen/logrus"
)

func main() {
	if err := postgres.Initialize(); err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", rest.Ping)
	r.Post("/bounty", rest.CreateBounty)

	http.ListenAndServe(":3000", r)
}
