package main

import (
	"github.com/gin-gonic/gin"
	pathFinder "github.com/reactivejson/flaypath/internal"
	"net/http"
)

// GetFlightPath returns the flight path of a person based on the input flights.
func GetFlightPath(c *gin.Context) {
	var flights []pathFinder.Flight
	if err := c.ShouldBindJSON(&flights); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find the starting and ending airports
	tracked := pathFinder.New(flights).PathFinder()

	c.JSON(http.StatusOK, gin.H{"flight_path": tracked})
}

func main() {
	router := gin.Default()
	router.POST("/calculate", GetFlightPath)
	router.Run(":8080")
}
