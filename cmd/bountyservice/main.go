package main

import (
	log "github.com/sirupsen/logrus"
	"bounty-poc/internal/database/postgres"
)

func main() {
	if err := postgres.Initialize(); err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}
}
