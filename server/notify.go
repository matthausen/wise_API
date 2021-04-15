package server

import (
	"fmt"
	"log"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

func GoDotEnvVariable(key string) string {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func Notify(rate float32) {
	r := fmt.Sprintf("%f", rate)
	emailAddress := GoDotEnvVariable("EMAIL_ADDRESS")
	password := GoDotEnvVariable("EMAIL_PASSWORD")

	from := emailAddress
	pass := password
	to := emailAddress

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Wise API Notification: \n\n" +
		"Today the exchange rate for GBP to EUR is" + r + ".\n Good time to transfer"

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}
}
