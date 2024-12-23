package greetings

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

func AddRoutes(api huma.API, controller Greetings) {
	huma.Register(api, huma.Operation{
		OperationID: "get-greeting",
		Method:      http.MethodGet,
		Path:        "/hello",
		Summary:     "Get a greeting",
		Description: "Get a greeting for a person by name, unless you are Bob.",
		Tags:        []string{"Greetings"},
	}, controller.GetGreeting)
}
