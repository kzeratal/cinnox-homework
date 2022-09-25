package main

import (
	"context"
	"fmt"
	"time"

	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	URL := fmt.Sprintf("%v", viper.Get("mongo.URL"))
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	if _, err := mongo.Connect(ctx, options.Client().ApplyURI(URL)); err != nil {
		panic(err)
	}
	channelSecret := fmt.Sprintf("%v", viper.Get("channel.Secret"))
	channelAccessToken := fmt.Sprintf("%v", viper.Get("channel.accessToken"))
	if _, err := linebot.New(channelSecret, channelAccessToken); err != nil {
		panic(err)
	}
}