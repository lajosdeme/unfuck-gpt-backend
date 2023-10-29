package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Gpt struct {
	url   string
	model string
}

func NewGpt(url string, model string) Gpt {
	return Gpt{
		url:   url,
		model: model,
	}
}

func (g *Gpt) GetResponse(prompt string) string {
	msg := ChatGPTApiRequestMessage{
		Role:    string(RoleUser),
		Content: prompt,
	}

	reqBody := ChatGPTApiRequestBody{
		Model:    g.model,
		Messages: []ChatGPTApiRequestMessage{msg},
	}

	reqBodyJson, err := json.Marshal(reqBody)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", g.url, bytes.NewBuffer(reqBodyJson))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", Config.ApiKey))

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var errResp ChatGPTApiErrorResponse
	json.Unmarshal(body, &errResp)
	if errResp.Error.Message != "" {
		panic(err)
	}

	// parse response body
	var chatGptResp ChatGPTApiResponse
	if err := json.Unmarshal(body, &chatGptResp); err != nil {
		panic(err)
	}

	finalReply := chatGptResp.Choices[0].Message.Content

	return finalReply
}

func (g *Gpt) GenerateHeadlinePrompt(headlines HeadlineInfo) string {
	var headlineStr string
	for _, h := range headlines.Headlines {
		headlineStr += fmt.Sprintf("%s\n", h)
	}
	return fmt.Sprintf("You are an esteemed financial analyist writing a report about the cryptocurrency market. Please read these headlines below and write a report on the current sentiment in the crypto market. It should have the following parts: Introduction, Emotions in the Market, Long term predictions, Short term perdictions, Conclusion. Please limit your response to 1000 words: %s", headlineStr)
}

func (g *Gpt) GenerateSentimentPrompt(headlines HeadlineInfo) string {
	var headlineStr string
	for _, h := range headlines.Headlines {
		headlineStr += fmt.Sprintf("%s\n", h)
	}
	return fmt.Sprintf("Please read these headlines, and based on them, rate the average sentiment in the crypto market from 0 (negative, fear) - 10 (positive, greed). Reply in this format \"rating number - short explanation... %s", headlineStr)
}

func (g *Gpt) GeneratePricePrompt(pricesBefore []Price, pricesNow []Price) string {
	priceStr := "Old prices: "
	for _, p := range pricesBefore {
		priceStr += fmt.Sprintf("%s - %.2f ", p.Symbol, p.Price)
	}
	priceStr += "\n"
	priceStr += "New prices: "
	for _, p := range pricesNow {
		priceStr += fmt.Sprintf("%s - %.2f", p.Symbol, p.Price)
	}
	return fmt.Sprintf("Please read these cryptocurrency prices, and compare the old prices to the current. Limit your answer to the comparison only. Do not return the prices: %s", priceStr)
}
