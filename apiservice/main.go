package main

import (
	"log"
	"time"

	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"

	//self defined api gin server
	apiGinServer "github.com/hyperbting/GoMicroservicePractice/apiservice/apiginserver"
)

func main() {

	apiSrv := apiGinServer.BuildServer()

	service := micro.NewService(
		micro.Name(apiSrv.Options().Name),
		micro.Server(apiSrv),
		micro.Registry(registry.NewRegistry()),
		//
		micro.RegisterTTL(time.Minute),
		micro.RegisterInterval(15 * time.Second),
	)

	service.Init()

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
