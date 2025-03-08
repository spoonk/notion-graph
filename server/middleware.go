package main

import (
	"log/slog"
	"net/http"
)

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

func installMiddleware(f http.HandlerFunc) http.HandlerFunc {
	slog.Info("abcdef???")
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("hellooo???")
		fn := setJSONResponse(logging(f))
		fn(w, r)
	}

