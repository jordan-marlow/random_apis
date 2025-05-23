package streams

import "github.com/gin-gonic/gin"

// Define API routes
func RegisterRoutes(r *gin.Engine) {
	r.GET("/tempstream", StreamTempDataHandler)
}
