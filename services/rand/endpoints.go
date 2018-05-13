package rand

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func MakeGetRandomInt64Endpoint(service Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		randRequest := request.(GetRandomInt64Request)

		value, err := service.GetRandomInt64(randRequest)

		if err != nil {
			return nil, err
		}

		return RandInt64Response{value}, nil
	}
}
