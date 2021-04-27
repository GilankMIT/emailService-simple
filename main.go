package main

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
	"net/http"
)

const (
	CONFIG_SMTP_HOST     = "srv58.niagahoster.com"
	CONFIG_SMTP_PORT     = 465
	CONFIG_SENDER_NAME   = "Test <system-one@pied-piper-coin.xyz>"
	CONFIG_AUTH_EMAIL    = "system-one@pied-piper-coin.xyz"
	CONFIG_AUTH_PASSWORD = "root12345"
)

type SendMailRequest struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Content string `json:"content"`
}

func main() {
	r := gin.Default()
	r.POST("/send-email", func(c *gin.Context) {
		var request SendMailRequest
		err := c.ShouldBindJSON(&request)
		if err != nil {
			c.JSON(400, gin.H{
				"message": "invalid request body",
			})
		}

		err = sendMail(request.To, request.Subject, request.Content)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "invalid request body",
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"message":   "email sent successfully",
			"recipient": request.To,
		})
	})

	r.Run("localhost:9001")
}

func sendMail(to, subject, content string) error {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
	mailer.SetHeader("To", to)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", content)

	dialer := gomail.NewDialer(
		CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT,
		CONFIG_AUTH_EMAIL,
		CONFIG_AUTH_PASSWORD,
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		return err
	}

	return nil
}
