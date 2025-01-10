package usecases

import (
	"context"

	"github.com/biswasRai/philTest/internal/adapters/dto"
	"github.com/biswasRai/philTest/internal/adapters/repository"
	"github.com/biswasRai/philTest/infrastructure/logger"
	"github.com/biswasRai/philTest/internal/core/entities"

)

type CustomersUsecase struct {
	customerRepo *repository.CustomerRepository
}

func NewCustomersUsecase(customerRepo *repository.CustomerRepository) *CustomersUsecase {
	return &CustomersUsecase{
		customerRepo: customerRepo,
	}
}

func init() {
	logger.Initialize()
}

var log = logger.GetLogger()

func (uc *CustomersUsecase) GetCustomers(ctx context.Context) ([]entities.Customer, error) {
	log.Info("Get customers...")
	report, err := uc.customerRepo.GetCustomers(ctx)
	if err != nil {
		log.WithError(err).Error("Error getting customers")
		return nil, err
	}

    return report, nil
}

func (uc *CustomersUsecase) CreateCustomer(ctx context.Context, CreateCustomerRequest *dto.CreateCustomerRequest) (int, error) {
	log.Info("Creating customer...")
	report, err := uc.customerRepo.CreateCustomer(ctx, CreateCustomerRequest)
	if err != nil {
		log.WithError(err).Error("Error creating customer")
		return 0, err
	}

	log.Info("Customer created successfully")

    return report, nil
}

func (uc *CustomersUsecase) UpdateCustomerById(ctx context.Context, id int, updateCustomerRequest *dto.UpdateCustomerRequest) (map[string]interface{}, error) {
	log.Info("Updating customer...")

	report, err := uc.customerRepo.UpdateCustomerById(ctx, id, updateCustomerRequest)

	if err != nil {
		log.WithError(err).Error("Error updating customer")
		return nil, err
	}

    return report, nil
}

func (uc *CustomersUsecase) DeleteCustomerById(ctx context.Context, id int) (int, error) {
	log.Info("Deleting customer...")
	report, err := uc.customerRepo.DeleteCustomerById(ctx, id)
	if err != nil {
		log.WithError(err).Error("Error deleting customer")
		return 0, err
	}

    return report, nil
}