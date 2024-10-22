package main

import (
	"fmt"
	"io"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	user := "james_liao@openfind.com.tw"
	msg := "Hello, this is a test message."

	resp, err := apiSendMessageUser(user, msg)
	if err != nil {
		fmt.Println("Error sending message:", err)
	} else {
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
		} else {
			fmt.Println("Message sent, response status:", resp.Status)
			fmt.Println("Response body:", string(body))
		}
	}
}
