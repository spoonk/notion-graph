package main

import (
	"context"
	"github.com/jomei/notionapi"
	"log/slog"
)

func doSomething() {
	ctx := context.Background()
	dbid := notionapi.DatabaseID(getEnvConfig().notionDbId)
	client := notionapi.NewClient(notionapi.Token(getEnvConfig().notionApiKey))
	// db, err := client.Database.Get(ctx, dbid)
	// if err != nil {
	// 	slog.Error(err.Error())
	// 	return
	// }

	qr, err := client.Database.Query(ctx, dbid, &notionapi.DatabaseQueryRequest{})
	if err != nil {
		slog.Error(err.Error())
	}

	// slog.Info(getJSON(qr.Results[0]))

	page := qr.Results[0]

	// slog.Info(getJSON(page.Properties["Name"].GetType()))

	name := page.Properties["Name"]

	slog.Info(getJSON(name))

	nameProperty, ok := name.(notionapi.TitleProperty)
	if !ok {
		slog.Info("name property failed")
		return
	}

	slog.Info(getJSON(nameProperty.Title))
}
