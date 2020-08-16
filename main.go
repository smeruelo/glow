package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gomodule/redigo/redis"
)

type Project struct {
	Name     string
	Goal     int
	Achieved int
}

var projects = []Project{
	{Name: "First", Goal: 3000, Achieved: 1030},
	{Name: "Second", Goal: 700, Achieved: 0},
	{Name: "Third", Goal: 5550, Achieved: 2200},
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
		b, err := json.Marshal(projects)
		if err != nil {
			log.Printf("Unable to marshal projects: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Write(b)
	})

	log.Fatal(http.ListenAndServe(":80", nil))
}
