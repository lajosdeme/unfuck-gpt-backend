package main

import (
	"fmt"

	"github.com/spf13/viper"
)

var Config ConfigInfo
var G Gpt
var S Saver

func LoadConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	fmt.Println(viper.AllKeys())
	fmt.Println(viper.Get("KEY"))
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
