package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go_mailservice/workers"
	"net/http"
	"os"
)

func main() {
	emailQueue := make(chan workers.EmailRequest, 100)
	workers.StartWorkerPool(emailQueue, 10)

	router := gin.Default()

	router.POST("/sendmail", func(c *gin.Context) {
		var request workers.EmailRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var recipients []string
		if err := json.Unmarshal(request.Recipients, &recipients); err != nil {
			var singleRecipient string
			if err := json.Unmarshal(request.Recipients, &singleRecipient); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid recipients format. Must be a string or array of strings."})
				return
			}
			recipients = append(recipients, singleRecipient)
		}

		for _, recipient := range recipients {
			recipientData, err := json.Marshal([]string{recipient})
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process recipient"})
				return
			}
			emailQueue <- workers.EmailRequest{
				Recipients: recipientData,
				Subject:    request.Subject,
				Message:    request.Message,
			}
		}

		c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Email(s) added to the queue!"})
	})

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8000" // Default port for localhost usage
	}

	router.Run(":" + httpPort)
}
