package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gomodule/redigo/redis"
	"github.com/google/uuid"
	"github.com/smeruelo/glow/graph"
	"github.com/smeruelo/glow/graph/generated"
)

type Project struct {
	Name     string `json:"Name"`
	Goal     int    `json:"Goal"`
	Achieved int    `json:"Achieved"`
}

func fetchProjectByID(c redis.Conn, uuid uuid.UUID) (Project, bool) {
	pJSON, err := redis.Bytes(c.Do("HGET", "projects", uuid.String()))
	if err != nil {
		return Project{}, false
	}
	var p Project
	if err := json.Unmarshal(pJSON, &p); err != nil {
		log.Printf("Unable to unmarshal project: %s", err)
		return Project{}, false
	}
	return p, true
}

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

	http.HandleFunc("/projects", func(w http.ResponseWriter, r *http.Request) {
		projectsJSON, err := redis.StringMap(db.Do("HGETALL", "projects"))
		if err != nil {
			log.Printf("Database error: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		projects := make([]Project, len(projectsJSON))
		i := 0
		for _, pJSON := range projectsJSON {
			var p Project
			if err := json.Unmarshal([]byte(pJSON), &p); err != nil {
				log.Printf("Unable to unmarshal project: %s", err)
			} else {
				projects[i] = p
				i++
			}
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		output, err := json.Marshal(projects)
		if err != nil {
			log.Printf("Unable to marshal projects: %s", err)
		}
		w.Write(output)
	})

	graphqlServer := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: graph.NewResolver(db),
	}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", graphqlServer)

	log.Fatal(http.ListenAndServe(":80", nil))
}
