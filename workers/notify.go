package workers

import (
	"fmt"
	"go_mailservice/utils"
	"os"
)

func notifyAdmin(failedRequest EmailRequest) {
	adminEmail := os.Getenv("ADMIN_EMAIL")
	if adminEmail == "" {
		fmt.Println("Admin email not configured. Skipping notification.")
		return
	}

	subject := fmt.Sprintf("Failed Email Delivery: %s", failedRequest.Recipient)
	message := fmt.Sprintf(
		"Failed to deliver email to %s after 3 attempts.\n\nSubject: %s\nMessage: %s",
		failedRequest.Recipient, failedRequest.Subject, failedRequest.Message,
	)

	err := utils.SendEmail(adminEmail, subject, message)
	if err != nil {
		fmt.Printf("Failed to notify admin about email failure: %s\n", err)
	} else {
		fmt.Printf("Admin notified about email failure to: %s\n", failedRequest.Recipient)
	}
}
