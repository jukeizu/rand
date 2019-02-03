package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/jukeizu/contract"
)

func Rand(request contract.Request) (*contract.Response, error) {
	randInt64 := getRandInt64(1000)

	message := contract.Message{
		Content: fmt.Sprintf("%d", randInt64),
	}

	return &contract.Response{Messages: []*contract.Message{&message}}, nil
}

func getRandInt64(maxValue int64) int64 {
	seed := rand.NewSource(time.Now().UnixNano())

	r := rand.New(seed)

	return r.Int63n(maxValue)
}
