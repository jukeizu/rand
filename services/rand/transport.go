package rand

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/log"
	httpTransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

type GetRandomInt64Request struct {
	HttpRequest *http.Request
	MaxValue    int64 `json:"maxValue"`
}

type RandInt64Response struct {
	Value int64 `json:"value"`
}

func DecodeRandInt64Request(_ context.Context, r *http.Request) (request interface{}, err error) {
	randRequest := GetRandomInt64Request{}

	err = json.NewDecoder(r.Body).Decode(&randRequest)

	if err != nil {
		return nil, err
	}

	return randRequest, nil
}

func EncodeRandResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func MakeHandler(ctx context.Context, s Service, logger log.Logger) http.Handler {
	opts := []httpTransport.ServerOption{
		httpTransport.ServerErrorLogger(logger),
	}

	getRandomIntHandler := httpTransport.NewServer(
		MakeGetRandomInt64Endpoint(s),
		DecodeRandInt64Request,
		EncodeRandResponse,
		opts...,
	)

	router := mux.NewRouter()

	router.Handle("/rand/int64", getRandomIntHandler).Methods("POST")

	return router
}
