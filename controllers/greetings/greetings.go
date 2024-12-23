package greetings

import (
	"context"
	"fmt"
	"mia_template_service_name_placeholder/config"

	"github.com/danielgtaylor/huma/v2"
	glogrus "github.com/mia-platform/glogger/v4/loggers/logrus"
)

type Greetings struct {
	Envs config.Envs
}

type GetGreetingInput struct {
	Name string `query:"name" maxLength:"30" example:"world" doc:"Name to greet" required:"true"`
	Age  int    `query:"age" example:"30" doc:"Your age"`
}

type GetGreetingOutputBody struct {
	Message string `json:"message" example:"Hello, world!" doc:"Greeting message"`
}
type GetGreetingOutput struct {
	Body GetGreetingOutputBody
}

func (g *Greetings) GetGreeting(ctx context.Context, input *GetGreetingInput) (*GetGreetingOutput, error) {
	log := glogrus.FromContext(ctx)
	log.Info(fmt.Sprintf("FOO = %s", g.Envs.Foo))

	if input.Name == "Bob" {
		return nil, huma.Error403Forbidden("no greeting for Bob")
	}

	resp := &GetGreetingOutput{
		Body: GetGreetingOutputBody{
			Message: fmt.Sprintf("Hello, %s!", input.Name),
		},
	}
	return resp, nil
}
