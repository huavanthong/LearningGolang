package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Gin instance
	r := gin.New()

	// Routes
	r.GET("/", HealthCheck)

	// Start server
	if err := r.Run(":3000"); err != nil {
		log.Fatal(err)
	}
}

func HealthCheck(c *gin.Context) {
	res := map[string]interface{}{
		"data": "Server is up and running",
	}

	c.JSON(http.StatusOK, res)
}
