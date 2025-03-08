package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/joho/godotenv/autoload" // underscore means import solely for side effects
	"github.com/lmittmann/tint"
	"log/slog"
	"net/http"
	"os"
	"time"
)

func getGraph(w http.ResponseWriter, r *http.Request) {
	notes := getAllPagesAsNotionNotes()
	g := buildGraph(notes)

	js, err := json.Marshal(g)

	if err != nil {
		fmt.Println(err.Error())
		panic("fuggg")
	}

	fmt.Fprintf(w, string(js)) // write to http response writer
}

func main() {
	initLogging()
	slog.Info("Starting server on port 8080")

	// http.HandleFunc("/graph", installMiddleware(getGraph))
	http.HandleFunc("/graph", installMiddleware(getGraph))

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
