package main

import (
	"os"
	"randomapis/apis/calculus"
	"randomapis/apis/coordinates"
	"randomapis/apis/geometry"
	"randomapis/frontend"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("/static", "./static")

	frontend.RegisterRoutes(r)    // Register frontend routes
	coordinates.RegisterRoutes(r) // Register Coordinate API routes
	geometry.RegisterRoutes(r)    // Register Geometry API routes
	calculus.RegisterRoutes(r)    // Register Calculus API routes

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default fallback for local development
	}

	r.Run(":" + port) // Run on the specified port
}
