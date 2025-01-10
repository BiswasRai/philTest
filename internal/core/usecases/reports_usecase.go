package usecases

import (
	"context"

	"github.com/biswasRai/philTest/internal/adapters/dto"
	"github.com/biswasRai/philTest/internal/adapters/repository"
)

type ReportsUsecase struct {
	reportRepo *repository.ReportRepository
}

func NewReportsUsecase(reportRepo *repository.ReportRepository) *ReportsUsecase {
	return &ReportsUsecase{
		reportRepo: reportRepo,
	}
}

func (uc *ReportsUsecase) GenerateSalesReport(ctx context.Context, reportSalesRequest *dto.ReportsSalesRequest) (map[string]interface{}, error) {
	log.Info("Generating sales report...")
	report, err := uc.reportRepo.GetSalesReport(ctx, reportSalesRequest)
	if err != nil {
		log.WithError(err).Error("Error generating sales report")
		return nil, err
	}

    return report, nil
}

func (uc *ReportsUsecase) GenerateReportsByCustomerSalesReport(ctx context.Context, customerSalesReportsRequest *dto.CustomerSalesReportsRequest) (map[string]interface{}, error) {
	log.Info("Generating customer sales report...")

	report, err := uc.reportRepo.GetCustomersSalesReport(ctx, customerSalesReportsRequest)

	if err != nil {
		log.WithError(err).Error("Error generating customer sales report")
		return nil, err
	}

    return report, nil
}