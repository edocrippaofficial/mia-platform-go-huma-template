package middlewares

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	glogrus "github.com/mia-platform/glogger/v4/loggers/logrus"
	gmux "github.com/mia-platform/glogger/v4/middleware/mux"
)

func SetLogging(r *mux.Router, log *logrus.Logger) {
	middlewareLog := glogrus.GetLogger(logrus.NewEntry(log))
	r.Use(gmux.RequestMiddlewareLogger(middlewareLog, []string{}))
}
