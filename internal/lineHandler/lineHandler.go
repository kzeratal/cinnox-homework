package lineHandler

import (
	"net/http"

	"github.com/kzeratal/cinnox-homework/internal/mongoHandler"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

var bot *linebot.Client

func Init(secret string, acessToken string) {
	client, err := linebot.New(secret, acessToken)
	if err != nil {
		panic(err)
	}
	bot = client
}

func GetMessages(req *http.Request) []interface{} {
	events, err := bot.ParseRequest(req)
	if err != nil {
		panic(err)
	}
	messages := []interface{}{}
	for _, event := range events {
			switch event.Message.(type) {
				case *linebot.TextMessage:
					message := mongoHandler.Message{
						UserID: event.Source.UserID,
						Text: event.Message.(*linebot.TextMessage).Text,
					}
					messages = append(messages, message)
		}
	}
	return messages
}

func Broadcast(text string) {
	message := linebot.NewTextMessage(text)
	bot.BroadcastMessage(message).Do()
}