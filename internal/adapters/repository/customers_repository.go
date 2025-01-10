package repository

import (
	"context"
	"fmt"
	"time"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/biswasRai/philTest/infrastructure/logger"
	"github.com/biswasRai/philTest/internal/adapters/dto"
	"github.com/biswasRai/philTest/internal/core/entities"
)

type CustomerRepository struct {
	db *pgxpool.Pool
}

func NewCustomerRepository(db *pgxpool.Pool) *CustomerRepository {
	return &CustomerRepository{
		db: db,
	}
}

func init() {
	logger.Initialize()
}

var log = logger.GetLogger()

func (r *CustomerRepository) GetCustomers(ctx context.Context) ([]entities.Customer, error) {
    defer func() {
        if err := recover(); err != nil {
            log.Printf("Recovered from panic: %v", err)
        }
    }()

    query := `
    SELECT
        id,
        name,
        email,
        location,
        lifetime_value,
        signup_date
    FROM
        customers
    `

    rows, err := r.db.Query(ctx, query)
    if err != nil {
        log.Printf("Error executing query: %v", err)
        return nil, err
    }
    defer rows.Close()

    var customers []entities.Customer
    for rows.Next() {
        var customer entities.Customer
        err := rows.Scan(&customer.ID, &customer.Name, &customer.Email, &customer.Location, &customer.LifetimeValue, &customer.SignupDate)
        if err != nil {
            log.Printf("Error scanning customer row: %v", err)
            return nil, err
        }
        customers = append(customers, customer)
    }

    return customers, nil
}

func (r *CustomerRepository) CreateCustomer(ctx context.Context, createCustomerRequest *dto.CreateCustomerRequest) (int, error) {
	defer func() {
        if err := recover(); err != nil {
            log.Info("Recovered from panic: %v", err)
        }
    }()

	log.WithFields(map[string]interface{}{"createCustomerRequest": createCustomerRequest}).Info("Creating customer...")

	query := `INSERT INTO customers (name, email, signup_date, location, lifetime_value)
			  VALUES ($1, $2, $3, $4, $5) RETURNING id`
	var id int
	 err := r.db.QueryRow(ctx, query, createCustomerRequest.Name, createCustomerRequest.Email, time.Now(), createCustomerRequest.Location, createCustomerRequest.LifetimeValue).Scan(&id)

	if err != nil {
		log.WithError(err).Error("Error creating customer")
		return 0, err
	}

	return id, nil
}

func (r *CustomerRepository) UpdateCustomerById(ctx context.Context, id int, updateCustomerRequest *dto.UpdateCustomerRequest) (map[string]interface{}, error) {
	defer func() {
        if err := recover(); err != nil {
            log.Info("Recovered from panic: %v", err)
        }
    }()

	log.WithFields(map[string]interface{}{"updateCustomerRequest": updateCustomerRequest}).Info("Updating customer...")

	query := `UPDATE customers SET name=$1, location=$2 WHERE id=$3`
	_, err := r.db.Exec(ctx, query, updateCustomerRequest.Name, updateCustomerRequest.Location, id)

	if err != nil {
		log.WithError(err).Error("Error updating customers")
		return nil, err
	}

	updatedCustomer := map[string]interface{}{
        "id":             id,
        "name":           updateCustomerRequest.Name,
        "location":       updateCustomerRequest.Location,
    }

	return updatedCustomer, nil
}

func (r *CustomerRepository) DeleteCustomerById(ctx context.Context, id int) (int, error) {
defer func() {
        if err := recover(); err != nil {
            log.Info("Recovered from panic: %v", err)
        }
    }()

	log.Info("Deleting customer...")

	query := fmt.Sprintf("DELETE FROM customers WHERE id=%d", id)
	_, err := r.db.Exec(ctx, query)

	if err != nil {
		log.WithError(err).Error("Error deleting customer")
		return 0, err
	}

	return id, nil
}