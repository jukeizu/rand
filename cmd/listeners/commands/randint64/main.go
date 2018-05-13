package main

import (
	"os"

	base "github.com/jukeizu/client-base"
	"github.com/jukeizu/handler"
	"github.com/jukeizu/rand/listeners/commands"
	"github.com/shawntoffel/services-core/command"
	"github.com/shawntoffel/services-core/config"
	"github.com/shawntoffel/services-core/logging"
)

var commandArgs command.CommandArgs

type RandHandlerConfig struct {
	HandlerConfig    handler.HandlerConfig
	RandClientConfig base.ClientConfig
}

func init() {
	commandArgs = command.ParseArgs()
}

func main() {

	logger := logging.GetLogger("command.rand.int64", os.Stdout)

	handlerConfig := RandHandlerConfig{}

	err := config.ReadConfig(commandArgs.ConfigFile, &handlerConfig)

	if err != nil {
		panic(err)
	}

	handler, err := handler.NewCommandHandler(logger, handlerConfig.HandlerConfig)

	if err != nil {
		panic(err)
	}

	c := commands.NewCommand(logger)

	handler.Start(c.RandInt64())
}
