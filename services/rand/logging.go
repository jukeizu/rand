package rand

import (
	"github.com/go-kit/kit/log"
	"time"
)

type loggingService struct {
	logger log.Logger
	Service
}

func NewLoggingService(logger log.Logger, s Service) Service {
	return &loggingService{logger, s}
}

func (s *loggingService) GetRandomInt64(getRandomInt64Request GetRandomInt64Request) (result int64, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "GetRandomInt",
			"maxValue", getRandomInt64Request.MaxValue,
			"result", result,
			"took", time.Since(begin),
		)

	}(time.Now())

	result, err = s.Service.GetRandomInt64(getRandomInt64Request)

	return
}
