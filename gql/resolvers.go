package gql

import (
	"../postgres"
	"github.com/graphql-go/graphql"
)


// sourced from https://medium.com/@bradford_hamilton/building-an-api-with-graphql-and-go-9350df5c9356
// Resolver struct holds a connection to our database
type Resolver struct {
	db *postgres.Db
}

// GetFeedbackResolver resolves our feedback query through a db call to GetFeedbacks
func (r *Resolver) GetFeedbacksResolver(p graphql.ResolveParams) (interface{}, error) {
	users := r.db.GetFeedbacks()
	return users, nil
}

// CreateFeedbackResolver resolves our feedback creation through a db call to CreateFeedback
func (r *Resolver) CreateFeedbackResolver(p graphql.ResolveParams) (interface{}, error) {
	rating, ok := p.Args["rating"].(int)
	email, ok := p.Args["email"].(string)
	details, ok := p.Args["details"].(string)

	if ok{
		user := r.db.CreateFeedback(rating, email, details)
		return user, nil
	}
	return nil, nil 
}