package main

import (
	"go_mailservice/utils"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type EmailRequest struct {
	Recipient string `json:"recipient" binding:"required"`
	Subject   string `json:"subject" binding:"required"`
	Message   string `json:"message" binding:"required"`
}

func main() {

	router := gin.Default()

	router.POST("/sendmail", func(c *gin.Context) {
		var request EmailRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := utils.SendEmail(request.Recipient, request.Subject, request.Message)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Email sent successfully!"})
	})

	httpPort := os.Getenv("HTTP_PORT")

	if httpPort == "" {
		httpPort = "8000" //for localhost usage
	}

	router.Run(":" + httpPort)
}
