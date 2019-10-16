package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
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

	return configuration
}
