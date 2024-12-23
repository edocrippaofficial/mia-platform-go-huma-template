package main

import (
	"context"
	"fmt"
	"mia_template_service_name_placeholder/config"
	"mia_template_service_name_placeholder/controllers"
	"mia_template_service_name_placeholder/router"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
)

func main() {
	entrypoint(make(chan os.Signal, 1))
	os.Exit(0)
}

func entrypoint(sigtermCh chan os.Signal) {
	// Init envs struct and logger
	envs := config.MustGetEnvs()
	log := config.MustGetLogger(envs.LogLevel)

	// Create the handlers
	handlers := controllers.NewHandlers(envs)

	// Create the router
	r, err := router.SetupRouter(envs, log, handlers)
	if err != nil {
		log.Fatal("Unable to setup the router!")
	}

	srv := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf("0.0.0.0:%d", envs.HttpPort),
	}

	// Run the server
	go func(srv *http.Server, log *logrus.Logger, envs config.Envs) {
		log.WithFields(map[string]any{
			"port": envs.HttpPort,
			"pid":  os.Getpid(),
		}).Info("Starting server")
		log.Fatal(srv.ListenAndServe())
	}(srv, log, envs)

	// Block and wait for the SIGTERM signal
	handleGracefulShutdown(sigtermCh, srv, log, envs)
}

func handleGracefulShutdown(sigtermCh chan os.Signal, srv *http.Server, log *logrus.Logger, envs config.Envs) {
	signal.Notify(sigtermCh, syscall.SIGTERM)
	<-sigtermCh

	log.Info("Received SIGTERM signal")
	time.Sleep(time.Duration(envs.DelayShutdownSeconds) * time.Second)
	log.Info("Gracefully shutting down...")
	if err := srv.Shutdown(context.Background()); err != nil {
		panic(err.Error())
	}
}
