package main

import (
    "log"
	"context"

    proto "github.com/aeciovc/go-image-search/proto"
	
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro"
)

func main() {

	service := micro.NewService()
	
	service.Init()

	// Create new greeter client
	imageService := proto.NewGoImageSearchService("go-image-search-service", service.Client())

	// Call the greeter
	rsp, err := imageService.Ping(context.TODO(), &proto.PingRequest{Name: "Testing service"})
	if err != nil {
		log.Println(err)
		return
	}

	// Print response
	log.Println(rsp.Message)

	
	log.Println("\n--- Multiple Requests ---")
	for i := 0; i < 10; i++ {
		call(i, service.Client())
	}
}

func call(i int, c client.Client) {
	
	req := c.NewRequest("go-image-search-service", "GoImageSearch.Ping", &proto.PingRequest{
		Name: "Testing service",
	})

	// create context with metadata
	ctx := metadata.NewContext(context.Background(), map[string]string{
		"X-User-Id": "john",
		"X-From-Id": "script",
	})

	rsp := &proto.PingResponse{}

	// Call service
	if err := c.Call(ctx, req, rsp); err != nil {
		log.Println("Error: ", err, rsp)
		return
	}

	log.Println("Call:", i, "rsp:", rsp.Message)
}