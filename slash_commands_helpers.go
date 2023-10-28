package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// The returning string may contain an error message if the http request could not be done.
func getResultOfHTTPQuery(request string, params url.Values) string {
	domain := "https://pewpew.live/"
	if UseLocalServer {
		fmt.Printf("Making request for %v using local server!\n", request)
		domain = "http://localhost:9999/"
	}

	paramsString := "?" + params.Encode()
	req, createErr := http.NewRequest("GET", domain+request+paramsString, nil)
	if createErr != nil {
		return "Failed to create request: " + createErr.Error()
	}
	req.Header.Add("Bot-Token", BotRequestToken)

	client := &http.Client{}
	resp, getErr := client.Do(req)
	if getErr != nil {
		return "Failed to contact server: " + getErr.Error()
	}

	body, readAllErr := io.ReadAll(resp.Body)
	if readAllErr != nil {
		return "Failed to read body of request: " + readAllErr.Error()
	}

	message := ""
	unmarshallErr := json.Unmarshal(body, &message)
	if unmarshallErr != nil {
		return "Failed to unmarshal JSON: " + unmarshallErr.Error()
	}

	if UseLocalServer {
		fmt.Println("Result:", message)
	}
	return message
}
