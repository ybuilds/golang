package main

import (
	"log"

	"github.com/graphql-go/handler"
)

func QueryBlogsHandler() (*handler.Handler, error) {
	blogSchema, err := CreateBlogSchema()
	if err != nil {
		log.Println("error fetching blog schema")
		return nil, err
	}

	handler := handler.New(&handler.Config{
		Schema: blogSchema,
		Pretty: true,
	})

	return handler, nil
}
