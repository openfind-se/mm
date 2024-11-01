package mm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

// Struct to hold the message data
type MessageData struct {
	Email string `json:"email"`
	Text  string `json:"text"`
}

// Create a global HTTP client with proxy settings
func CreateHttpClient() *http.Client {
	proxyURLStr := os.Getenv("HTTP_PROXY")
	if proxyURLStr != "" {
		proxyURL, err := url.Parse(proxyURLStr)
		if err != nil {
			fmt.Println("Error parsing proxy URL:", err)
			return &http.Client{}
		}
		transport := &http.Transport{Proxy: http.ProxyURL(proxyURL)}
		return &http.Client{Transport: transport}
	}
	return &http.Client{}
}

var httpClient = CreateHttpClient()

func SendMessageUser(user, msg string) (*http.Response, error) {
	mmApiHost := os.Getenv("MM_API_HOST")
	mmAppId := os.Getenv("MM_APP_ID")
	mmAppKey := os.Getenv("MM_APP_KEY")

	urlStr := fmt.Sprintf("https://%s/api/messages/sendFromApp?applicationId=%s&clientKey=%s", mmApiHost, mmAppId, mmAppKey)

	data := MessageData{
		Email: user,
		Text:  msg,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", urlStr, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
