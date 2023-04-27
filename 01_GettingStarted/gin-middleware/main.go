package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", Ping)
	r.Use(PrintRequest)
	r.Use(PrintResponse)
	// add your handlers here
	r.Run()
}

func PrintRequest(c *gin.Context) {
	fmt.Println("this is request")
}

func PrintResponse(c *gin.Context) {
	now := time.Now()
	//First, execute the middleware and handler under the handler
	c.Next()
	fmt.Printf("this is response, cost: %d",
		time.Since(now).Nanoseconds())
}
