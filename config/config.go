package config

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Configuration exported
type Configuration struct {
	Server   ServerConfiguration
	Database DatabaseConfiguration
	Slack    SlackConfiguration
}

// ServerConfiguration exported
type ServerConfiguration struct {
	Host string
}

// DatabaseConfiguration exported
type DatabaseConfiguration struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     int
	Client     *mongo.Client
}

// SlackConfiguration exported
type SlackConfiguration struct {
	BotToken          string
	VerificationToken string
}

// GetConfig exported
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

	return configuration
}

// Private function to get database connection based on config
func getDatabaseConnection(configuration Configuration) *mongo.Client {
	// Get application context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Get database connection
	databaseString := fmt.Sprintf("mongodb://%s:%s@%s:%d", configuration.Database.DBUser, configuration.Database.DBPassword, configuration.Database.DBHost, configuration.Database.DBPort)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(databaseString))
	if err != nil {
		fmt.Println("Can't connect to Mongo")
	}

	return client
}
