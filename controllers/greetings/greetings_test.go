package greetings

import (
	"mia_template_service_name_placeholder/tests"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/sirupsen/logrus/hooks/test"
	"gotest.tools/assert"
)

func TestGetGreeting(t *testing.T) {
	greetings := Greetings{}
	log, _ := test.NewNullLogger()

	router, api := tests.SetupTestRouter(log)
	AddRoutes(api, greetings)

	r := httptest.NewRequest(http.MethodGet, "/hello?name=john", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, r)

	assert.Equal(t, w.Code, http.StatusOK, "Unexpected status code.")
	assert.Assert(t, strings.Contains(w.Body.String(), "Hello, john!"), "Unexpected body")
}

func TestGetGreetingFailsForBob(t *testing.T) {
	greetings := Greetings{}
	log, _ := test.NewNullLogger()

	router, api := tests.SetupTestRouter(log)
	AddRoutes(api, greetings)

	r := httptest.NewRequest(http.MethodGet, "/hello?name=Bob", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, r)

	assert.Equal(t, w.Code, http.StatusForbidden, "Unexpected status code.")
	assert.Assert(t, strings.Contains(w.Body.String(), "no greeting for Bob"), "Unexpected body")
}
