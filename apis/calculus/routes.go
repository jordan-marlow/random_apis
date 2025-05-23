package calculus

import "github.com/gin-gonic/gin"

// Define API routes
func RegisterRoutes(r *gin.Engine) {
	r.GET("/limit", LimitHandler)
	r.GET("/derivative", DerivativeHandler)
	r.GET("/integral", IntegralHandler)
	r.GET("/plotdata", PlotDataHandler)
	r.GET("/tangentline", TangentLineHandler)
	r.GET("/evaluate2Dfunctionpoint", evaluate2DFunctionPointHandler)
}
