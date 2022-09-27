package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kzeratal/cinnox-homework/internal/ginHandler"
	"github.com/kzeratal/cinnox-homework/internal/lineHandler"
	"github.com/kzeratal/cinnox-homework/internal/mongoHandler"
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

	for _, message := range(mongoHandler.FindMessages()) {
		fmt.Println(message)
	}

	secret := fmt.Sprintf("%v", viper.Get("line.secret"))
	accessToken := fmt.Sprintf("%v", viper.Get("line.accessToken"))
	lineHandler.Init(secret, accessToken)

	server := gin.Default()
	server.POST("/receive", ginHandler.Revceive)
	server.Run()
}