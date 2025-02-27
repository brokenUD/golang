package main

import (
	"context"
	// hello "github.com/micro/go-grpc/examples/greeter/server/proto/hello"
	"github.com/micro/go-micro"
	// "github.com/micro/go-micro/registry"
	// "github.com/micro/go-plugins/registry/consul"
	"log"
)

type Say struct{}

func (s *Say) Hello(ctx context.Context, req *hello.Request, rsp *hello.Response) error {
	log.Print("Received Say.Hello request")
	rsp.Msg = "Hello " + req.Name
	return nil

}

func main() {
	reg := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"123.207.79.96:8500",
		}
	})

	service := micro.NewService(
		micro.Name("go.micro.srv.greeter"),
		micro.Registry(reg),
	)

	// optionally setup command line usage
	service.Init()

	// Register Handlers
	hello.RegisterSayHandler(service.Server(), new(Say))

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
