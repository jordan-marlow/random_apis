package coordinates

import (
	"math"
	"testing"
)

func TestCartesianToPolar(t *testing.T) {
	tests := []struct {
		x, y, z float64
		r       float64
		theta   float64
		pz      float64
	}{
		{1, 1, 0, math.Sqrt(2), 45, 0},
		{4, 0, 0, 4, 0, 0},
		{0, 4, 0, 4, 90, 0},
	}

	for _, test := range tests {
		r, theta, pz := cartesianToPolar(test.x, test.y, test.z)
		if r != test.r || theta != test.theta {
			t.Errorf("cartesianToPolar(%f, %f, %f) = (%f, %f, %f); want (%f, %f, %f)", test.x, test.y, test.z, r, theta, pz, test.r, test.theta, test.pz)
		}
	}
}
