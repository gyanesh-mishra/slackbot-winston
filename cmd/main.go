package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gyanesh-mishra/slackbot-winston/config"
	"github.com/gyanesh-mishra/slackbot-winston/internal/routing"
)

func main() {

	// Get configuration object
	configuration := config.GetConfig()

	// Get the router object
	router := routing.GetRouter()

	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	// client := configuration.Database.Client

	// var result struct {
	// 	Value string
	// }
	// filter := bson.M{"keyword": "guiding"}

	// collection := client.Database("winston").Collection("questions")
	// collection.InsertOne(ctx, bson.M{"keyword": "guiding", "value": "check drive"})
	// err := collection.FindOne(ctx, filter).Decode(&result)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(result)

	log.Print(fmt.Sprintf("Running server on %s", configuration.Server.Host))
	log.Fatal(http.ListenAndServe(configuration.Server.Host, router))

}
