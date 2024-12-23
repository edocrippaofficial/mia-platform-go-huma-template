package router

import (
	"mia_template_service_name_placeholder/config"
	"mia_template_service_name_placeholder/controllers"
	"mia_template_service_name_placeholder/controllers/greetings"
	"mia_template_service_name_placeholder/middlewares"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humamux"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func SetupRouter(envs config.Envs, log *logrus.Logger, handlers *controllers.Handlers) (*mux.Router, error) {
	// Init Mux
	router := mux.NewRouter()
	middlewares.SetLogging(router, log)

	// Init Huma
	humaConfig := huma.DefaultConfig("mia_template_service_name_placeholder", envs.ServiceVersion)
	humaConfig.DocsPath = "/documentation/"
	humaConfig.OpenAPIPath = "/documentation/openapi"
	humaConfig.CreateHooks = make([]func(huma.Config) huma.Config, 0)
	api := humamux.New(router, humaConfig)

	// Register routes
	addRoutes(api, handlers)

	// Register health routes
	addHealthRoutes(api)

	return router, nil
}

func addRoutes(api huma.API, handlers *controllers.Handlers) {
	greetings.AddRoutes(api, handlers.Greetings)
}
