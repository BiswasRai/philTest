package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	
	"github.com/gorilla/mux"
	"github.com/biswasRai/philTest/internal/adapters/dto"
	"github.com/biswasRai/philTest/internal/adapters/jsend"
	"github.com/biswasRai/philTest/internal/core/usecases"
)

type CustomersHandler struct {
	customerUsecase *usecases.CustomersUsecase
}

func NewCustomersHandler(customerUsecase *usecases.CustomersUsecase) *CustomersHandler {
	return &CustomersHandler{
		customerUsecase: customerUsecase,
	}
}

func (h *CustomersHandler) GetCustomers(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	 w.Header().Set("Content-Type", "application/json")

	report, err := h.customerUsecase.GetCustomers(ctx)
	
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(jsend.Error(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(jsend.Success(report))
}

func (h *CustomersHandler) CreateCustomer(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	 createCustomerRequest, err := dto.ValidatePayloadCreateCustomerRequest(r)
	 w.Header().Set("Content-Type", "application/json")

    if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(jsend.Error(err.Error()))
        return
    }

	report, err := h.customerUsecase.CreateCustomer(ctx, createCustomerRequest)
	
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(jsend.Error(err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(jsend.Success(report))
}

func (h *CustomersHandler) UpdateCustomer(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr :=vars["id"]
	updateCustomerRequest, err := dto.ValidatePayloadUpdateCustomerRequest(r)
	w.Header().Set("Content-Type", "application/json")
    if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(jsend.Error(err.Error()))
        return
    }

	id, err := strconv.Atoi(idStr)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(jsend.Error(err.Error()))
		return
	}

	report, err := h.customerUsecase.UpdateCustomerById(ctx, id, updateCustomerRequest)
	
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(jsend.Error(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(jsend.Success(report))
}

func (h *CustomersHandler) DeleteCustomer(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr :=vars["id"]
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(jsend.Error(err.Error()))
		return
	}

	report, err := h.customerUsecase.DeleteCustomerById(ctx, id)
	
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(jsend.Error(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(jsend.Success(report))
}
