package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload" // underscore means import solely for side effects
	"github.com/lmittmann/tint"
)

func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my website!")
}

func database(w http.ResponseWriter, r *http.Request) {
	doSomething()
}

// middleware pattern:
//
//	wrap a route call in a chain of middleware functions
//	route fn called after all middleware fns called
func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Debug(r.URL.Path)
		f(w, r)
	}
}

func main() {
	initLogging()

	slog.Info(getEnvConfig().notionApiKey)

	slog.Info("Starting server on port 8080")
	http.HandleFunc("/", logging(welcome))
	http.HandleFunc("/database", logging(database))
	// getPagesFromDB(getEnvConfig().notionDbId)
	getAllPagesAsNotionNotes()

	// http.ListenAndServe(":8080", nil)
}

func getAllPagesAsNotionNotes() {
	pages := getPagesFromDB(getEnvConfig().notionDbId)

	notionNotes := []NotionNote{}

	for _, page := range pages {
		notionNotes = append(notionNotes, parsePageToNotionNote(page))
	}

	slog.Info(getJSON(notionNotes))

}

func initLogging() {
	slog.SetDefault(slog.New(
		tint.NewHandler(os.Stderr, &tint.Options{
			Level:      slog.LevelDebug,
			TimeFormat: time.Kitchen,
		}),
	))
}
