package coordinates

import "github.com/gin-gonic/gin"

// Define API routes
func RegisterRoutes(r *gin.Engine) {
	r.GET("/cartesian-to-polar", CartesianToPolarHandler)
	r.GET("/cartesian-to-spherical", CartesianToSphericalHandler)
	r.GET("/polar-to-cartesian", PolarToCartesianHandler)
	r.GET("/polar-to-spherical", PolarToSphericalHandler)
	r.GET("/spherical-to-cartesian", SphericalToCartesianHandler)
	r.GET("/spherical-to-polar", SphericalToPolarHandler)
}
