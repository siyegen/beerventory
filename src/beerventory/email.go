package main

import (
    "fmt"
    "net/http"
    "os"
    "github.com/sendgrid/sendgrid-go"
)

//func main() {
//    addresses := make([]string, 2)
//    addresses = append(addresses, "suchit@sendgrid.com", "richard.the@sendgrid.com")
//    msg := "Yo! We ran out of Corona..."
//    subject := "Beerventory Update"
//
//    sendEmail(addresses, msg, subject)
//}

func sendEmail(addresses []string, msg string, subject string) {
    sg := sendgrid.NewSendGridClient(os.Getenv("SG_USER"), os.Getenv("SG_KEY"))

    // Because SG's default client times out
    sg.Client = http.DefaultClient

    message := sendgrid.NewMail()

    for _, address := range addresses {
        message.AddTo(address)
    }

    message.SetSubject(subject)
    message.SetText(msg)
    message.SetFrom(os.Getenv("SG_FROM"))
    if r := sg.Send(message); r == nil {
        fmt.Println("Email sent!")
    } else {
        fmt.Println(r)
    }
}
