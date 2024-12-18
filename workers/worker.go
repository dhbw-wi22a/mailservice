package workers

import (
	"errors"
	"fmt"
	_ "go_mailservice/utils"
	"time"
)

type EmailRequest struct {
	Recipient string
	Subject   string
	Message   string
}

func StartWorkerPool(emailQueue chan EmailRequest, workerCount int) {
	for i := 0; i < workerCount; i++ {
		go worker(i, emailQueue)
	}
}

func worker(id int, emailQueue chan EmailRequest) {
	for request := range emailQueue {
		fmt.Printf("[Worker %d] Processing email to: %s\n", id, request.Recipient)
		sendEmailWithRetry(request, 3)
	}
}

func sendEmailWithRetry(request EmailRequest, retries int) {
	for i := 0; i < retries; i++ {
		err := errors.New("simulated failure")
		//err := utils.SendEmail(request.Recipient, request.Subject, request.Message)
		if err == nil {
			fmt.Printf("Email successfully sent to: %s\n", request.Recipient)
			return
		}
		fmt.Printf("Retrying email to %s (attempt %d/%d)\n", request.Recipient, i+1, retries)
		time.Sleep(2 * time.Second)
	}
	fmt.Printf("Failed to send email to %s after %d retries\n", request.Recipient, retries)
	notifyAdmin(request)
}
