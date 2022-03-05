package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"

	gomail "gopkg.in/mail.v2"
)

// first seen on https://www.loginradius.com/blog/async/sending-emails-with-golang/

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

	msg, _ := ioutil.ReadFile("msg.txt")

	m.SetBody("text/plain", string(msg))

	d := gomail.NewDialer("smtp.gmail.com", 587, from, pass)

	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}
}
