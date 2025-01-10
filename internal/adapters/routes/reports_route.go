package routes

import (
    "context"
    "net/http"

    "github.com/biswasRai/philTest/internal/adapters/handlers"
    "github.com/gorilla/mux"
    "github.com/sirupsen/logrus"
)

func ReportsRoutes(r *mux.Router, reportsHandler *handlers.ReportsHandler, log *logrus.Logger) {
    r.HandleFunc("/reports/sales", func(w http.ResponseWriter, r *http.Request) {
        reportsHandler.GetReportsBySales(context.Background(), w, r)
    }).Methods("GET")

    r.HandleFunc("/reports/customers", func(w http.ResponseWriter, r *http.Request) {
        reportsHandler.GetReportsByCustomers(context.Background(), w, r)
    }).Methods("GET")
}