package handler

import (
	"context"

	"github.com/micro/go-micro/util/log"

	microtest "github.com/Cactush/microtest/proto/microtest"
)

type Microtest struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Microtest) Call(ctx context.Context, req *microtest.Request, rsp *microtest.Response) error {
	log.Log("Received Microtest.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Microtest) Stream(ctx context.Context, req *microtest.StreamingRequest, stream microtest.Microtest_StreamStream) error {
	log.Logf("Received Microtest.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&microtest.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Microtest) PingPong(ctx context.Context, stream microtest.Microtest_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&microtest.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
