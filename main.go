package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Serve frontend
	r.Static("/static", "./static")

	// API route
	r.GET("/summarize", StreamSummaryHandler)

	r.Run(":8080")
}

func sendSSE(c *gin.Context, message string) {
	fmt.Fprintf(c.Writer, "data: %s\n\n", message)
	c.Writer.Flush()
}
