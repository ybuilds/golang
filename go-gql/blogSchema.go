package main

import (
	"log"

	"github.com/graphql-go/graphql"
)

func CreateBlogSchema() (*graphql.Schema, error) {
	blogType := CreateBlogType()

	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query: QueryType(blogType),
		},
	)
	if err != nil {
		log.Println("error creating new schema for blog type")
		return nil, err
	}

	return &schema, nil
}
