package routes

import (
    "context"
    "net/http"

    "github.com/biswasRai/philTest/internal/adapters/handlers"
    "github.com/gorilla/mux"
    "github.com/sirupsen/logrus"
)

func CustomersRoutes(r *mux.Router, customersHandler *handlers.CustomersHandler, log *logrus.Logger) {
    r.HandleFunc("/customers", func(w http.ResponseWriter, r *http.Request) {
        customersHandler.CreateCustomer(context.Background(), w, r)
    }).Methods("POST")

	r.HandleFunc("/customers", func(w http.ResponseWriter, r *http.Request) {
        customersHandler.GetCustomers(context.Background(), w, r)
    }).Methods("GET")

	r.HandleFunc("/customers/{id}", func(w http.ResponseWriter, r *http.Request) {
        customersHandler.UpdateCustomer(context.Background(), w, r)
    }).Methods("PUT")

	r.HandleFunc("/customers/{id}", func(w http.ResponseWriter, r *http.Request) {
        customersHandler.DeleteCustomer(context.Background(), w, r)
    }).Methods("DELETE")
}