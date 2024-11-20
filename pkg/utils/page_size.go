package utils

func PageSize(limit int) int {
	if limit <= 0 {
		limit = 20
	} else if limit > 100 {
		limit = 100
	}
	return limit
}