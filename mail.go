package main

import (
	"errors"
	"fmt"
	"net/smtp"
)

func sendMail(host string, port uint16, username, password, subject, body string, recipients []string) error {
	if len(recipients) == 0 {
		return errors.New("no recipients found on e-mail notification")
	}

	auth := smtp.PlainAuth("", username, password, host)
	msg := fmt.Sprintf(
		"From: %v\r\nTo: %v\r\nSubject: %v\r\n\r\n%v",
		username,
		username,
		subject,
		body,
	)
	err := smtp.SendMail(
		fmt.Sprintf("%v:%v", host, port),
		auth,
		username,
		recipients,
		[]byte(msg),
	)
	if err != nil {
		return fmt.Errorf("error sending email: %w", err)
	}

	return nil
}
