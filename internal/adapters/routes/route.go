package routes

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/biswasRai/philTest/internal/adapters/repository"
	"github.com/biswasRai/philTest/internal/adapters/handlers"
	"github.com/biswasRai/philTest/internal/core/usecases"

)

func InitializeRoutes(db *pgxpool.Pool, log *logrus.Logger) *mux.Router {
	r := mux.NewRouter()

	reportRepository := repository.NewReportRepository(db)
	reportUsecase := usecases.NewReportsUsecase(reportRepository)
	reportHandler := handlers.NewReportHandler(reportUsecase)

	customerRepository := repository.NewCustomerRepository(db)
	customerUsecase := usecases.NewCustomersUsecase(customerRepository)
	customerHandler := handlers.NewCustomersHandler(customerUsecase)

	ReportsRoutes(r, reportHandler, log)
	CustomersRoutes(r, customerHandler, log)

	return r
}