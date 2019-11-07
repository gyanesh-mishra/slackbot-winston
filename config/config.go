package config

import (
	"context"
	"fmt"
	"strings"

	"github.com/nlopes/slack"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Configuration for the application
type Configuration struct {
	Server   ServerConfiguration
	Database DatabaseConfiguration
	Slack    SlackConfiguration
}

// ServerConfiguration holds the webserver configuration
type ServerConfiguration struct {
	Host string
}

// DatabaseConfiguration holds the database connection configuration
type DatabaseConfiguration struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     int
	Client     *mongo.Client
}

// SlackConfiguration holds the slack app connection configuration
type SlackConfiguration struct {
	BotToken          string
	VerificationToken string
	Client            *slack.Client
}

// GetConfig reads the configuration file and returns a Configuration object
func GetConfig() Configuration {
	// Set the file name of the configurations file
	viper.SetConfigName("config/config")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")

	// Enable VIPER to read Environment Variables
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.SetEnvPrefix("winston")
	viper.AutomaticEnv()

	viper.SetConfigType("yml")
	var configuration Configuration

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode config into struct, %v", err)
	}

	// Add database connection client to the global configuration object
	configuration.Database.Client = getDatabaseConnection(configuration)
	// Add slack client to global configuration object
	configuration.Slack.Client = slack.New(configuration.Slack.BotToken)

	return configuration
}

// Private function to get database connection based on config
func getDatabaseConnection(configuration Configuration) *mongo.Client {

	// Get database connection
	databaseString := fmt.Sprintf("mongodb://%s:%s@%s:%d", configuration.Database.DBUser, configuration.Database.DBPassword, configuration.Database.DBHost, configuration.Database.DBPort)
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(databaseString))
	if err != nil {
		fmt.Println("Can't connect to Mongo")
	}

	return client
}
