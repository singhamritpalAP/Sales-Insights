package repository

import (
	"log"
	"sales/internal/models"

	"gorm.io/gorm"
)

// GetTopProductsOverall retrieves the top N products overall based on quantity sold within a date range.
func GetTopProductsOverall(db *gorm.DB, n int, startDate string, endDate string) ([]models.Product, error) {
	var topProducts []models.Product
	log.Printf("Executing GetTopProductsOverall: startDate=%s, endDate=%s, limit=%d", startDate, endDate, n)
	// Single query with JOIN to get full product details
	query := db.Model(&models.OrderItem{}).
		Select("products.product_id, products.product_name, products.category, products.unit_price, SUM(order_items.quantity_sold) as quantity_sold").
		Joins("JOIN orders ON order_items.order_id = orders.order_id").
		Joins("JOIN products ON order_items.product_id = products.product_id").
		Where("orders.date_of_sale BETWEEN ? AND ?", startDate, endDate).
		Group("products.product_id, products.product_name, products.category, products.unit_price").
		Order("quantity_sold DESC").
		Limit(n).
		Find(&topProducts)

	if query.Error != nil {
		log.Printf("Query failed: %v", query.Error)
		return nil, query.Error
	}

	return topProducts, nil
}

// GetTopProductsByCategory retrieves the top N products by category based on quantity sold within a date range.
func GetTopProductsByCategory(db *gorm.DB, n int, startDate string, endDate string) (map[string][]models.Product, error) {
	log.Printf("Executing GetTopProductsByCategory: startDate=%s, endDate=%s, limit=%d", startDate, endDate, n)
	var results []models.ProductResult
	query := db.Model(&models.OrderItem{}).
		Select("products.category, products.product_id, products.product_name, products.unit_price, SUM(order_items.quantity_sold) as quantity_sold").
		Joins("JOIN orders ON order_items.order_id = orders.order_id").
		Joins("JOIN products ON order_items.product_id = products.product_id").
		Where("orders.date_of_sale BETWEEN ? AND ?", startDate, endDate).
		Group("products.category, products.product_id, products.product_name, products.unit_price").
		Order("products.category ASC, quantity_sold DESC").
		Find(&results)

	if query.Error != nil {
		log.Printf("Query failed: %v", query.Error)
		return nil, query.Error
	}

	// Group and limit to top N per category
	topProductsByCategory := make(map[string][]models.Product, len(results))
	for _, res := range results {
		if _, ok := topProductsByCategory[res.Category]; !ok {
			topProductsByCategory[res.Category] = make([]models.Product, 0, n)
		}

		// Only append if under the limit of n for this category
		if len(topProductsByCategory[res.Category]) < n {
			topProductsByCategory[res.Category] = append(topProductsByCategory[res.Category], models.Product{
				ProductID:   res.ProductID,
				ProductName: res.ProductName,
				Category:    res.Category,
				UnitPrice:   res.UnitPrice,
			})
		}
	}

	return topProductsByCategory, nil
}

// GetTopProductsByRegion retrieves the top N products by region based on quantity sold within a date range.
func GetTopProductsByRegion(db *gorm.DB, n int, startDate string, endDate string) (map[string][]models.Product, error) {
	//type productResult struct {
	//	Region       string
	//	ProductID    string  `gorm:"column:product_id"`
	//	ProductName  string  `gorm:"column:product_name"`
	//	Category     string  `gorm:"column:category"`
	//	UnitPrice    float64 `gorm:"column:unit_price"`
	//	QuantitySold int     `gorm:"column:quantity_sold"`
	//}
	log.Printf("Executing GetTopProductsByRegion: startDate=%s, endDate=%s, limit=%d", startDate, endDate, n)
	var results []models.ProductResult
	query := db.Model(&models.OrderItem{}).
		Select("orders.region, products.product_id, products.product_name, products.category, products.unit_price, SUM(order_items.quantity_sold) as quantity_sold").
		Joins("JOIN orders ON order_items.order_id = orders.order_id").
		Joins("JOIN products ON order_items.product_id = products.product_id").
		Where("orders.date_of_sale BETWEEN ? AND ?", startDate, endDate).
		Group("orders.region, products.product_id, products.product_name, products.category, products.unit_price").
		Order("orders.region ASC, quantity_sold DESC").
		Find(&results)

	if query.Error != nil {
		log.Printf("Query failed: %v", query.Error)
		return nil, query.Error
	}

	// Group and limit to top N per region
	topProductsByRegion := make(map[string][]models.Product, len(results))
	for _, res := range results {
		if _, ok := topProductsByRegion[res.Region]; !ok {
			topProductsByRegion[res.Region] = make([]models.Product, 0, n)
		}

		// Only append if under the limit of n for this region
		if len(topProductsByRegion[res.Region]) < n {
			topProductsByRegion[res.Region] = append(topProductsByRegion[res.Region], models.Product{
				ProductID:   res.ProductID,
				ProductName: res.ProductName,
				Category:    res.Category,
				UnitPrice:   res.UnitPrice,
			})
		}
	}

	return topProductsByRegion, nil
}
