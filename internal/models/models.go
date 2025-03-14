package models

import "time"

type Product struct {
	ProductID   string  `gorm:"primaryKey;type:TEXT;column:product_id"`
	ProductName string  `gorm:"type:TEXT"`
	Category    string  `gorm:"type:TEXT"`
	UnitPrice   float64 `gorm:"type:REAL"`
}

type Customer struct {
	CustomerID      string `gorm:"primaryKey;type:TEXT;column:customer_id"`
	CustomerName    string `gorm:"type:TEXT"`
	CustomerEmail   string `gorm:"type:TEXT"`
	CustomerAddress string `gorm:"type:TEXT"`
}

type Order struct {
	OrderID       string    `gorm:"primaryKey;type:TEXT;column:order_id"`
	CustomerID    string    `gorm:"index;type:TEXT;column:customer_id"`
	DateOfSale    time.Time `gorm:"type:TEXT"`
	ShippingCost  float64   `gorm:"type:REAL"`
	PaymentMethod string    `gorm:"type:TEXT"`
	Region        string    `gorm:"type:TEXT;column:region"`
	Customer      Customer  `gorm:"foreignKey:CustomerID;references:CustomerID"`
}

type OrderItem struct {
	OrderItemID  uint    `gorm:"primaryKey;autoIncrement;type:INTEGER;column:order_item_id"`
	OrderID      string  `gorm:"index;type:TEXT;column:order_id"`
	ProductID    string  `gorm:"index;type:TEXT;column:product_id"`
	QuantitySold int     `gorm:"type:INTEGER"`
	Discount     float64 `gorm:"type:REAL"`
	Order        Order   `gorm:"foreignKey:OrderID;references:OrderID"`
	Product      Product `gorm:"foreignKey:ProductID;references:ProductID"`
}

type ProductResult struct {
	Category     string
	ProductID    string  `gorm:"column:product_id"`
	ProductName  string  `gorm:"column:product_name"`
	UnitPrice    float64 `gorm:"column:unit_price"`
	QuantitySold int     `gorm:"column:quantity_sold"`
	Region       string  `gorm:"column:region"`
}

type CustomError struct {
	Prefix  string
	Message string
}

// Error returns the string representation of the custom error.
func (e *CustomError) Error() string {
	return e.Prefix + ": " + e.Message
}
