package calculus

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func LimitHandler(c *gin.Context) {
	// Extract parameters from the request
	expression := c.Query("expression")
	point := c.Query("point")

	// Validate input
	if expression == "" || point == "" {
		c.JSON(400, gin.H{"error": "Missing required parameters"})
		return
	}
	pVal, _ := strconv.ParseFloat(point, 64)

	// Calculate the limit using a library or custom logic
	f, err := parseFunction(expression)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid expression" + err.Error()})
		return
	}
	result := limit(f, pVal)

	// Return the result as JSON
	c.JSON(200, gin.H{"limit": result})
}

func DerivativeHandler(c *gin.Context) {
	// Extract parameters from the request
	expression := c.Query("expression")
	point := c.Query("point")

	// Validate input
	if expression == "" || point == "" {
		c.JSON(400, gin.H{"error": "Missing required parameters"})
		return
	}
	pVal, _ := strconv.ParseFloat(point, 64)

	// Calculate the derivative using a library or custom logic
	f, err := parseFunction(expression)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid expression" + err.Error()})
		return
	}
	result := derivative(f, pVal)

	// Return the result as JSON
	c.JSON(200, gin.H{"derivative": result})
}
func IntegralHandler(c *gin.Context) {
	// Extract parameters from the request
	expression := c.Query("expression")
	lowerBound := c.Query("lower_bound")
	upperBound := c.Query("upper_bound")
	steps := c.Query("steps")

	// Validate input
	if expression == "" || lowerBound == "" || upperBound == "" || steps == "" {
		c.JSON(400, gin.H{"error": "Missing required parameters"})
		return
	}
	lbVal, _ := strconv.ParseFloat(lowerBound, 64)
	ubVal, _ := strconv.ParseFloat(upperBound, 64)
	stepsVal, _ := strconv.Atoi(steps)
	if stepsVal <= 0 {
		c.JSON(400, gin.H{"error": "Steps must be a positive integer"})
		return
	}
	if stepsVal > 100_000 {
		c.JSON(400, gin.H{"error": "Steps must be less than or equal to 100,000"})
		return
	}

	// Calculate the integral using a library or custom logic
	f, err := parseFunction(expression)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid expression" + err.Error()})
		return
	}
	result := integral(f, lbVal, ubVal, stepsVal)

	// Return the result as JSON
	c.JSON(200, gin.H{"integral": result})
}
