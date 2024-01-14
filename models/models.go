package models

import (
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator/v10"
)

// Resume is a struct that represents a resume in JSON format.
type Resume struct {
	Basics struct {
		Name     string `json:"name" validate:"required"`
		Label    string `json:"label"`
		Image    string `json:"image"`
		Email    string `json:"email" validate:"required,email"`
		Phone    string `json:"phone" validate:"required"`
		URL      string `json:"url"`
		Summary  string `json:"summary" validate:"required"`
		Location struct {
			Address     string `json:"address"`
			PostalCode  string `json:"postalCode"`
			City        string `json:"city"`
			CountryCode string `json:"countryCode"`
			Region      string `json:"region"`
		} `json:"location"`
		Profiles []struct {
			Network  string `json:"network"`
			Username string `json:"username"`
			URL      string `json:"url"`
		} `json:"profiles"`
	} `json:"basics"`
	Work []struct {
		Name       string   `json:"name" validate:"required"`
		Position   string   `json:"position" validate:"required"`
		URL        string   `json:"url"`
		StartDate  string   `json:"startDate" validate:"required"`
		EndDate    string   `json:"endDate" validate:"required"`
		Summary    string   `json:"summary" validate:"required"`
		Highlights []string `json:"highlights"`
	} `json:"work"`
	Volunteer []struct {
		Organization string   `json:"organization"`
		Position     string   `json:"position"`
		URL          string   `json:"url"`
		StartDate    string   `json:"startDate"`
		EndDate      string   `json:"endDate"`
		Summary      string   `json:"summary"`
		Highlights   []string `json:"highlights"`
	} `json:"volunteer"`
	Education []struct {
		Institution string   `json:"institution" `
		URL         string   `json:"url"`
		Area        string   `json:"area"`
		StudyType   string   `json:"studyType"`
		StartDate   string   `json:"startDate"`
		EndDate     string   `json:"endDate"`
		Score       string   `json:"score"`
		Courses     []string `json:"courses"`
	} `json:"education"`
	Awards []struct {
		Title   string `json:"title"`
		Date    string `json:"date"`
		Awarder string `json:"awarder"`
		Summary string `json:"summary"`
	} `json:"awards"`
	Certificates []struct {
		Name   string `json:"name"`
		Date   string `json:"date"`
		Issuer string `json:"issuer"`
		URL    string `json:"url"`
	} `json:"certificates"`
	Publications []struct {
		Name        string `json:"name"`
		Publisher   string `json:"publisher"`
		ReleaseDate string `json:"releaseDate"`
		URL         string `json:"url"`
		Summary     string `json:"summary"`
	} `json:"publications"`
	Skills []struct {
		Name     string   `json:"name" validate:"required"`
		Level    string   `json:"level" validate:"required"`
		Keywords []string `json:"keywords" validate:"required"`
	} `json:"skills"`
	Languages []struct {
		Language string `json:"language"`
		Fluency  string `json:"fluency"`
	} `json:"languages"`
	Interests []struct {
		Name     string   `json:"name"`
		Keywords []string `json:"keywords"`
	} `json:"interests"`
	References []struct {
		Name      string `json:"name"`
		Reference string `json:"reference"`
	} `json:"references"`
	Projects []struct {
		Name        string   `json:"name"`
		StartDate   string   `json:"startDate"`
		EndDate     string   `json:"endDate"`
		Description string   `json:"description"`
		Highlights  []string `json:"highlights"`
		URL         string   `json:"url"`
	} `json:"projects"`
}

// CreateResumeStructFromJSON validates a resume against the JSON resume schema.
func CreateResumeStructFromJSON(jsonResume []byte) (Resume, error) {
	var resume Resume
	err := json.Unmarshal(jsonResume, &resume)
	if err != nil {
		return Resume{}, fmt.Errorf("Unable to marshal JSON due to %s", err)
	}
	err = validator.New().Struct(resume)
	if err != nil {
		return Resume{}, fmt.Errorf("Validation Errors Found %s", err)
	}
	return resume, nil
}

// String returns a string representation of a resume.
func (r Resume) String() string {
	return "TODO"
}
