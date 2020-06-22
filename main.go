package main
import (
	"fmt"
	"log"
	"net/http"
	"./postgres"
	"./gql"
	"./server"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/graphql-go/graphql"
)

func main() {
	router, db := initializeAPI()
	defer db.Close()

	log.Fatal(http.ListenAndServe(":4000",router))
}

// sourced from https://medium.com/@bradford_hamilton/building-an-api-with-graphql-and-go-9350df5c9356
func initializeAPI() (*chi.Mux, *postgres.Db) {
	// Create a new router
	router := chi.NewRouter()

	// Create a new connection to our pg database
	db, err := postgres.New()

	if err != nil {
		log.Fatal(err)
	}

	// Create our root query for graphql
	rootQuery := gql.NewRoot(db)

	//rootMutation := gql.NewMutation(db)

	// Create a new graphql schema, passing in the the root query
	sc, err := graphql.NewSchema(graphql.SchemaConfig{
			Query: rootQuery.Query,
			Mutation: rootQuery.Mutation,
	})

	if err != nil {
		fmt.Println("Error creating schema: ", err)
	}

	// Create a server struct that holds a pointer to our database as well
	// as the address of our graphql schema
	s := server.Server{
		GqlSchema: &sc,
	}

	// Add some middleware to our router
	router.Use(
		render.SetContentType(render.ContentTypeJSON), // set content-type headers as application/json
		middleware.Logger,          // log api request calls
		middleware.Compress(5),
		middleware.StripSlashes,    // match paths with a trailing slash, strip it, and continue routing through the mux
		middleware.Recoverer,       // recover from panics without crashing server
	)

	// Create the graphql route with a Server method to handle it
	router.Post("/graphql", s.GraphQL())

	return router, db
}