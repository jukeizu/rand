package main

import (
	"flag"
	"net/http"

	"github.com/jukeizu/contract"
)

var (
	flagPort = "10000"
)

func parseCli() {
	flag.StringVar(&flagPort, "p", flagPort, "port")
	flag.Parse()
}

func main() {
	parseCli()

	mux := http.NewServeMux()

	mux.HandleFunc("/rand", contract.NewHttpHandler(Rand))

	http.ListenAndServe(":"+flagPort, mux)
}
