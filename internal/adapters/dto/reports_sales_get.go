package dto

import (
    "time"
    "net/url"
    "errors"
    "strconv"

	"github.com/biswasRai/philTest/infrastructure/logger"
)

type ReportsSalesRequest struct {
    // Filters
    StartDate       *time.Time `json:"startDate,omitempty"`
    EndDate         *time.Time `json:"endDate,omitempty"`
    ProductCategory *string    `json:"productCategory,omitempty"`
    ProductID       *int       `json:"productId,omitempty"`
    CustomerLocation *string   `json:"customerLocation,omitempty"`

    // Aggregations
    TotalSales       float64 `json:"totalSales"`
    AverageOrderValue float64 `json:"averageOrderValue"`
    NumberOfProductsSold int  `json:"numberOfProductsSold"`
    TotalRevenueByCustomer map[int]float64 `json:"totalRevenueByCustomer"`
    TotalRevenueByProduct  map[int]float64 `json:"totalRevenueByProduct"`
    TotalRevenueByRegion   map[string]float64 `json:"totalRevenueByRegion"`
}

func NewReportsSalesDTO() *ReportsSalesRequest {
    return &ReportsSalesRequest{
        TotalRevenueByCustomer: make(map[int]float64),
        TotalRevenueByProduct:  make(map[int]float64),
        TotalRevenueByRegion:   make(map[string]float64),
    }
}

type CustomerSalesReportsRequest struct {
    // Filters
    StartDate       *time.Time `json:"startDate,omitempty"`
    EndDate         *time.Time `json:"endDate,omitempty"`
}


func init() {
    logger.Initialize()
}

var log = logger.GetLogger()

func ConvertURLValuesToReportsSalesRequest(values url.Values) (*ReportsSalesRequest, error) {
    var reportSalesRequest ReportsSalesRequest

    layout := "2006-01-02"

    if startDate := values.Get("start_date"); startDate != "" {
        parsedStartDate, err := time.Parse(layout, startDate)
        if err != nil {
            log.WithError(err).Error("Error parsing start date")
            return nil, errors.New("invalid start date format, expected YYYY-MM-DD")
        }
        reportSalesRequest.StartDate = &parsedStartDate
    }

    if endDate := values.Get("end_date"); endDate != "" {
        parsedEndDate, err := time.Parse(layout, endDate)
        if err != nil {
            log.WithError(err).Error("Error parsing end date")
            return nil, err
        }
        reportSalesRequest.EndDate = &parsedEndDate
    }

    if productCategory := values.Get("product_category"); productCategory != "" {
        reportSalesRequest.ProductCategory = &productCategory
    }

    if customerLocation := values.Get("customer_location"); customerLocation != "" {
        reportSalesRequest.CustomerLocation = &customerLocation
    }

    if productID := values.Get("product_id"); productID != "" {
        parsedProductID, err := strconv.Atoi(productID)
        if err != nil {
            log.WithError(err).Error("Error parsing product ID")
            return nil, err
        }
        reportSalesRequest.ProductID = &parsedProductID
    }

    return &reportSalesRequest, nil
}

func ConvertURLValuesToCustomerSalesReportRequest(values url.Values) (*CustomerSalesReportsRequest, error) {
    var CustomerSalesReportsRequest CustomerSalesReportsRequest

     layout := "2006-01-02"

    if startDate := values.Get("start_date"); startDate != "" {
        parsedStartDate, err := time.Parse(layout, startDate)
        if err != nil {
            log.WithError(err).Error("Error parsing start date")
            return nil, errors.New("invalid start date format, expected YYYY-MM-DD")
        }
        CustomerSalesReportsRequest.StartDate = &parsedStartDate
    }

    if endDate := values.Get("end_date"); endDate != "" {
        parsedEndDate, err := time.Parse(layout, endDate)
        if err != nil {
            log.WithError(err).Error("Error parsing end date")
            return nil, err
        }
        CustomerSalesReportsRequest.EndDate = &parsedEndDate
    }

    return &CustomerSalesReportsRequest, nil
}