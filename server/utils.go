package main

import (
	json "github.com/neilotoole/jsoncolor"
	"log/slog"
)

func getJSON(obj any) string {
	encoded, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		slog.Info(err.Error())
	}

	return string(encoded)
}
