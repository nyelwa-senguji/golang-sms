package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/twilio/twilio-go"

	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

func main() {

	// load .env file from given path
	// we keep it empty it will load .env from current directory
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")
	from := os.Getenv("TWILIO_FROM_PHONE_NUMBER")
	to := os.Getenv("TWILIO_TO_PHONE_NUMBER")

	client := twilio.NewRestClientWithParams(accountSid, authToken, twilio.RestClientParams{
		AccountSid: accountSid,
	})

	params := &openapi.CreateMessageParams{}
	params.SetTo(to)
	params.SetFrom(from)
	params.SetBody("Hello there")

	resp, err := client.ApiV2010.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
		err = nil
	} else {
		response, _ := json.Marshal(*resp)
		fmt.Println("Response: " + string(response))
	}
}
