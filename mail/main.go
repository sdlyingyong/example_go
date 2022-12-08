package main

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

func main() {
	// Create the message.
	m := gomail.NewMessage()
	m.SetHeader("From", "your_qq_email@qq.com")
	m.SetHeader("To", "to_email@163.com")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/plain", "Hi!\nThis is an email send by golang.\n")

	// Create the dialer to connect to the server.
	d := gomail.NewDialer("smtp.qq.com", 465, "your_qq_account", "your_qq_mail_smtp_password")

	// Send the email.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

	fmt.Println("Email sent!")
}
