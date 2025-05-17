package geometry

import "github.com/gin-gonic/gin"

// Define API routes
func RegisterRoutes(r *gin.Engine) {
	r.GET("/area-circle", CircleAreaHandler)
	r.GET("/circumference-circle", CircleCircumferenceHandler)
	r.GET("/area-rectangle", RectangleAreaHandler)
	r.GET("/perimeter-rectangle", RectanglePerimeterHandler)
}
