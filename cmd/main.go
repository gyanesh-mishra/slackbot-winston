package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gyanesh-mishra/slackbot-winston/config"
	"github.com/gyanesh-mishra/slackbot-winston/internal/dao"
	"github.com/gyanesh-mishra/slackbot-winston/internal/routing"
)

func main() {

	// Get configuration object
	configuration := config.GetConfig()

	// Create all database indexes
	err := dao.CreateAllIndexes()
	if err != nil {
		log.Print(fmt.Sprintf("Error creating indexes %+v", err))
	}

	// Get the router object
	router := routing.GetRouter()

	// Start the server
	log.Print(fmt.Sprintf("Running server on %s", configuration.Server.Host))
	log.Fatal(http.ListenAndServe(configuration.Server.Host, router))

}
