package main

import (
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

		emailQueue <- request
		c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Email added to the queue!"})
	})

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8000" // Default port for localhost usage
	}

	router.Run(":" + httpPort)
}
