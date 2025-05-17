package geometry

import "math"

func circleArea(radius float64) float64 {
	return math.Pi * radius * radius
}

func circleCircumference(radius float64) float64 {
	return 2 * math.Pi * radius
}

func rectangleArea(length, width float64) float64 {
	return length * width
}

func rectanglePerimeter(length, width float64) float64 {
	return 2 * (length + width)
}
