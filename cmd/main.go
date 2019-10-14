package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/jdkato/prose.v2"
)

func main() {
	doc, _ := prose.NewDocument("What are the guiding principles?")
	// Iterate over the doc's tokens:
	for _, tok := range doc.Tokens() {
		fmt.Println(tok.Text, tok.Tag, tok.Label)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://user:password@winston-database:27017"))
	if err != nil {
		fmt.Println("Can't connect to Mongo")
	}

	var result struct {
		Value string
	}
	filter := bson.M{"keyword": "guiding"}

	collection := client.Database("winston").Collection("questions")
	collection.InsertOne(ctx, bson.M{"keyword": "guiding", "value": "check drive"})
	err = collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)

}
