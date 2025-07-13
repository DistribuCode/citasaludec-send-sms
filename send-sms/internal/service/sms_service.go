package service

import "log"

func SendSMS(to, message string) error {
	log.Printf("📲 Sending SMS to %s: %s", to, message)
	return nil
}
