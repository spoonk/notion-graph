package main

import (
	"fmt"
	"github.com/lmittmann/tint"
	"log/slog"
	"net/http"
	"os"
	"time"
)

func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my website!")
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

	slog.Info("Starting server on port 8080")
	http.HandleFunc("/", logging(welcome))

	http.ListenAndServe(":8080", nil)
}

func initLogging() {
	slog.SetDefault(slog.New(
		tint.NewHandler(os.Stderr, &tint.Options{
			Level:      slog.LevelDebug,
			TimeFormat: time.Kitchen,
		}),
	))
}
