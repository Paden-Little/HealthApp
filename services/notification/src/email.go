package main

/*
"Service","Name","Username","Password","Hostname","Port","Security"
"SMTP","Tom Hamill","tom.hamill@ethereal.email","X9VjtwqtRpGJYBZ8jx","smtp.ethereal.email",587,"STARTTLS"
"IMAP","Tom Hamill","tom.hamill@ethereal.email","X9VjtwqtRpGJYBZ8jx","imap.ethereal.email",993,"TLS"
"POP3","Tom Hamill","tom.hamill@ethereal.email","X9VjtwqtRpGJYBZ8jx","pop3.ethereal.email",995,"TLS"
*/

import (
	"fmt"
	"net/smtp"
)

func SendEmail() {
	// Sender data.
	from := "tom.hamill@ethereal.email"
	password := "X9VjtwqtRpGJYBZ8jx"

	// Receiver email address.
	to := []string{
		"sender@example.com",
	}

	// smtp server configuration.
	smtpHost := "smtp.ethereal.email"
	smtpPort := "587"

	// Message.
	message := []byte("This is a test email message.")

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")
}
