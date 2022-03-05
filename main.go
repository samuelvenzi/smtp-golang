package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"

	gomail "gopkg.in/mail.v2"
)

func main() {
	sendTo("example@to.com")
}

func sendTo(to string) {
	m := gomail.NewMessage()

	from := "example@from.com"
	pass := "passoutintheopen-becareful"
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "My Subject")

	// read from msg.txt
	msg, _ := ioutil.ReadFile("msg.txt")

	m.SetBody("text/plain", string(msg))

	// Settings for SMTP server
	d := gomail.NewDialer("smtp.gmail.com", 587, from, pass)

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}
}
