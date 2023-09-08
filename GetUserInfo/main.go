package main

import (
	"GetUserInfo/handler"
	example "GetUserInfo/proto/example"
	"github.com/micro/go-micro"
	"log"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.GetUserInfo"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
