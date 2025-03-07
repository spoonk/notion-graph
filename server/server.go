package main

import (
	"encoding/json"
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
func getGraph(w http.ResponseWriter, r *http.Request) {
	notes := getAllPagesAsNotionNotes()
	g := buildGraph(notes)

	js, err := json.Marshal(g)

	if err != nil {
		fmt.Println(err.Error())
		panic("fuggg")
	}

	fmt.Fprintf(w, string(js))
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

func setJSONResponse(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		f(w, r)
	}
}

func main() {
	initLogging()

	slog.Info(getEnvConfig().notionApiKey)

	slog.Info("Starting server on port 8080")
	http.HandleFunc("/", logging(welcome))
	http.HandleFunc("/database", setJSONResponse(logging(database)))
	http.HandleFunc("/graph", setJSONResponse(logging(getGraph)))

	http.ListenAndServe(":8080", nil)
}

func getAllPagesAsNotionNotes() []NotionNote {
	pages := getPagesFromDB(getEnvConfig().notionDbId)

	notionNotes := []NotionNote{}

	for _, page := range pages {
		notionNotes = append(notionNotes, parsePageToNotionNote(page))
	}

	return notionNotes
}

func buildGraph(notes []NotionNote) Graph {
	graph := Graph{
		G: make(map[NodeId]Node),
	}

	for _, note := range notes {
		graph.addNode(note)
	}

	return graph
}

func initLogging() {
	slog.SetDefault(slog.New(
		tint.NewHandler(os.Stderr, &tint.Options{
			Level:      slog.LevelDebug,
			TimeFormat: time.Kitchen,
		}),
	))
}
