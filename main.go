package main

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	"github.com/Cactush/microtest/handler"
	"github.com/Cactush/microtest/subscriber"

	microtest "github.com/Cactush/microtest/proto/microtest"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.microtest"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	microtest.RegisterMicrotestHandler(service.Server(), new(handler.Microtest))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.microtest", service.Server(), new(subscriber.Microtest))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.microtest", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
/*

反正我只是测试一下
*/
/*
再次测试
*/