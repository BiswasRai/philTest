package repository

import (
    "context"
    "fmt"
	"github.com/jackc/pgx/v5/pgxpool"
    "database/sql"
    "errors"

	"github.com/biswasRai/philTest/infrastructure/logger"
    "github.com/biswasRai/philTest/internal/adapters/dto"
)

type ReportRepository struct {
    db *pgxpool.Pool
}

func NewReportRepository(db *pgxpool.Pool) *ReportRepository {
    return &ReportRepository{
        db: db,
    }
}

func (r *ReportRepository) GetSalesReport(ctx context.Context, reportSalesRequest *dto.ReportsSalesRequest) (map[string]interface{}, error) {
	
	logger.Initialize()

    log := logger.GetLogger()
 	defer func() {
        if err := recover(); err != nil {
            log.Info("Recovered from panic: %v", err)
        }
    }()
	
    log.WithFields(map[string]interface{}{"reportSalesRequest": reportSalesRequest}).Info("Getting sales report...")

    query := `
    WITH filtered_orders AS (
      SELECT
          o.id AS order_id,
          o.order_date,
          o.customer_id,
          c.location AS customer_location,
          oi.product_id,
          oi.quantity,
          p.category AS product_category,
          oi.price AS total_product_sales
      FROM
          orders o
          JOIN order_items oi ON o.id = oi.order_id
          JOIN customers c ON o.customer_id = c.id
          JOIN products p ON oi.product_id = p.id
      WHERE 1 = 1
      `

      log.WithFields(map[string]interface{}{"productId": reportSalesRequest.ProductID}).Info("Executing query...")

      if reportSalesRequest.StartDate != nil {
		query += fmt.Sprintf(`AND o.order_date >= '%s' `, reportSalesRequest.StartDate.Format("2006-01-02"))
	}
	if reportSalesRequest.EndDate != nil {
		query += fmt.Sprintf(`AND o.order_date <= '%s' `, reportSalesRequest.EndDate.Format("2006-01-02"))
	}
	if reportSalesRequest.ProductCategory != nil {
		query += fmt.Sprintf(`AND p.category = '%s' `, reportSalesRequest.ProductCategory)
	}
	if reportSalesRequest.CustomerLocation != nil {
		query += fmt.Sprintf(`AND c.location = '%s' `, reportSalesRequest.CustomerLocation)
	}
	if reportSalesRequest.ProductID != nil {
		query += fmt.Sprintf(`AND oi.product_id = %d `, *reportSalesRequest.ProductID)
	}

	query +=
		`),
    aggregated_data AS (
        SELECT
            SUM(total_product_sales) AS total_sales,
            AVG(total_product_sales) AS avg_order_value,
            SUM(quantity) AS num_products_sold
        FROM
            filtered_orders
    )
    SELECT * from aggregated_data
  `

    log.WithFields(map[string]interface{}{"query": query}).Info("Executing query...")
	rows, err := r.db.Query(ctx, query)

	if err != nil {
        log.WithError(err).Error("Error executing query")
		return nil, errors.New("Something went wrong")
	}
	defer rows.Close()

	report := make(map[string]interface{})
	for rows.Next() {
		var totalSales sql.NullFloat64
		var avgOrderValue sql.NullFloat64
		var numProductsSold sql.NullInt64

		err = rows.Scan(&totalSales, &avgOrderValue, &numProductsSold)

		if err != nil {
			return nil, err
		}

		if totalSales.Valid {
			report["totalSales"] = totalSales.Float64
		} else {
			report["totalSales"] = 0
		}

		if avgOrderValue.Valid {
			report["averageOrderValue"] = avgOrderValue.Float64
		} else {
			report["averageOrderValue"] = 0
		}

		if numProductsSold.Valid {
			report["numberOfProductsSold"] = numProductsSold.Int64
		} else {
			report["numberOfProductsSold"] = 0
		}

	}
	return report, nil
}

func (r *ReportRepository) GetCustomersSalesReport(ctx context.Context, reportCustomerSalesRequest *dto.CustomerSalesReportsRequest) (map[string]interface{}, error) {
	
	logger.Initialize()

    log := logger.GetLogger()
 	defer func() {
        if err := recover(); err != nil {
            log.Info("Recovered from panic: %v", err)
        }
    }()
	
    log.WithFields(map[string]interface{}{"reportCustomerSalesRequest": reportCustomerSalesRequest}).Info("Getting customer sales report...")

	query := `
    WITH filtered_customers AS (
        SELECT 
            c.id AS customer_id,
            c.signup_date,
            c.lifetime_value,
            COUNT(o.id) AS total_orders
        FROM customers c
        LEFT JOIN orders o ON c.id = o.customer_id
        WHERE 1 = 1
    `

	if reportCustomerSalesRequest.StartDate != nil {
		query += fmt.Sprintf(`AND c.signup_date >= '%s' `, reportCustomerSalesRequest.StartDate.Format("2006-01-02"))
	}

	if reportCustomerSalesRequest.EndDate != nil {
		query += fmt.Sprintf(`AND c.signup_date <= '%s' `, reportCustomerSalesRequest.EndDate.Format("2006-01-02"))
	}

	query += `
      GROUP BY c.id
  ),
  -- Aggregations
  aggregated_data as (
      SELECT 
      COUNT(customer_id) AS total_customers,
      AVG(total_orders) AS avg_order_frequency
      from filtered_customers
  )
  select * from aggregated_data
  `

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	report := make(map[string]interface{})
	for rows.Next() {
		var totalCustomers sql.NullInt64
		var avgOrderFrequency sql.NullFloat64

		err = rows.Scan(&totalCustomers, &avgOrderFrequency)
		if err != nil {
			return nil, err
		}

		if totalCustomers.Valid {
			report["totalCustomers"] = totalCustomers.Int64
		} else {
			report["totalCustomers"] = 0
		}

		if avgOrderFrequency.Valid {
			report["avgOrderFrequency"] = avgOrderFrequency.Float64
		} else {
			report["avgOrderFrequency"] = nil
		}
	}

	return report, nil
}
