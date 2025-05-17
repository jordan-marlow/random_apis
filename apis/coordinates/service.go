package coordinates

import "math"

// All angles are expected to come and and be returned in degrees
// Must convert them to radians to use with the math functions.

func cartesianToPolar(x, y, z float64) (float64, float64, float64) {
	r := math.Sqrt(x*x + y*y)
	theta := math.Atan2(y, x) * 180 / math.Pi // Convert radians to degrees
	return r, theta, z
}

func cartesianToSpherical(x, y, z float64) (float64, float64, float64) {
	r := math.Sqrt(x*x + y*y + z*z)
	theta := math.Atan2(y, x) * 180 / math.Pi
	phi := math.Acos(z/r) * 180 / math.Pi
	return r, theta, phi
}

func polarToCartesian(r, theta, z float64) (float64, float64, float64) {
	radians := theta * math.Pi / 180 // Convert degrees to radians
	x := r * math.Cos(radians)
	y := r * math.Sin(radians)
	return x, y, z
}

func polarToSpherical(r, theta, z float64) (float64, float64, float64) {
	phi := math.Acos(z/r) * 180 / math.Pi
	return r, theta, phi
}

func sphericalToCartesian(r, theta, phi float64) (float64, float64, float64) {
	rtheta := theta * math.Pi / 180
	rphi := phi * math.Pi / 180
	x := r * math.Sin(rphi) * math.Cos(rtheta)
	y := r * math.Sin(rphi) * math.Sin(rtheta)
	z := r * math.Cos(rphi)
	return x, y, z
}

func sphericalToPolar(r, theta, phi float64) (float64, float64, float64) {
	rphi := phi * math.Pi / 180
	z := r * math.Cos(rphi)
	return r, theta, z
}
