package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func StreamSummaryHandler(c *gin.Context) {
	text := c.Query("text")

	if text == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "text query parameter is required",
		})
		return
	}

	// Set SSE headers
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")
	c.Writer.Header().Set("Transfer-Encoding", "chunked")

	streamSummary(text, c)
}
