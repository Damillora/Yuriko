package database

import (
	"log"

	"github.com/Damillora/Yuriko/pkg/config"
	"github.com/typesense/typesense-go/typesense"
	"github.com/typesense/typesense-go/typesense/api"
)

var Client *typesense.Client

func Initialize(config config.Config) {
	Client = typesense.NewClient(
		typesense.WithServer(config.TypesenseApiUrl),
		typesense.WithAPIKey(config.TypesenseApiKey))

	CreateDatabaseSchema()
}

func CreateDatabaseSchema() {
	schema := &api.CollectionSchema{
		Name: "articles",
		Fields: []api.Field{
			{
				Name: "id",
				Type: "string",
			},
			{
				Name: "url",
				Type: "string",
			},
			{
				Name: "slug",
				Type: "string",
			},
			{
				Name: "title",
				Type: "string",
			},
			{
				Name: "excerpt",
				Type: "string",
			},
			{
				Name: "html",
				Type: "string",
			},
			{
				Name: "primary_tag",
				Type: "string",
			},
			{
				Name: "published_at",
				Type: "int64",
			},
		},
		DefaultSortingField: "published_at",
	}
	_, err := Client.Collections().Create(schema)
	if err != nil {
		log.Println(err.Error())
	}
}
