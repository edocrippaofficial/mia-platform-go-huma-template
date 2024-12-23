package router

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
)

func addHealthRoutes(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "ready",
		Method:      "GET",
		Path:        "/-/ready",
		Summary:     "Ready check",
		Description: "Ready check",
		Tags:        []string{"Health"},
	}, simpleOkHandler)

	huma.Register(api, huma.Operation{
		OperationID: "healthz",
		Method:      "GET",
		Path:        "/-/healthz",
		Summary:     "Health check",
		Description: "Health check",
		Tags:        []string{"Health"},
	}, simpleOkHandler)
}

type healthBody struct {
	Status string `json:"status"`
}
type healthOutput struct {
	Body healthBody
}

func simpleOkHandler(ctx context.Context, input *struct{}) (*healthOutput, error) {
	return &healthOutput{Body: healthBody{Status: "OK"}}, nil
}
