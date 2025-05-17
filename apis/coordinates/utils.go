package coordinates

import "strconv"

// Convert query parameters to float values
func convertToFloat(a, b, c string) (float64, float64, float64) {
	aVal, _ := strconv.ParseFloat(a, 64)
	bVal, _ := strconv.ParseFloat(b, 64)
	cVal, _ := strconv.ParseFloat(c, 64)
	return aVal, bVal, cVal
}
