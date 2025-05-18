package geometry

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func CircleAreaHandler(c *gin.Context) {
	r := c.Query("radius")
	rVal, err := strconv.ParseFloat(r, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid radius"})
		return
	}
	area := circleArea(rVal)
	c.JSON(200, gin.H{"area": area})
}

func CircleCircumferenceHandler(c *gin.Context) {
	r := c.Query("radius")
	rVal, err := strconv.ParseFloat(r, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid radius"})
		return
	}
	circumference := circleCircumference(rVal)
	c.JSON(200, gin.H{"circumference": circumference})
}

func RectangleAreaHandler(c *gin.Context) {
	l := c.Query("length")
	w := c.Query("width")
	lVal, err := strconv.ParseFloat(l, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid length"})
		return
	}
	wVal, err := strconv.ParseFloat(w, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid width"})
		return
	}
	area := rectangleArea(lVal, wVal)
	c.JSON(200, gin.H{"area": area})
}

func RectanglePerimeterHandler(c *gin.Context) {
	l := c.Query("length")
	w := c.Query("width")
	lVal, err := strconv.ParseFloat(l, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid length"})
		return
	}
	wVal, err := strconv.ParseFloat(w, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid width"})
		return
	}
	perimeter := rectanglePerimeter(lVal, wVal)
	c.JSON(200, gin.H{"perimeter": perimeter})
}
