package main

import (
	"github.com/gin-gonic/gin"
	pathFinder "github.com/reactivejson/flaypath/internal"
	"net/http"
)

/**
 * @author Mohamed-Aly Bou-Hanane
 * © 2023
 */

// GetFlightPath returns the flight path of a person based on the input flights.
func GetFlightPath(c *gin.Context) {
	var flights []pathFinder.Flight
	if err := c.ShouldBindJSON(&flights); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find the starting and ending airports
	path := pathFinder.NewTracker(flights).PathFinder()

	c.JSON(http.StatusOK, gin.H{"flight_path": path})
}

// The core entry point into the app. will setup the config, and run the App
func main() {
	router := gin.Default()
	router.POST("/calculate", GetFlightPath)
	router.Run(":8080")
}
