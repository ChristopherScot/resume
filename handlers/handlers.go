package handlers

import (
	"context"
	"log/slog"
	"time"

	"github.com/ChristopherScot/resume/db"
	"github.com/ChristopherScot/resume/models"
	"github.com/ChristopherScot/resume/restapi/operations/open"
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
		slog.Warn("resume not found", "id", params.ID.String())
		return open.NewGetResumeNotFound().WithPayload(&models.SimpleErrorResponse{Error: swag.String("resume not found. id: " + params.ID.String())})
	}

	slog.Info("retrieved resume", "id", params.ID.String(), "resume", resume)

	resp := open.NewGetResumeOK().WithPayload(&resume)
	return resp
}

func UpdateResume(params open.UpdateResumeParams) middleware.Responder {

	metadata := &models.ResumeMetadata{
		Updated: strfmt.DateTime(time.Now()),
	}
	err := db.UpdateResume(context.Background(), params.ID.String(), *params.Resume, *metadata)
	if err != nil {
		slog.Error("error updating resume", "error", err, "id", params.ID.String(), "resume", params.Resume, "metadata", metadata)
		return open.NewUpdateResumeInternalServerError().WithPayload(&models.SimpleErrorResponse{Error: swag.String(err.Error())})
	}

	slog.Debug("updated resume", "id", params.ID.String(), "resume", params.Resume)

	resp := open.NewUpdateResumeOK().WithPayload(&models.UUIDResponse{ID: &params.ID})

	return resp
}

func DeleteResume(params open.DeleteResumeParams) middleware.Responder {

	resume, err := db.GetResume(context.Background(), params.ID.String())
	if err != nil {
		slog.Error("error when finding resume", "id", params.ID.String())
		if resume.Basics == nil {
			slog.Warn("resume not found", "id", params.ID.String())
			return open.NewDeleteResumeNotFound().WithPayload(&models.SimpleErrorResponse{Error: swag.String("resume not found. id: " + params.ID.String())})
		}
		return open.NewDeleteResumeNotFound().WithPayload(&models.SimpleErrorResponse{Error: swag.String("resume not found. id: " + params.ID.String())})
	}

	err = db.DeleteResume(context.Background(), params.ID.String())
	if err != nil {
		slog.Error("error deleting resume", "error", err)
		return open.NewDeleteResumeInternalServerError().WithPayload(&models.SimpleErrorResponse{Error: swag.String(err.Error())})
	}

	slog.Debug("deleted resume", "id", params.ID.String())

	resp := open.NewDeleteResumeOK().WithPayload(&models.UUIDResponse{ID: &params.ID})

	return resp
}

func CreateResume(params open.CreateResumeParams) middleware.Responder {

	uuid := uuid.New()
	uuidStr := strfmt.UUID(uuid.String())
	metadata := &models.ResumeMetadata{
		Created: strfmt.DateTime(time.Now()),
		Updated: strfmt.DateTime(time.Now()),
	}
	err := db.CreateResume(context.Background(), uuidStr.String(), *params.Resume, *metadata)
	if err != nil {
		slog.Error("error creating resume", "error", err)
		return open.NewCreateResumeInternalServerError().WithPayload(&models.SimpleErrorResponse{Error: swag.String(err.Error())})

	}

	slog.Debug("created resume", "id", uuid.String(), "resume", params.Resume)

	resp := open.NewCreateResumeCreated().WithPayload(&models.UUIDResponse{ID: &uuidStr})

	return resp
}
