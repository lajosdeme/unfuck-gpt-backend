package main

type Role string

const (
	RoleSystem    Role = "system"
	RoleUser      Role = "user"
	RoleAssistant Role = "assistant"
)

type ConfigInfo struct {
	ApiKey        string `mapstructure:"API_KEY"`
	FlareNode     string `mapstructure:"FLARE_NODE_URL"`
	MumbaiUrl     string `mapstructure:"MUMBAI_URL"`
	PriceDataAddr string `mapstructure:"UNFUCK_PRICE_DATA_ADDRESS"`
	UnfuckCore    string `mapstructure:"UNFUCK_CORE_ADDRESS"`
	Key           string `mapstructure:"KEY"`
}

type ReportInfo struct {
	Tstamp int64  `json:"tstamp"`
	Report string `json:"report"`
}

type Sentiment struct {
	Tstamp    int64  `json:"tstamp"`
	Sentiment string `json:"sentiment"`
}

type HeadlineInfo struct {
	Time      string   `json:"time"`
	Headlines []string `json:"headlines"`
}

type PriceInfo struct {
	Tstamp int64   `json:"tstamp"`
	Prices []Price `json:"prices"`
}

type ChatGPTApiRequestMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatGPTApiRequestBody struct {
	Model    string                     `json:"model"`
	Messages []ChatGPTApiRequestMessage `json:"messages"`
}

type ChatGPTApiResponse struct {
	Choices []Choice `json:"choices"`
}

type ChatGPTApiErrorResponse struct {
	Error ChatGPTApiError `json:"error"`
}

type ChatGPTApiError struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

type Choice struct {
	Index   int                      `json:"index"`
	Message ChatGPTApiRequestMessage `json:"message"`
}
