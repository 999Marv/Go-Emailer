package main

import (
	"fmt"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

func sendEmail() {
	godotenv.Load()
	auth := smtp.PlainAuth(
		"",
		os.Getenv("MYEMAIL"),
		os.Getenv("PASSWORD"),
		"smtp.gmail.com",
	)

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		os.Getenv("MYEMAIL"),
		[]string{os.Getenv("MYEMAIL")},
		[]byte("Subject: HI\nThis is you"),
	)

	if err != nil {
		fmt.Println(err)
	}

}

func main() {
	sendEmail()
}
