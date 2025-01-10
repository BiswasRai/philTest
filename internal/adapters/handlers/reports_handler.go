package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/biswasRai/philTest/internal/adapters/dto"
	"github.com/biswasRai/philTest/internal/adapters/jsend"
	"github.com/biswasRai/philTest/internal/core/usecases"
)

type ReportsHandler struct {
	reportUsecase *usecases.ReportsUsecase
}

func NewReportHandler(reportUsecase *usecases.ReportsUsecase) *ReportsHandler {
	return &ReportsHandler{
		reportUsecase: reportUsecase,
	}
}


func (h *ReportsHandler) GetReportsBySales(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	 defer func() {
        if err := recover(); err != nil {
            // json.NewEncoder(w).Encode(jsend.Error(err.Error()))
			return
        }
    }()

	 reportSalesRequest, err := dto.ConvertURLValuesToReportsSalesRequest(r.URL.Query())
	 w.Header().Set("Content-Type", "application/json")
    if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(jsend.Error(err.Error()))
        return
    }

	report, err := h.reportUsecase.GenerateSalesReport(ctx, reportSalesRequest)
	
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(jsend.Error(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(jsend.Success(report))
}

func (h *ReportsHandler) GetReportsByCustomers(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	 defer func() {
        if err := recover(); err != nil {
            // json.NewEncoder(w).Encode(jsend.Error(err.Error()))
			return
        }
    }()

	 customerSalesReportRequest, err := dto.ConvertURLValuesToCustomerSalesReportRequest(r.URL.Query())
	 w.Header().Set("Content-Type", "application/json")
    if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(jsend.Error(err.Error()))
        return
    }

	report, err := h.reportUsecase.GenerateReportsByCustomerSalesReport(ctx, customerSalesReportRequest)
	
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(jsend.Error(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(jsend.Success(report))
}