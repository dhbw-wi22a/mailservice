package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go_mailservice/utils"
	"go_mailservice/workers"
	"net/http"
)

func SendInvoiceHandler(c *gin.Context, emailQueue *workers.Queue) {
	var request workers.EmailRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	recipients, err := utils.ParseRecipients(request.Recipients)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, recipient := range recipients {
		recipientData, _ := json.Marshal([]string{recipient})
		emailQueue.Add(workers.EmailRequest{
			Recipients: recipientData,
			Subject:    request.Subject,
			Message:    request.Message,
		})
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Email(s) added to the queue!"})
}
