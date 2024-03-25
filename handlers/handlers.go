package handlers

import (
	"github.com/ChristopherScot/resume/models"
	"github.com/ChristopherScot/resume/restapi/operations/open"
	"github.com/go-openapi/runtime/middleware"
)

func GetApiIdentifier(params open.GetAPIIdentifierParams) middleware.Responder {
	str := "go-swagger Lambda integration API - version 1.0"

	r := models.SimpleMessageResponse{
		Message: &str,
	}

	resp := open.NewGetAPIIdentifierOK().WithPayload(&r)

	return resp
}

func GetResume(params open.GetResumeParams) middleware.Responder {
	resume := models.Resume{
		Basics: &models.ResumeBasics{
			Name:  "John Doe",
			Label: "Programmer",
			Image: "https://john.doe.com/picture.jpg",
		},
	}

	resp := open.NewGetResumeOK().WithPayload(&resume)
	return resp
}
