package mongoHandler

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

type Message struct {
	ID 			primitive.ObjectID	`bson:"_id,omitempty"`
	UserID	string							`bson:"name,omitempty"`
	Text		string							`bson:"email,omitempty"`
}

func Connect(URL string) {
	c, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(URL))
	if err != nil {
		panic(err)
	}
	client = c
}

func Disconnect() {
	if err := client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}

func FindMessages() []Message {
	cursor, err := client.Database("testing").Collection("messages").Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}
	var messages []Message
	if err := cursor.All(context.TODO(), &messages); err != nil {
		panic(err)
	}
	return messages
}

func InsertOne(message Message) {
	if _, err := client.Database("testing").Collection("messages").InsertOne(context.TODO(), message); err != nil {
		panic(err)
	}
}

func InsertMany(messages []interface{}) {
	if _, err := client.Database("testing").Collection("messages").InsertMany(context.TODO(), messages); err != nil {
		panic(err)
	}
}