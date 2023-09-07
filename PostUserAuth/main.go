package main

import (
	"IHome/PostUserAuth/handler"
	example "IHome/PostUserAuth/proto/example"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/service/grpc"
)

func main() {
	// New Service
	service := grpc.NewService(
		micro.Name("go.micro.srv.PostUserAuth"),
		micro.Version("latest"),
	)

	service.Init()

	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
