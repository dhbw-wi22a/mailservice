package main

import (
	"github.com/gin-gonic/gin"
	"go_mailservice/handlers"
	"go_mailservice/workers"
	"os"
)

func main() {
	emailQueue := workers.NewQueue(100)
	workers.StartWorkerPool(emailQueue, 10)

	router := gin.Default()

	mailservice := router.Group("/mailservice")
	{
		mailservice.POST("/registration", func(c *gin.Context) {
			handlers.RegistrationHandler(c, emailQueue)
		})
		mailservice.POST("/orderconf", func(c *gin.Context) {
			handlers.ConfOrderHandler(c, emailQueue)
		})
		mailservice.POST("/notifyshipment", func(c *gin.Context) {
			handlers.NotifyShipmentHandler(c, emailQueue)
		})
		mailservice.POST("/sendinvoice", func(c *gin.Context) {
			handlers.SendInvoiceHandler(c, emailQueue)
		})
		mailservice.POST("/groupinvitation", func(c *gin.Context) {
			handlers.GroupInvitationHandler(c, emailQueue)
		})
		mailservice.POST("/resetpassword", func(c *gin.Context) {
			handlers.ResetPasswordHandler(c, emailQueue)
		})
		mailservice.POST("/setinactive", func(c *gin.Context) {
			handlers.SetInactiveHandler(c, emailQueue)
		})

	}

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8000" // Default port for localhost usage (check .env file)
	}

	router.Run(":" + httpPort)
}
