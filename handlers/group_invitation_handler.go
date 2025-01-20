package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go_mailservice/utils"
	"go_mailservice/workers"
	"net/http"
)

func GroupInvitationHandler(c *gin.Context, emailQueue *workers.Queue) {
	var request struct {
		Recipients []struct {
			Email          string `json:"email" binding:"required"`
			Fname          string `json:"fname" binding:"required"`
			Lname          string `json:"lname" binding:"required"`
			GroupName      string `json:"group_name" binding:"required"`
			InvitedBy      string `json:"invited_by" binding:"required"`
			InvitationLink string `json:"invitation_link" binding:"required"`
		} `json:"recipients"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Loop through all recipients and render email for each
	for _, recipient := range request.Recipients {
		htmlBody, err := utils.RenderTemplate("templates/groupinvitation.html", map[string]interface{}{
			"Fname":     recipient.Fname,
			"Lname":     recipient.Lname,
			"Email":     recipient.Email,
			"GroupName": recipient.GroupName,
			"InvitedBy": recipient.InvitedBy,
			"Accept":    fmt.Sprintf("%s/accept/", recipient.InvitationLink),
			"Decline":   fmt.Sprintf("%s/decline/", recipient.InvitationLink),
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to render email template"})
			return
		}

		// Add the personalized email to the queue
		recipientData, _ := json.Marshal([]string{recipient.Email})
		emailQueue.Add(workers.EmailRequest{
			Recipients: recipientData,
			Subject:    "Du wurdest zu einer Gruppe eingeladen - jetzt bestätigen!",
			Message:    htmlBody,
		})
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Emails added to the queue!"})
}
