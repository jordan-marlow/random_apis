package main

import (
	"randomapis/apis/coordinates"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	coordinates.RegisterRoutes(r) // Register Coordinate API routes

	r.Run(":8080") // Start server
}
