package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"github.com/vncnttan/TrainingGraphQL/database"
	"github.com/vncnttan/TrainingGraphQL/graph"
	"github.com/vncnttan/TrainingGraphQL/middleware"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	database.MigrateTable()

	c := graph.Config{Resolvers: &graph.Resolver{
		DB: database.GetInstance(),
	}}

	c.Directives.Auth = func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
		// Harusnya disini ambil token
		token := ctx.Value("TokenValue")

		if token == nil {
			// Harusnya disini decode token
			return nil, &gqlerror.Error{
				Message: "Please Login",
			}
		}

		ctx = context.WithValue(ctx, "UserID", token)
		return next(ctx)
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(c))

	router := mux.NewRouter()

	router.Use(middleware.AuthMiddleware)
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
