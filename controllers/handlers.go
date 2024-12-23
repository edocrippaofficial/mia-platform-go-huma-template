package controllers

import (
	"mia_template_service_name_placeholder/config"
	"mia_template_service_name_placeholder/controllers/greetings"
)

type Handlers struct {
	Greetings greetings.Greetings
}

func NewHandlers(envs config.Envs) *Handlers {
	return &Handlers{
		Greetings: greetings.Greetings{
			Envs: envs,
		},
	}
}
