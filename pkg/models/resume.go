package models

import (
	"strings"
	"time"
)

const timeLayout = "2006-01-02"

type jsonTime struct {
	time.Time
}

// Validate validates the resume and returns a list of problems it found if any
// TODO: Make this do actual validations
func (r Resume) Validate() error {
	return nil
}

func (ts *jsonTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")

	t, err := time.Parse(timeLayout, s)
	if err != nil {
		return err
	}

	ts.Time = t
	return nil
}

type Interval struct {
	StartDate jsonTime `json:"startDate" validate:"ltfield=EndDate"`
	EndDate   jsonTime `json:"endDate" validate:"gtfield=StartDate"`
}

// Resume holds the information according to schema
// described in https://jsonresume.org/schema/
type Resume struct {
	Basics       Basics        `json:"basics"`
	Work         []Work        `json:"work"`
	Volunteer    []Volunteer   `json:"volunteer"`
	Education    []Education   `json:"education"`
	Awards       []Award       `json:"awards"`
	Publications []Publication `json:"publications"`
	Skills       []Skill       `json:"skills"`
	Languages    []Language    `json:"languages"`
	Interests    []Interest    `json:"interests"`
	References   []Reference   `json:"references"`
}

type Location struct {
	Address     string `json:"address"`
	PostalCode  string `json:"postalCode"`
	City        string `json:"city"`
	CountryCode string `json:"countryCode"`
	Region      string `json:"region"`
}

type Profile struct {
	Network  string `json:"network"`
	Username string `json:"username"`
	URL      string `json:"url" validate:"url"`
}

type Basics struct {
	Name     string    `json:"name" validate:"required"`
	Label    string    `json:"label" validate:"required"`
	Picture  string    `json:"picture"`
	Email    string    `json:"email" validate:"required,email"`
	Phone    string    `json:"phone"`
	Website  string    `json:"website" validate:"url"`
	Summary  string    `json:"summary"`
	Location Location  `json:"location"`
	Profiles []Profile `json:"profiles"`
}

type Work struct {
	Interval
	Name       string   `json:"name"`
	Position   string   `json:"position"`
	Website    string   `json:"website" validate:"url"`
	Summary    string   `json:"summary"`
	Highlights []string `json:"highlights"`
}

type Volunteer struct {
	Interval
	Organization string   `json:"organization"`
	Position     string   `json:"position"`
	Website      string   `json:"website" validate:"url"`
	Summary      string   `json:"summary"`
	Highlights   []string `json:"highlights"`
}

type Education struct {
	Interval
	Institution string   `json:"institution"`
	Area        string   `json:"area"`
	StudyType   string   `json:"studyType"`
	Gpa         string   `json:"gpa"`
	Courses     []string `json:"courses"`
}

type Award struct {
	Title   string   `json:"title"`
	Date    jsonTime `json:"date" validate:"lte"`
	Awarder string   `json:"awarder"`
	Summary string   `json:"summary"`
}

type Publication struct {
	Name        string   `json:"name"`
	Publisher   string   `json:"publisher"`
	ReleaseDate jsonTime `json:"releaseDate" validate:"lte"`
	Website     string   `json:"website" validate:"url"`
	Summary     string   `json:"summary"`
}

type Skill struct {
	Name     string   `json:"name"`
	Level    string   `json:"level"`
	Keywords []string `json:"keywords"`
}

type Language struct {
	Language string `json:"language"`
	Fluency  string `json:"fluency"`
}

type Interest struct {
	Name     string   `json:"name"`
	Keywords []string `json:"keywords"`
}

type Reference struct {
	Name      string `json:"name"`
	Reference string `json:"reference"`
}
