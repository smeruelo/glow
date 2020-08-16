package main

import (
	"encoding/json"
	"log"
	"net/http"
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
