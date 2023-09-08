package main

import (
	"DeleteSession/handler"
	example "DeleteSession/proto/example"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/service/grpc"
)

func main() {
	// New Service
	service := grpc.NewService(
		micro.Name("go.micro.srv.DeleteSession"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
