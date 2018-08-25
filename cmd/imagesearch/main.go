package main

import (
    "log"
   
    proto "github.com/aeciovc/go-image-search/proto"
	imagesearch "github.com/aeciovc/go-image-search"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/cmd"
)

func main() {
	
	// Parse CLI flags
	cmd.Init()

	config := imagesearch.LoadConfigs()

	//Load Service
	service := micro.NewService(
		micro.Name(config.Service.Name),
	)

	service.Init()

	proto.RegisterGoImageSearchHandler(service.Server(), new(imagesearch.Ping))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

    log.Println("Image Search Service is up!")
}

