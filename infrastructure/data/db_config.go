package data

import (
	"context"
	"log"

	"github.com/biswasRai/philTest/infrastructure/logger"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Config() *pgxpool.Config {
	DATABASE_URL := "postgres://postgres:postgres@localhost:5432/phil_test" // TODO handle through env variables

	dbConfig, err := pgxpool.ParseConfig(DATABASE_URL)

	if err != nil {
		log.Fatal("Failed to parse database URL:", err)  // Handle with logger
	}

	return dbConfig
}

func Connect() (*pgxpool.Pool, error) {
	logger.Initialize()

	log := logger.GetLogger()

	conn, err := pgxpool.NewWithConfig(context.Background(), Config())

	if err != nil {
		log.WithError(err).Error("Error connecting to the database")

		return nil, err
	}

	connection, err := conn.Acquire(context.Background())
	if err != nil {
		log.WithError(err).Error("Error connecting to the database")
		return nil, err
	}
	defer connection.Release()


	log.WithFields(map[string]interface{}{
		"conn": conn,
		"env":  "dev",
	}).Info("Successfully connected to the database!")

	return conn, nil
}
