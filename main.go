package main

import (
	"fmt"

	"github.com/kzeratal/cinnox-homework/internal/mongoHandler"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	URL := fmt.Sprintf("%v", viper.Get("mongo.URL"))
	mongoHandler.Connect(URL)
	defer mongoHandler.Disconnect()
	fmt.Println("Connected to MongoDB")

	message := mongoHandler.Message{
		UserID: "ID",
		Text: 	"test",
	}
	mongoHandler.InsertOne(message)
	for _, message := range(mongoHandler.FindMessages()) {
		fmt.Println(message)
	}

	channelSecret := fmt.Sprintf("%v", viper.Get("channel.Secret"))
	channelAccessToken := fmt.Sprintf("%v", viper.Get("channel.accessToken"))
	if _, err := linebot.New(channelSecret, channelAccessToken); err != nil {
		panic(err)
	}
}