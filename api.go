package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// Struct to hold the message data
type MessageData struct {
	Email string `json:"email"`
	Text  string `json:"text"`
}

func apiSendMessageUser(user, msg string) (*http.Response, error) {
	mmApiHost := os.Getenv("MM_API_HOST")
	mmAppId := os.Getenv("MM_APP_ID")
	mmAppKey := os.Getenv("MM_APP_KEY")

	url := fmt.Sprintf("https://%s/api/messages/sendFromApp?applicationId=%s&clientKey=%s", mmApiHost, mmAppId, mmAppKey)

	data := MessageData{
		Email: user,
		Text:  msg,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	return resp, nil
}