package main

import (
	_ "github.com/joho/godotenv/autoload" // underscore means import solely for side effects
	"os"
)

func getEnvConfig() EnvConfig {
	return EnvConfig{
		notionApiKey: os.Getenv("NOTION_API_KEY"),
	}
}
