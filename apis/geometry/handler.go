package geometry

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func CircleAreaHandler(c *gin.Context) {
	r := c.Query("r")
	rVal, err := strconv.ParseFloat(r, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid radius"})
	}
	area := circleArea(rVal)
	c.JSON(200, gin.H{"area": area})
}

func CircleCircumferenceHandler(c *gin.Context) {
	r := c.Query("r")
	rVal, err := strconv.ParseFloat(r, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid radius"})
	}
	circumference := circleCircumference(rVal)
	c.JSON(200, gin.H{"circumference": circumference})
}

func RectangleAreaHandler(c *gin.Context) {
	l := c.Query("l")
	w := c.Query("w")
	lVal, err := strconv.ParseFloat(l, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid length"})
	}
	wVal, err := strconv.ParseFloat(w, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid width"})
	}
	area := rectangleArea(lVal, wVal)
	c.JSON(200, gin.H{"area": area})
}

func RectanglePerimeterHandler(c *gin.Context) {
	l := c.Query("l")
	w := c.Query("w")
	lVal, err := strconv.ParseFloat(l, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid length"})
	}
	wVal, err := strconv.ParseFloat(w, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid width"})
	}
	perimeter := rectanglePerimeter(lVal, wVal)
	c.JSON(200, gin.H{"perimeter": perimeter})
}
