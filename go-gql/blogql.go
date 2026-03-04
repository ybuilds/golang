package main

import (
	"log"

	"github.com/graphql-go/graphql"
)

func CreateBlogType() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Blog",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"title": &graphql.Field{
					Type: graphql.String,
				},
				"content": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)
}

func QueryType(blog *graphql.Object) *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"blogs": &graphql.Field{
					Type: graphql.NewList(blog),
					Args: graphql.FieldConfigArgument{
						"limit": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
						"offset": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
					},
					Resolve: func(p graphql.ResolveParams) (any, error) {
						limit, _ := p.Args["limit"].(int)
						if limit <= 0 || limit > 20 {
							limit = 10
						}

						offset, _ := p.Args["offset"].(int)
						if offset < 0 {
							offset = 0
						}

						var blogs []Blog

						rows, err := DB.Query("SELECT id, title, content FROM blogs LIMIT $1 OFFSET $2", limit, offset)
						if err != nil {
							log.Println("error fetching blogs from db")
							return nil, err
						}
						defer rows.Close()

						for rows.Next() {
							var blog Blog
							err := rows.Scan(&blog.Id, &blog.Title, &blog.Content)
							if err != nil {
								log.Println("error scanning blog from db")
								return nil, err
							}
							blogs = append(blogs, blog)
						}

						return blogs, nil
					},
				},
			},
		},
	)
}
