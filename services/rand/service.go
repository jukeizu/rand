package rand

import (
	"math/rand"
	"time"
)

type Service interface {
	GetRandomInt64(GetRandomInt64Request) (int64, error)
}

type service struct{}

func (service) GetRandomInt64(getRandomInt64Request GetRandomInt64Request) (int64, error) {
	var seed = rand.NewSource(time.Now().UnixNano())

	var randGen = rand.New(seed)

	var maxValue = getRandomInt64Request.MaxValue

	return randGen.Int63n(maxValue), nil
}

func NewService() Service {
	return &service{}
}
