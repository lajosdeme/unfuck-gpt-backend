package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

var Config ConfigInfo
var G Gpt
var S Saver

func LoadConfig() {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		fmt.Println("error getting port: ", err)
	}
	if os.Getenv("MODE") == "heroku" {
		Config = ConfigInfo{
			ApiKey:        os.Getenv("API_KEY"),
			FlareNode:     os.Getenv("FLARE_NODE_URL"),
			MumbaiUrl:     os.Getenv("MUMBAI_URL"),
			PriceDataAddr: os.Getenv("UNFUCK_PRICE_DATA_ADDRESS"),
			UnfuckCore:    os.Getenv("UNFUCK_CORE_ADDRESS"),
			Key:           os.Getenv("KEY"),
			Port:          port,
		}

		G = NewGpt("http://flock.tools:8001/v1/chat/completions", "hackathon-chat")
		S = NewSaver()
		return
	}
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()

	if err != nil {
		panic(fmt.Sprintf("Failed to read env file: %s", err.Error()))
	}

	err = viper.Unmarshal(&Config)

	if err != nil {
		panic(fmt.Sprintf("Failed to decode env file: %s", err.Error()))
	}

	G = NewGpt("http://flock.tools:8001/v1/chat/completions", "hackathon-chat")
	S = NewSaver()
}
