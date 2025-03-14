package utils

import (
	"fmt"
	"sales/internal/constants"
	"sales/internal/models"
	"strconv"
	"strings"
	"time"
)

// logError creates a new error with the given message.
func logError(message string) error {
	return logErrorWithCustomMessage(message)
}

// logErrorWithCustomMessage creates a new error with a custom message.
func logErrorWithCustomMessage(message string) error {
	return logErrorWithPrefix("Validation Error", message)
}

// logErrorWithPrefix creates a new error with a custom prefix.
func logErrorWithPrefix(prefix string, message string) error {
	return &models.CustomError{
		Prefix:  prefix,
		Message: message,
	}
}

// ParsePrice parses and validates a price string.
func ParsePrice(priceStr string) (float64, error) {
	priceStr = strings.TrimSpace(priceStr)
	if priceStr == "" {
		return 0, logError("price cannot be empty")
	}
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		return 0, logError("invalid price format")
	}
	if price < 0 {
		return 0, logError("price cannot be negative")
	}
	return price, nil
}

// ParseInt parses and validates an integer string.
func ParseInt(intStr string) (int, error) {
	intStr = strings.TrimSpace(intStr)
	if intStr == "" {
		return 0, logError("integer cannot be empty")
	}
	quantity, err := strconv.Atoi(intStr)
	if err != nil {
		return 0, logError("invalid integer format")
	}
	if quantity <= 0 {
		return 0, logError("quantity must be greater than zero")
	}
	return quantity, nil
}

// ParseDiscount parses and validates a discount string.
func ParseDiscount(discountStr string) (float64, error) {
	discountStr = strings.TrimSpace(discountStr)
	if discountStr == "" {
		return 0, logError("discount cannot be empty")
	}
	discount, err := strconv.ParseFloat(discountStr, 64)
	if err != nil {
		return 0, logError("invalid discount format")
	}
	if discount < 0 || discount > 1 {
		return 0, logError("discount must be between 0 and 1")
	}
	return discount, nil
}

// ParseDate parses and validates a date string.
func ParseDate(dateStr string) (time.Time, error) {
	dateStr = strings.TrimSpace(dateStr)
	if dateStr == "" {
		return time.Time{}, logError("date cannot be empty")
	}
	date, err := time.Parse(constants.DateFormat, dateStr)
	if err != nil {
		return time.Time{}, logError("invalid date format")
	}
	return date, nil
}

// ValidateDateFormat checks if a date string is in YYYY-MM-DD format.
func ValidateDateFormat(dateStr string) error {
	_, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return fmt.Errorf("date must be in YYYY-MM-DD format: %v", err)
	}
	return nil
}

// ValidateParamsAndGetLimit validates params received in request
func ValidateParamsAndGetLimit(nStr, startDate, endDate string) (int, error) {
	// limit validation
	n, err := strconv.Atoi(nStr)
	if err != nil || n <= 0 {
		return 0, constants.ErrInvalidLimit
	}

	// start date validation
	if err := ValidateDateFormat(startDate); err != nil {
		return 0, constants.ErrInvalidStartDate
	}

	// end date validation
	if err := ValidateDateFormat(endDate); err != nil {
		return 0, constants.ErrInvalidEndDate
	}

	return n, nil
}
