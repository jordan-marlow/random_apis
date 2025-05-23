package coordinates

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func CartesianToPolarHandler(c *gin.Context) {
	x := c.Query("x")
	y := c.Query("y")
	z := c.Query("z")

	xVal, err := strconv.ParseFloat(x, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid x value"})
		return
	}
	yVal, err := strconv.ParseFloat(y, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid y value"})
		return
	}
	zVal, err := strconv.ParseFloat(z, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid z value"})
		return
	}
	r, theta, pz := cartesianToPolar(xVal, yVal, zVal)

	c.JSON(200, gin.H{"r": r, "theta": theta, "z": pz})
}

func CartesianToSphericalHandler(c *gin.Context) {
	x := c.Query("x")
	y := c.Query("y")
	z := c.Query("z")

	xVal, err := strconv.ParseFloat(x, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid x value"})
		return
	}
	yVal, err := strconv.ParseFloat(y, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid y value"})
		return
	}
	zVal, err := strconv.ParseFloat(z, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid z value"})
		return
	}
	r, theta, phi := cartesianToSpherical(xVal, yVal, zVal)

	c.JSON(200, gin.H{"r": r, "theta": theta, "phi": phi})
}

func PolarToCartesianHandler(c *gin.Context) {
	r := c.Query("r")
	theta := c.Query("theta")
	z := c.Query("z")

	rVal, err := strconv.ParseFloat(r, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid r value"})
		return
	}
	thetaVal, err := strconv.ParseFloat(theta, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid theta value"})
		return
	}
	zVal, err := strconv.ParseFloat(z, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid z value"})
		return
	}
	cx, cy, cz := polarToCartesian(rVal, thetaVal, zVal)

	c.JSON(200, gin.H{"x": cx, "y": cy, "z": cz})
}

func PolarToSphericalHandler(c *gin.Context) {
	r := c.Query("r")
	theta := c.Query("theta")
	z := c.Query("z")

	rVal, err := strconv.ParseFloat(r, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid r value"})
		return
	}
	thetaVal, err := strconv.ParseFloat(theta, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid theta value"})
		return
	}
	zVal, err := strconv.ParseFloat(z, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid z value"})
		return
	}
	sr, stheta, sphi := polarToSpherical(rVal, thetaVal, zVal)

	c.JSON(200, gin.H{"r": sr, "theta": stheta, "phi": sphi})
}

func SphericalToCartesianHandler(c *gin.Context) {
	r := c.Query("r")
	theta := c.Query("theta")
	phi := c.Query("phi")

	rVal, err := strconv.ParseFloat(r, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid r value"})
		return
	}
	thetaVal, err := strconv.ParseFloat(theta, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid theta value"})
		return
	}
	phiVal, err := strconv.ParseFloat(phi, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid phi value"})
		return
	}
	cx, cy, cz := sphericalToCartesian(rVal, thetaVal, phiVal)

	c.JSON(200, gin.H{"x": cx, "y": cy, "z": cz})
}

func SphericalToPolarHandler(c *gin.Context) {
	r := c.Query("r")
	theta := c.Query("theta")
	phi := c.Query("phi")

	rVal, err := strconv.ParseFloat(r, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid r value"})
		return
	}
	thetaVal, err := strconv.ParseFloat(theta, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid theta value"})
		return
	}
	phiVal, err := strconv.ParseFloat(phi, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid phi value"})
		return
	}
	pr, ptheta, pz := sphericalToPolar(rVal, thetaVal, phiVal)

	c.JSON(200, gin.H{"r": pr, "theta": ptheta, "z": pz})
}
