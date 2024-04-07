package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"

	"github.com/ChristopherScot/resume/db"
	"github.com/ChristopherScot/resume/handlers"
	"github.com/ChristopherScot/resume/lib"
	"github.com/ChristopherScot/resume/restapi"
	"github.com/ChristopherScot/resume/restapi/operations"
	"github.com/ChristopherScot/resume/restapi/operations/open"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	loads "github.com/go-openapi/loads"
)

var httpAdapter *httpadapter.HandlerAdapter
var nullHandler = false

func setupHandlers() *operations.ResumeAPI {

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		panic(err)
	}

	api := operations.NewResumeAPI(swaggerSpec)
	api.OpenGetAPIIdentifierHandler = open.GetAPIIdentifierHandlerFunc(handlers.GetApiIdentifier)
	api.OpenGetResumeHandler = open.GetResumeHandlerFunc(handlers.GetResume)
	api.OpenCreateResumeHandler = open.CreateResumeHandlerFunc(handlers.CreateResume)
	return api
}
func init() {
	// Setup Logging
	slog.SetDefault(lib.NewAugmentedLogger())
	// Setup AWS SDK
	sdkConfig, err := config.LoadDefaultConfig(
		context.Background(),
		config.WithRegion(os.Getenv("REGION")),
	)
	if err != nil {
		slog.Error("unable to load SDK config", "error", err)
		os.Exit(1)
	}
	// Setup DynamoDB
	db.DBClient = db.NewClient(sdkConfig)
	api := setupHandlers()
	server := restapi.NewServer(api)
	server.ConfigureAPI()

	httpAdapter = httpadapter.New(server.GetHandler())
}

// Handler handles API requests
func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if nullHandler {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Body:       `{"message": "Null handler operational - always returns this message"}`,
		}, nil
	} else {
		return httpAdapter.Proxy(req)
	}
}

func main() {
	lambda.Start(Handler)
}
