package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Define the structure for the API request body
type OpenAIRequest struct {
	Prompt  string `json:"prompt"`
	Model   string `json:"model"`
	MaxTokens int `json:"max_tokens"`
}

// Define the structure for the API response
type OpenAIResponse struct {
	Choices []struct {
		Text string `json:"text"`
	} `json:"choices"`
}

// Function to summarize content using OpenAI's ChatGPT
func SummarizeContent(content string) (string, error) {
	requestBody := OpenAIRequest{
		Prompt:  content,
		Model:   "gpt-3.5-turbo-instruct", // Or the latest available model
		MaxTokens: 1000,
	}
	requestBodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}

	request, err := http.NewRequest("POST", "https://api.openai.com/v1/completions", bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return "", err
	}

	apikey, err := ReadAPIKey()
	if err != nil {
		fmt.Println("Error reading API key:", err)
		return "", err
	}

	// Replace "Your_OpenAI_API_Key" with your actual OpenAI API key
	request.Header.Set("Authorization", "Bearer " + apikey)
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	var openAIResponse OpenAIResponse
	err = json.Unmarshal(body, &openAIResponse)
	if err != nil {
		return "", err
	}

	if len(openAIResponse.Choices) > 0 {
		return openAIResponse.Choices[0].Text, nil
	}

	return "", fmt.Errorf("no summary received from OpenAI")
}

func ReadAPIKey() (string, error) {
	data, err := ioutil.ReadFile("./apikey.txt")
	if err != nil {
		// Return an empty string and the error if the file cannot be read
		return "", err
	}

	// Convert the data to a string and trim any whitespace
	apiKey := strings.TrimSpace(string(data))
	return apiKey, nil
}