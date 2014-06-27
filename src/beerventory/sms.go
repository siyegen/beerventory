package main

import (
    "fmt"
    "os"
    twiliogo "github.com/carlosdp/twilio-go"
)

func main() {
    fmt.Println("send sms")

    sendSms(os.Getenv("ANAHEIM"))
}

func sendSms(to string) {

    client := twiliogo.NewClient(os.Getenv("SID"), os.Getenv("AUTH_TOKEN"))

    message := "Get me some beer, running out of it!"
    _, err := twiliogo.NewMessage(client,
                 os.Getenv("FROM"),
                 to,
                 twiliogo.Body(message))
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Sprintf("Message sent!")
    }
}
