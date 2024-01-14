package main

import (
	"html/template"
	"net/http"

	"github.com/ChristopherScot/resume/models"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	ResumesByUser := make(map[string]models.Resume)
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	// A put request to resume will create a resume object from a json resume
	// which is passed in the request body.
	// The json resume is validated against the json resume schema.
	r.Post("/resume/{user}", func(w http.ResponseWriter, r *http.Request) {
		// Create a resume object from the json resume.
		input := make([]byte, r.ContentLength)
		_, err := r.Body.Read(input)
		resume, err := models.CreateResumeStructFromJSON(input)
		if err != nil {
			w.Write([]byte(err.Error()))
		}
		// Add the resume to the map of resumes.
		ResumesByUser[chi.URLParam(r, "user")] = resume
		// respond with just a 200 status code.
		w.WriteHeader(http.StatusOK)
	})

	r.Get("/resume/{user}", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("resumeTemplateExample.go.template"))

		// if the user has a resume, return it.
		if resume, ok := ResumesByUser[chi.URLParam(r, "user")]; ok {
			//use go templates to make an actual html resume page

			err := tmpl.Execute(w, resume)
			if err != nil {
				w.Write([]byte(err.Error()))
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	})

	http.ListenAndServe(":3000", r)
}
