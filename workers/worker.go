package workers

import (
	"encoding/json"
	"fmt"
	"go_mailservice/utils"
	"time"
)

type EmailRequest struct {
	Recipients json.RawMessage `json:"recipients" binding:"required"`
	Subject    string          `json:"subject" binding:"required"`
	Message    string          `json:"message" binding:"required"`
}

func StartWorkerPool(emailQueue chan EmailRequest, workerCount int) {
	for i := 0; i < workerCount; i++ {
		go worker(i, emailQueue)
	}
}

func worker(id int, emailQueue chan EmailRequest) {
	for request := range emailQueue {
		var recipients []string
		if err := json.Unmarshal(request.Recipients, &recipients); err != nil {
			var singleRecipient string
			if err := json.Unmarshal(request.Recipients, &singleRecipient); err != nil {
				continue
			}
			recipients = append(recipients, singleRecipient)
		}

		for _, recipient := range recipients {
			sendEmailWithRetry(EmailRequest{
				Recipients: json.RawMessage(fmt.Sprintf(`"%s"`, recipient)),
				Subject:    request.Subject,
				Message:    request.Message,
			}, 3)
		}
	}
}

func sendEmailWithRetry(request EmailRequest, retries int) {
	var recipient string
	if err := json.Unmarshal(request.Recipients, &recipient); err != nil {
		fmt.Printf("Invalid recipient format: %v\n", err)
		return
	}

	for i := 0; i < retries; i++ {
		err := utils.SendEmail(recipient, request.Subject, request.Message)
		if err == nil {
			fmt.Printf("Email successfully sent to: %s\n", recipient)
			return
		}
		fmt.Printf("Retrying email to %s (attempt %d/%d)\n", recipient, i+1, retries)
		time.Sleep(2 * time.Second)
	}
	fmt.Printf("Failed to send email to %s after %d retries\n", recipient, retries)
	notifyAdmin(request)
}
