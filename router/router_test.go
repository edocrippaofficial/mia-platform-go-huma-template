package router

import (
	"mia_template_service_name_placeholder/config"
	"mia_template_service_name_placeholder/controllers"
	"testing"

	"github.com/sirupsen/logrus/hooks/test"
	"gotest.tools/assert"
)

func TestSetupRouter(t *testing.T) {
	envs := config.Envs{}
	log, _ := test.NewNullLogger()

	handlers := controllers.NewHandlers(envs)

	_, err := SetupRouter(envs, log, handlers)
	assert.NilError(t, err, "unexpected error")
}
