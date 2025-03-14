package services

import (
	"encoding/csv"
	"gorm.io/gorm"
	"log"
	"os"
	"sales/internal/constants"
	"sales/internal/models"
	"sales/internal/utils"
	"strings"
)

// loadCSVData reads the CSV file and returns the data as a slice of string slices.
func loadCSVData(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Reading all rows at once
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}

func validateAndTransformData(records [][]string) ([]models.Product, []models.Customer, []models.Order, []models.OrderItem, error) {
	var products []models.Product
	var customers []models.Customer
	var orders []models.Order
	var orderItems []models.OrderItem

	// Processing rows; skipping the header
	for i, record := range records {
		if i == 0 {
			continue // Skip header row
		}

		// record having less than 15 values is incomplete
		if len(record) < 15 {
			log.Printf("Skipping record %d due to insufficient fields: %v\n", i+1, record)
			continue
		}

		// Validate and transform data
		productID := strings.TrimSpace(record[1])
		customerID := strings.TrimSpace(record[2])

		// Parse and validate numeric fields
		unitPrice, err := utils.ParsePrice(record[8])
		if err != nil {
			log.Printf("Skipping record %d due to invalid Unit Price: %v\n", i+1, err)
			continue
		}

		discount, err := utils.ParseDiscount(record[9])
		if err != nil {
			log.Printf("Skipping record %d due to invalid Discount: %v\n", i+1, err)
			continue
		}

		shippingCost, err := utils.ParsePrice(record[10])
		if err != nil {
			log.Printf("Skipping record %d due to invalid Shipping Cost: %v\n", i+1, err)
			continue
		}

		quantitySold, err := utils.ParseInt(record[7])
		if err != nil {
			log.Printf("Skipping record %d due to invalid Quantity Sold: %v\n", i+1, err)
			continue
		}

		dateOfSale, err := utils.ParseDate(record[6])
		if err != nil {
			log.Printf("Skipping record %d due to invalid Date of Sale: %v\n", i+1, err)
			continue
		}

		// Creating models to feed in database
		product := models.Product{
			ProductID:   productID,
			ProductName: strings.TrimSpace(record[3]),
			Category:    strings.TrimSpace(record[4]),
			UnitPrice:   unitPrice,
		}

		customer := models.Customer{
			CustomerID:      customerID,
			CustomerName:    strings.TrimSpace(record[12]),
			CustomerEmail:   strings.TrimSpace(record[13]),
			CustomerAddress: strings.TrimSpace(record[14]),
		}

		order := models.Order{
			OrderID:       strings.TrimSpace(record[0]),
			CustomerID:    customerID,
			DateOfSale:    dateOfSale,
			ShippingCost:  shippingCost,
			PaymentMethod: strings.TrimSpace(record[11]),
			Customer:      customer,
			Region:        strings.TrimSpace(record[5]),
		}

		orderItem := models.OrderItem{
			OrderID:      strings.TrimSpace(record[0]),
			ProductID:    productID,
			QuantitySold: quantitySold,
			Discount:     discount,
			Order:        order,
			Product:      product,
		}

		// Append to slices
		products = append(products, product)
		customers = append(customers, customer)
		orders = append(orders, order)
		orderItems = append(orderItems, orderItem)
	}

	return products, customers, orders, orderItems, nil
}

// RefreshDatabase refreshes the database with data from the CSV file. todo implement batching
func RefreshDatabase(db *gorm.DB) error {
	filePath := constants.CSVFilePath

	// Load CSV data
	records, err := loadCSVData(filePath)
	if err != nil {
		return err
	}

	// Validate and transform data
	products, customers, orders, orderItems, err := validateAndTransformData(records)
	if err != nil {
		return err
	}

	// Begin transaction
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	// Create or update customers
	for _, customer := range customers {
		if err := createOrUpdateCustomer(tx, customer); err != nil {
			tx.Rollback()
			return err
		}
	}

	// Create or update products
	for _, product := range products {
		if err := createOrUpdateProduct(tx, product); err != nil {
			tx.Rollback()
			return err
		}
	}

	// Create orders and order items
	for _, order := range orders {
		if err := createOrder(tx, order); err != nil {
			tx.Rollback()
			return err
		}

		// Find associated order item
		for _, orderItem := range orderItems {
			if orderItem.OrderID == order.OrderID {
				if err := createOrderItem(tx, orderItem); err != nil {
					tx.Rollback()
					return err
				}
			}
		}
	}

	// Commit transaction
	return tx.Commit().Error
}

// createOrUpdateCustomer creates or updates a customer in the database.
func createOrUpdateCustomer(db *gorm.DB, customer models.Customer) error {
	result := db.FirstOrCreate(&customer, models.Customer{CustomerID: customer.CustomerID})
	return result.Error
}

// createOrUpdateProduct creates or updates a product in the database.
func createOrUpdateProduct(db *gorm.DB, product models.Product) error {
	result := db.FirstOrCreate(&product, models.Product{ProductID: product.ProductID})
	return result.Error
}

// createOrder creates an order in the database.
func createOrder(db *gorm.DB, order models.Order) error {
	result := db.Create(&order)
	return result.Error
}

// createOrderItem creates an order item in the database.
func createOrderItem(db *gorm.DB, orderItem models.OrderItem) error {
	result := db.Create(&orderItem)
	return result.Error
}
