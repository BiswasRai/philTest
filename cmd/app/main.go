package main

import (
	"net/http"

	"github.com/biswasRai/philTest/infrastructure/data"
	"github.com/biswasRai/philTest/infrastructure/logger"
	"github.com/biswasRai/philTest/internal/adapters/routes"
)

func main() {
	logger.Initialize()

	log := logger.GetLogger()

	log.Info("Starting the application...")

	db, err := data.Connect()
	if err != nil {
		log.WithError(err).Error("Error connecting to the database")
		return
	}
	defer db.Close()

	r := routes.InitializeRoutes(db, log)
	
	log.Info("Starting the server at 8080...")
	http.ListenAndServe(":8080", r)
}
