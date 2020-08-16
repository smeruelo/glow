package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gomodule/redigo/redis"
)

type Project struct {
	Name     string `json:"Name"`
	Goal     int    `json:"Goal"`
	Achieved int    `json:"Achieved"`
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

	c, err := redis.Dial("tcp", dbHost+":"+dbPort)
	if err != nil {
		log.Printf("Unable to connect to database: %s", err)
	}
	defer c.Close()

	http.HandleFunc("/projects", func(w http.ResponseWriter, r *http.Request) {
		projectsJSON, err := redis.StringMap(c.Do("HGETALL", "projects"))
		if err != nil {
			log.Printf("Database error: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		projects := make([]Project, len(projectsJSON))
		i := 0
		for _, pJSON := range projectsJSON {
			var p Project
			if err = json.Unmarshal([]byte(pJSON), &p); err != nil {
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

	log.Fatal(http.ListenAndServe(":80", nil))
}
