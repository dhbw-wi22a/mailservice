package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go_mailservice/utils"
	"go_mailservice/workers"
	"net/http"
)

func ConfOrderHandler(c *gin.Context, emailQueue *workers.Queue) {
	var request struct {
		Recipients []struct {
			Email string `json:"email" binding:"required"`
			Name  string `json:"name" binding:"required"`
		} `json:"recipients"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Loop through all recipients and render email for each
	for _, recipient := range request.Recipients {
		htmlBody, err := utils.RenderTemplate("templates/orderconf.html", map[string]interface{}{
			"Name":  recipient.Name,
			"Email": recipient.Email,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to render email template"})
			return
		}

		// Add the personalized email to the queue
		recipientData, _ := json.Marshal([]string{recipient.Email})
		emailQueue.Add(workers.EmailRequest{
			Recipients: recipientData,
			Subject:    "Deine Bestellung bei HotHardwareHub!",
			Message:    htmlBody,
		})
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Emails added to the queue!"})
}
