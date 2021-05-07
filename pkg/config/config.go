package config

import (
	"log"
	"os"
)

type Config struct {
	TypesenseApiUrl string
	TypesenseApiKey string
	WebhookApiKey   string
}

var CurrentConfig Config

func InitializeConfig() {
	CurrentConfig = Config{
		TypesenseApiUrl: os.Getenv("TYPESENSE_API_URL"),
		TypesenseApiKey: os.Getenv("TYPESENSE_API_KEY"),
		WebhookApiKey:   os.Getenv("WEBHOOK_API_KEY"),
	}
	log.Println("Typesense API URL: " + os.Getenv("TYPESENSE_API_URL"))
}
