package main

import (
	"GetHouses/handler"
	example "GetHouses/proto/example"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
)

func main() {

	service := micro.NewService(
		micro.Name("go.micro.srv.GetHouses"),
		micro.Version("latest"),
	)

	service.Init()

	err := example.RegisterExampleHandler(service.Server(), new(handler.Example))
	if err != nil {
		return
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
