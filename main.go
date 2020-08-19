package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gomodule/redigo/redis"
	"github.com/smeruelo/glow/graph"
	"github.com/smeruelo/glow/graph/generated"
)

func main() {
	dbHost, ok := os.LookupEnv("DB_HOST")
	if !ok {
		log.Fatal("Environment variable DB_HOST not found")
	}

	dbPort, ok := os.LookupEnv("DB_PORT")
	if !ok {
		log.Fatal("Environment variable DB_PORT not found")
	}

	db, err := redis.Dial("tcp", dbHost+":"+dbPort)
	if err != nil {
		log.Printf("Unable to connect to database: %s", err)
	}
	defer db.Close()

	graphqlServer := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: graph.NewResolver(db),
	}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", graphqlServer)

	log.Fatal(http.ListenAndServe(":80", nil))
}
