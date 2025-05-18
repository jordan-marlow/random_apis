package geometry

import (
	"math"
	"randomapis/utilities"
)

func circleArea(radius float64) float64 {
	return utilities.RoundTo5Decimals(math.Pi * radius * radius)
}

func circleCircumference(radius float64) float64 {
	return utilities.RoundTo5Decimals(2 * math.Pi * radius)
}

func rectangleArea(length, width float64) float64 {
	return utilities.RoundTo5Decimals(length * width)
}

func rectanglePerimeter(length, width float64) float64 {
	return utilities.RoundTo5Decimals(2 * (length + width))
}
