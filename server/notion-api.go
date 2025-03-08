package main

import (
	"context"
	"github.com/jomei/notionapi"
	"log/slog"
)

type NotionNote struct {
	ID         string   `json:"id"`
	Title      string   `json:"title"`
	RelatedIds []string `json:"relatedIds"`
}

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

func getPagesFromDB(databaseIdString string) []notionapi.Page {
	dbid := notionapi.DatabaseID(databaseIdString)

	ctx := context.Background()
	client := notionapi.NewClient(notionapi.Token(getEnvConfig().notionApiKey))

	qr, err := client.Database.Query(ctx, dbid, &notionapi.DatabaseQueryRequest{})
	if err != nil {
		slog.Error(err.Error())
		return []notionapi.Page{}
	}

	return qr.Results
}

func parsePageToNotionNote(page notionapi.Page) NotionNote {

	title, titleOk := page.Properties["Name"].(*notionapi.TitleProperty)

	if !titleOk {
		slog.Info(getJSON(page))
		panic("raaa") // tee hee
	}

	related, ok := page.Properties["Sub-item"].(*notionapi.RelationProperty)

	if !ok {
		slog.Info(getJSON(page))
		panic("parsing related failed") // tee hee
	}

	relatedIDs := []string{}
	for _, relation := range related.Relation {
		relatedIDs = append(relatedIDs, relation.ID.String())
	}

	return NotionNote{
		ID:         string(page.ID),
		Title:      title.Title[0].PlainText,
		RelatedIds: relatedIDs,
	}

}
