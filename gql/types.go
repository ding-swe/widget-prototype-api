package gql

import (
	"github.com/graphql-go/graphql"
)

// "fmt"
// pg "../postgres"


// sourced from https://github.com/sohelamin/graphql-postgres-go/blob/master/main.go
var feedbackType = graphql.NewObject(graphql.ObjectConfig{
	Name:	"Feedback",
	Description: "User Feedback",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.Int),
			Description: "The identifier of the feedback.",
		},
		"email": &graphql.Field{
			Type:       graphql.String,
			Description: "The email related to the feedback.",
		},
		"rating": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.Int),
			Description: "The rating",
		},
		"details": &graphql.Field{
			Type:        graphql.String,
			Description: "The detail text of the feedback.",
		},
	},
})
