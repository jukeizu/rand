package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/jukeizu/contract"
	"github.com/rs/xid"
	"github.com/rs/zerolog"
)

var Version = ""

var (
	flagPort = "10000"
)

func parseCli() {
	flag.StringVar(&flagPort, "p", flagPort, "port")
	flag.Parse()
}

func main() {
	parseCli()
	port := ":" + flagPort

	logger := zerolog.New(os.Stdout).With().Timestamp().
		Str("instance", xid.New().String()).
		Str("component", "intent.endpoint.rand").
		Str("version", Version).
		Logger()

	mux := http.NewServeMux()
	mux.HandleFunc("/rand", contract.NewHttpHandler(Rand))

	logger.Info().Str("address", port).Msg("listening")
	http.ListenAndServe(port, mux)
}
