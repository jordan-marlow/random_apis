package coordinates

import (
	"github.com/gin-gonic/gin"
)

func CartesianToPolarHandler(c *gin.Context) {
	x := c.Query("x")
	y := c.Query("y")
	z := c.Query("z")

	xVal, yVal, zVal := convertToFloat(x, y, z)
	r, theta, pz := cartesianToPolar(xVal, yVal, zVal)

	c.JSON(200, gin.H{"r": r, "theta": theta, "z": pz})
}

func CartesianToSphericalHandler(c *gin.Context) {
	x := c.Query("x")
	y := c.Query("y")
	z := c.Query("z")

	xVal, yVal, zVal := convertToFloat(x, y, z)
	r, theta, phi := cartesianToSpherical(xVal, yVal, zVal)

	c.JSON(200, gin.H{"r": r, "theta": theta, "phi": phi})
}

func PolarToCartesianHandler(c *gin.Context) {
	r := c.Query("r")
	theta := c.Query("theta")
	z := c.Query("z")

	rVal, thetaVal, zVal := convertToFloat(r, theta, z)
	cx, cy, cz := polarToCartesian(rVal, thetaVal, zVal)

	c.JSON(200, gin.H{"x": cx, "y": cy, "z": cz})
}

func PolarToSphericalHandler(c *gin.Context) {
	r := c.Query("r")
	theta := c.Query("theta")
	z := c.Query("z")

	rVal, thetaVal, zVal := convertToFloat(r, theta, z)
	sr, stheta, sphi := polarToSpherical(rVal, thetaVal, zVal)

	c.JSON(200, gin.H{"r": sr, "theta": stheta, "phi": sphi})
}

func SphericalToCartesianHandler(c *gin.Context) {
	r := c.Query("r")
	theta := c.Query("theta")
	phi := c.Query("phi")

	rVal, thetaVal, phiVal := convertToFloat(r, theta, phi)
	cx, cy, cz := sphericalToCartesian(rVal, thetaVal, phiVal)

	c.JSON(200, gin.H{"x": cx, "y": cy, "z": cz})
}

func SphericalToPolarHandler(c *gin.Context) {
	r := c.Query("r")
	theta := c.Query("theta")
	phi := c.Query("phi")

	rVal, thetaVal, phiVal := convertToFloat(r, theta, phi)
	pr, ptheta, pz := sphericalToPolar(rVal, thetaVal, phiVal)

	c.JSON(200, gin.H{"r": pr, "theta": ptheta, "z": pz})
}
