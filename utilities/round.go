package utilities

import "math"

func RoundTo5Decimals(x float64) float64 {
	return math.Round(x*1e5) / 1e5
}
