package util

import (
	"fmt"
	"net/smtp"
	"strconv"
	"strings"
)

func SendEmail(server string, port int, sender string, nick string, password string, recipient []string, contentType string, subject string, body string) {
	auth := smtp.PlainAuth("", sender, password, server)
	msg := []byte("To: " + strings.Join(recipient, ",") + "\r\nFrom: " + nick +
		"<" + sender + ">\r\nSubject: " + subject + "\r\n" + contentType + "\r\n\r\n" + body)
	err := smtp.SendMail(server+strconv.Itoa(port), auth, sender, recipient, msg)
	if err != nil {
		fmt.Printf("send mail error: %v\n", err)
	}
}
