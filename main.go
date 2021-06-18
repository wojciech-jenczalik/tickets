package main

import (
	"fmt"
	"io"
	"net/http"
	"net/smtp"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println(os.Getenv("GMAIL_PASS"))
	for true {
		resp, err := http.Get("https://www.stockholmlive.com/en/events/detail/csgo-major")
		if(err != nil) {
			fmt.Println("Error while visiting website")
			continue
		}
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)

		if strings.Contains(string(body), "no_ticket_link") {
			fmt.Println("No tickets")
		} else {
			from := "ticket.maintain.fpl.major@gmail.com"
			password := os.Getenv("GMAIL_PASS")

			to := []string{
				"waflu918@gmail.com",
				"michalbabinski@protonmail.ch",
			}

			smtpHost := "smtp.gmail.com"
			smtpPort := "587"

			message := []byte("BILETY DOSTĘPNE")

			auth := smtp.PlainAuth("Tickets FPL", from, password, smtpHost)

			err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("Notification about tickets sent successfully")
		}

		time.Sleep(30 * time.Minute)
	}
}
