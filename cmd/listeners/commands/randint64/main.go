package main

import (
	"os"

	"github.com/jukeizu/handler"
	"github.com/jukeizu/rand/listeners/commands"
	"github.com/shawntoffel/services-core/command"
	configreader "github.com/shawntoffel/services-core/config"
	"github.com/shawntoffel/services-core/logging"
)

var commandArgs command.CommandArgs

type Config struct {
	HandlerConfig handler.HandlerConfig
}

func init() {
	commandArgs = command.ParseArgs()
}

func main() {
	logger := logging.GetLogger("command.rand.int64", os.Stdout)

	config := Config{}

	err := configreader.ReadConfig(commandArgs.ConfigFile, &config)

	if err != nil {
		panic(err)
	}

	handler, err := handler.NewCommandHandler(logger, config.HandlerConfig)

	if err != nil {
		panic(err)
	}

	c := commands.NewCommand(logger)

	handler.Start(c.RandInt64())
}
