package main

import (
	"fmt"
	twiliogo "github.com/carlosdp/twilio-go"
	"os"
)

// func main() {
//     fmt.Println("send sms")

//     sendSms(os.Getenv("ANAHEIM"))
// }

func sendSms(to string, beer string) {

	client := twiliogo.NewClient(os.Getenv("SID"), os.Getenv("AUTH_TOKEN"))

	message := fmt.Sprintf("We've ran out of %s!fdfadkjadh", beer)
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
