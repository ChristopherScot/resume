package handlers

import (
	"context"
	"encoding/json"
	"io"
	"log/slog"

	"github.com/ChristopherScot/resume/db"
	"github.com/ChristopherScot/resume/models"
	"github.com/ChristopherScot/resume/restapi/operations/open"
	"github.com/davecgh/go-spew/spew"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/google/uuid"
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
	resume, err := db.GetResume(context.Background(), params.ID.String())
	if err != nil {
		slog.Error("resume not found", "id", params.ID.String())
		return open.NewGetResumeNotFound().WithPayload(&models.SimpleErrorResponse{Error: swag.String("resume not found. id: " + params.ID.String())})
	}

	resp := open.NewGetResumeOK().WithPayload(&resume)
	return resp
}

func CreateResume(params open.CreateResumeParams) middleware.Responder {
	resume := models.Resume{}
	decoder := json.NewDecoder(params.HTTPRequest.Body)
	err := decoder.Decode(&resume)
	if err != nil && err != io.EOF {
		slog.Error("error decoding request body", "error", err)
		return open.NewCreateResumeBadRequest().WithPayload(&models.SimpleErrorResponse{Error: swag.String(err.Error())})
	}

	slog.Info("Resume Content", "resume", spew.Sdump(resume))

	uuid := uuid.New()
	uuidStr := strfmt.UUID(uuid.String())
	err = db.CreateResume(context.Background(), uuidStr.String(), resume)
	// err = db.CreateBasicThing(context.Background(), uuidStr.String(), jsonResume)
	if err != nil {
		slog.Error("error creating resume", "error", err)
		return open.NewCreateResumeInternalServerError().WithPayload(&models.SimpleErrorResponse{Error: swag.String(err.Error())})

	}

	slog.Debug("created resume", "id", uuid.String())

	resp := open.NewCreateResumeCreated().WithPayload(&models.UUIDResponse{ID: &uuidStr})

	return resp
}
