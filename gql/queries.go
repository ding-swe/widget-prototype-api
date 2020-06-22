package gql 

import (
	"../postgres"
	"github.com/graphql-go/graphql"
)

type RootQuery struct {
	Query *graphql.Object
	Mutation *graphql.Object
}

// sourced from https://medium.com/@bradford_hamilton/building-an-api-with-graphql-and-go-9350df5c9356
func NewRoot(db *postgres.Db) *RootQuery {
	// Create a resolver holding our databse. Resolver can be found in resolvers.go
	resolver := Resolver{db: db}

	// Create a new Root that describes our base query set up. In this
	// example we have a user query that takes one argument called name
	root := RootQuery{
		Query: graphql.NewObject(
			graphql.ObjectConfig{
				Name: "Query",
				Fields: graphql.Fields{
					"feedback": &graphql.Field{
						// Slice of User type which can be found in types.go
						Type: graphql.NewList(feedbackType),
						Resolve: resolver.GetFeedbacksResolver,
					},
				},
			},
		),
		Mutation: graphql.NewObject(
			graphql.ObjectConfig{
				Name: "Mutation",
				Fields: graphql.Fields{
					"createFeedback": &graphql.Field{
						// Slice of User type which can be found in types.go
						Type: feedbackType,
						Resolve: resolver.CreateFeedbackResolver,
					},
				},
			},
		),
	}
	return &root
}