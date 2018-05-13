package commands

import (
	"github.com/go-kit/kit/log"
)

type command struct {
	logger log.Logger
}

type Command interface {
	RandInt64() RandInt64
}

func NewCommand(logger log.Logger) Command {
	return &command{logger}
}
