package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go_mailservice/utils"
	"go_mailservice/workers"
	"net/http"
)

func ResetPasswordHandler(c *gin.Context, emailQueue *workers.Queue) {
	var request struct {
		Recipients []struct {
			Email     string `json:"email" binding:"required"`
			Fname     string `json:"fname" binding:"required"`
			Lname     string `json:"lname" binding:"required"`
			ResetLink string `json:"reset_link" binding:"required"`
		} `json:"recipients"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Loop through all recipients and render email for each
	for _, recipient := range request.Recipients {
		htmlBody, err := utils.RenderTemplate("templates/passwordreset.html", map[string]interface{}{
			"Fname":     recipient.Fname,
			"Lname":     recipient.Lname,
			"Email":     recipient.Email,
			"ResetLink": recipient.ResetLink,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to render email template"})
			return
		}

		// Add the personalized email to the queue
		recipientData, _ := json.Marshal([]string{recipient.Email})
		emailQueue.Add(workers.EmailRequest{
			Recipients: recipientData,
			Subject:    "HotHardwareHub - Passwort zurücksetzen",
			Message:    htmlBody,
		})
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Emails added to the queue!"})
}
