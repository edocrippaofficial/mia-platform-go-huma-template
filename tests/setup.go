package tests

import (
	"mia_template_service_name_placeholder/middlewares"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humamux"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func SetupTestRouter(log *logrus.Logger) (*mux.Router, huma.API) {
	router := mux.NewRouter()
	middlewares.SetLogging(router, log)
	api := humamux.New(router, huma.DefaultConfig("test", "1.0.0"))

	return router, api
}
