package utils

import "math"

// ConvertPriceToStoredFormat converts a price from a user-friendly format (e.g., 100 for 100.00€)
// to the stored integer format (e.g., 10000).
func ConvertPriceToStoredFormat(price int) int {
	return price * 100
}

// ConvertPriceToDisplayFormat converts a stored price format (e.g., 10000) back to a user-friendly format (e.g., 100 for 100.00€).
func ConvertPriceToDisplayFormat(price int) float64 {
	return math.Round(float64(price)*100) / 10000
}