package gql

import (
	"fmt"
	"github.com/graphql-go/graphql"
)

// sourced from https://medium.com/@bradford_hamilton/building-an-api-with-graphql-and-go-9350df5c9356
// ExecuteQuery runs our graphql queries
func ExecuteQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})

	// Error check
	if len(result.Errors) > 0 {
		fmt.Printf("Unexpected errors inside ExecuteQuery: %v", result.Errors)
	}

	return result
}

func ExecuteMutation(query string, schema graphql.Schema) *graphql.Result{
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})

	// Error check
	if len(result.Errors) > 0 {
		fmt.Printf("Unexpected errors inside ExecuteMutation: %v", result.Errors)
	}

	return result
}