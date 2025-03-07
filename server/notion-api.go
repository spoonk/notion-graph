package main

import (
	"context"
	// "fmt"
	"log/slog"

	"github.com/jomei/notionapi"
)

func doSomething() {
	ctx := context.Background()
	dbid := notionapi.DatabaseID(getEnvConfig().notionDbId)
	client := notionapi.NewClient(notionapi.Token(getEnvConfig().notionApiKey))

	qr, err := client.Database.Query(ctx, dbid, &notionapi.DatabaseQueryRequest{})
	if err != nil {
		slog.Error(err.Error())
	}

	page, ok := qr.Results[0].Properties["Sub-item"].(*notionapi.RelationProperty)
	if !ok {

		slog.Info(getJSON(qr.Results[0].Properties["Sub-item"]))
		slog.Info(getJSON(page))
		slog.Info("fuggg")
		return
	}
	dumb := notionapi.RelationProperty{
		ID:       notionapi.ObjectID(page.GetID()),
		Type:     page.GetType(),
		Relation: page.Relation,
	}

	slog.Info(getJSON(dumb))
}
