package main

import (
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

func sendEmail(subject string, body string, to string) {
	godotenv.Load()

	message := gomail.NewMessage()
	message.SetHeader("From", os.Getenv("MYEMAIL"))
	message.SetHeader("To", to)
	message.SetHeader("Subject", subject)
	message.SetBody("text/html", body)
	message.Attach("./Siri_Marvin_Resume.pdf")

	d := gomail.NewDialer("smtp.gmail.com", 587, os.Getenv("MYEMAIL"), os.Getenv("PASSWORD"))

	if err := d.DialAndSend(message); err != nil {
		panic(err)
	}
}

func main() {
	sendEmail("here", "you go", "marvinsiri123@gmail.com")
}
