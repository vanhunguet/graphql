package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/vanhunguet/GrabhQL/graph"
	"github.com/vanhunguet/GrabhQL/graph/generated"
	"github.com/vanhunguet/GrabhQL/internal/app/adapter/database/db/mysql"
	"log"
	"net/http"
	"os"
)

const defaultPort = "8080"

func main() {
	mysql.SetupDatabaseConnection()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
