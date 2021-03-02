package notifications

import (
	"net/smtp"

	"code.jtg.tools/ayush.singhal/notifications-microservice/configuration"
	"github.com/pkg/errors"
)

// Notifications interface is used to send different types of notificaitons
type Notifications interface {
	SendNotification() error
	NewNotification(to string, title string, body string)
}

// Email struct implements Notifications interface
type Email struct {
	To      string
	Subject string
	Message string
}

// NewNotification creates fills the values in the struct with the provided ones
func (email *Email) NewNotification(to string, title string, body string) {
	email.Message = body
	email.To = to
	email.Subject = title
}

// SendNotification method send email notifications
func (email *Email) SendNotification() error {
	from := configuration.GetResp().EmailNotification.Email
	password := configuration.GetResp().EmailNotification.Password
	smtpHost := configuration.GetResp().EmailNotification.SMTPHost
	smtpPort := configuration.GetResp().EmailNotification.SMTPPort
	addr := smtpHost + ":" + smtpPort
	msg := []byte("Subject: " + email.Subject + "\r\n" +
		"\r\n" + email.Message + "\r\n")

	// Authentication
	auth := smtp.PlainAuth("", from, password, addr)

	//  Sending email.
	err := smtp.SendMail(addr, auth, from, []string{email.To}, msg)
	if err != nil {
		return errors.Wrap(err, "Unable to send email")
	}
	return nil
}
