package main

import (
	"net/http"
	"os"

	"context"

	"github.com/go-kit/kit/log"
	"github.com/jukeizu/rand/services/rand"
	"github.com/shawntoffel/services-core/command"
	configreader "github.com/shawntoffel/services-core/config"
	"github.com/shawntoffel/services-core/logging"
	"github.com/shawntoffel/services-core/runner"
)

var serviceArgs command.CommandArgs

func init() {
	serviceArgs = command.ParseArgs()
}

func main() {
	backgroundContext := context.Background()
	logger := logging.GetLogger("rand", os.Stdout)

	config := configreader.ServiceConfig{}
	err := configreader.ReadConfig(serviceArgs.ConfigFile, &config)

	if err != nil {
		panic(err)
	}

	randService := rand.NewService()
	randService = rand.NewLoggingService(log.With(logger, "component", "randService"), randService)

	httpLogger := log.With(logger, "component", "randService-http")
	mux := http.NewServeMux()
	mux.Handle("/rand/", rand.MakeHandler(backgroundContext, randService, httpLogger))

	runner.StartService(mux, logger, config)
}
