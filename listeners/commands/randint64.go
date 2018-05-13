package commands

import (
	"regexp"

	"github.com/go-kit/kit/log"

	"github.com/jukeizu/handler"
)

type RandInt64 interface {
	handler.Command
}

type randInt64 struct {
	command *command
	logger  log.Logger
}

func (c *command) RandInt64() RandInt64 {
	logger := log.With(c.logger, "command", "rand.int64")

	return &randInt64{c, logger}
}

func (r *randInt64) IsCommand(request handler.Request) (bool, error) {
	return regexp.MatchString(`!rand (.*?)`, request.Content)
}

func (r *randInt64) Handle(request handler.Request) (handler.Results, error) {
	result := handler.Result{
		Content: "test",
	}

	return handler.Results{result}, nil
}
