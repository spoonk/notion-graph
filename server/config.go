package main

import (
	_ "github.com/joho/godotenv/autoload" // underscore means import solely for side effects
	"os"
)

type EnvConfig struct {
	notionApiKey string
	notionDbId   string
}

func getEnvConfig() EnvConfig {
	return EnvConfig{
		notionApiKey: os.Getenv("NOTION_API_KEY"),
		notionDbId:   os.Getenv("NOTION_DB_ID"),
	}
}
