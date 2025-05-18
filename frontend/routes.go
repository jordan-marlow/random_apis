package frontend

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		renderTemplate(c, "home.html", gin.H{})
	})
	r.GET("/calculus", func(c *gin.Context) {
		renderTemplate(c, "calculus.html", gin.H{})
	})
	r.GET("/geometry", func(c *gin.Context) {
		renderTemplate(c, "geometry.html", gin.H{})
	})
	r.GET("/coordinates", func(c *gin.Context) {
		renderTemplate(c, "coordinates.html", gin.H{})
	})
}
