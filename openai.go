package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func streamSummary(text string, c *gin.Context) {
	apiKey := os.Getenv("OPENAI_API_KEY")

	url := "https://api.openai.com/v1/responses"

	payload := []byte(fmt.Sprintf(`{
		"model": "gpt-4.1-mini",
		"input": "Summarize this text:\n\n%s",
		"stream": true
	}`, text))

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		sendSSE(c, "ERROR: "+err.Error())
		return
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)

	for scanner.Scan() {
		line := scanner.Text()

		// SSE format: lines start with "data: "
		if len(line) > 6 && line[:6] == "data: " {
			data := line[6:]

			if data == "[DONE]" {
				break
			}

			sendSSE(c, data)
		}
	}
}
