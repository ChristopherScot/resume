package main

import (
	"net/http"

	"github.com/ChristopherScot/resume/pkg/models"
	"github.com/gorilla/mux"
)

// create a gorilla/mux router with endpoints for
// creating, reading, updating, and deleting resumes
// using the models.Resume struct
// as well as a view endpoint for rendering the resume with a template

// instead of using a database for now we'll just use a map of resumes
var Resumes = make(map[string]models.Resume)

//go:generate swagger generate spec -o $PWD/../swagger.yaml
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/resume", createResume).Methods("POST")
	r.HandleFunc("/resume/{id}", getResume).Methods("GET")
	r.HandleFunc("/resume/{id}", updateResume).Methods("PUT")
	r.HandleFunc("/resume/{id}", deleteResume).Methods("DELETE")
	r.HandleFunc("/resume/{id}/view", viewResume).Methods("GET")
	// throw a basic error if the route is not found
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "And I still haven't found... what I'm looking for.", http.StatusNotFound)
	})
	http.ListenAndServe(":2020", r)

}
