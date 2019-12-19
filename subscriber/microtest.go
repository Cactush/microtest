package subscriber

import (
	"context"
	"github.com/micro/go-micro/util/log"

	microtest "github.com/Cactush/microtest/proto/microtest"
)

type Microtest struct{}

func (e *Microtest) Handle(ctx context.Context, msg *microtest.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *microtest.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
