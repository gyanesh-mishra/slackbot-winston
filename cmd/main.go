package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gyanesh-mishra/slackbot-winston/config"
	"github.com/gyanesh-mishra/slackbot-winston/internal/routing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"gopkg.in/jdkato/prose.v2"
)

func main() {
	configuration := config.GetConfig()

	doc, _ := prose.NewDocument("What is our time off policy?", prose.WithExtraction(false))
	router := routing.GetRouter()
	// Iterate over the doc's tokens:
	for _, tok := range doc.Tokens() {
		fmt.Println(tok.Text, tok.Tag)
		if tok.Tag == "NNS" {
			fmt.Println("Is noun")
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	databaseString := fmt.Sprintf("mongodb://%s:%s@%s:%d", configuration.Database.DBUser, configuration.Database.DBPassword, configuration.Database.DBHost, configuration.Database.DBPort)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(databaseString))
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

	log.Print(fmt.Sprintf("Running server on %s", configuration.Server.Host))
	log.Fatal(http.ListenAndServe(configuration.Server.Host, router))

}
