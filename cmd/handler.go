package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log/slog"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"github.com/ChristopherScot/resume/pkg/models"
)

func getJSONResumeFromBody(body io.ReadCloser) (models.Resume, error) {
	// decode the body of the request into a models.Resume struct
	var jsonResume models.Resume

	err := json.NewDecoder(body).Decode(&jsonResume)
	if err != nil {
		return models.Resume{}, err
	}
	if err := jsonResume.Validate(); err != nil {
		return models.Resume{}, err
	}
	return jsonResume, nil
}

// handler functions for the endpoints defined in main()

func createResume(w http.ResponseWriter, r *http.Request) {
	// create a new resume

	jsonResume, err := getJSONResumeFromBody(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("error decoding request body: %s", err), http.StatusBadRequest)
		slog.Error("error decoding request body", "error", err)
		return
	}

	// generate a uuid for the resume
	// and then add the resume to the map of resumes

	uuid := uuid.New().String()

	Resumes[uuid] = jsonResume

	slog.Default().Debug("created resume", "id", uuid)

	// return the id of the new resume
	w.Write([]byte(uuid))

}

func getResume(w http.ResponseWriter, r *http.Request) {
	Resume, ok := Resumes[mux.Vars(r)["id"]]
	if !ok {
		http.Error(w, "resume not found", http.StatusNotFound)
		return
	}

	// return the resume
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Resume)
}

func updateResume(w http.ResponseWriter, r *http.Request) {
	// update a resume by id
	jsonResume, err := getJSONResumeFromBody(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("error decoding request body: %s", err), http.StatusBadRequest)
		slog.Error("error decoding request body", "error", err)
		return
	}

	// This step is definitely not neccesary here but it's a good practice
	// that I don't want to skip over when I use a real database
	if _, ok := Resumes[mux.Vars(r)["id"]]; !ok {
		http.Error(w, "resume not found", http.StatusNotFound)
		return
	}

	Resumes[mux.Vars(r)["id"]] = jsonResume

}

func deleteResume(w http.ResponseWriter, r *http.Request) {
	// delete a resume by id
	if _, ok := Resumes[mux.Vars(r)["id"]]; !ok {
		http.Error(w, "resume not found", http.StatusNotFound)
		return
	}

	delete(Resumes, mux.Vars(r)["id"])
	slog.Default().Debug("deleted resume", "id", mux.Vars(r)["id"])
}

func viewResume(w http.ResponseWriter, r *http.Request) {

	desiredTemplate := "default"
	if r.URL.Query().Get("template") != "" {
		desiredTemplate = r.URL.Query().Get("template")
	}

	// get the resume by id
	resume, ok := Resumes[mux.Vars(r)["id"]]
	if !ok {
		http.Error(w, "resume not found", http.StatusNotFound)
		return
	}

	tmpl := template.Must(template.ParseFiles(fmt.Sprintf("../internal/templates/%s.template", desiredTemplate)))

	// Let's not go straight to the writer but instead use a buffer
	// to capture any errors that might occur during template execution
	// and then write the buffer to the writer
	// This way we can handle the error and return a 500 status code
	buffer := new(bytes.Buffer)
	err := tmpl.Execute(buffer, resume)
	if err != nil {
		http.Error(w, fmt.Sprintf("error rendering template: %s", err), http.StatusInternalServerError)
		return
	}
	// write the buffer to the writer
	w.Write(buffer.Bytes())

}
