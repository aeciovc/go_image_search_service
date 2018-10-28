package main

import (
	"log"
	"fmt"
	"github.com/aeciovc/go-image-search/rabbitmq"
	
	imagesearch "github.com/aeciovc/go-image-search"
)

func main() {
	log.Println("[main] Initializing service...")

	//Load Configs
	config := imagesearch.LoadConfigs()

	// Build Server configs
	serverConfig := rabbitmq.ServerConfig{URI:config.Service.Broker}
	queueConfig := rabbitmq.QueueConfig{Name:config.Service.Queue}

	// Register functions services
	service := &imagesearch.RabbitMQService{}

	rabbitmq.Register("ping", service.Ping)
	rabbitmq.Register("search", service.Search)

	// Initialize the config
	rabbitmq.Init(serverConfig, queueConfig)

	// Connect
	rabbitmq.Connect(onError)

	// Run
	rabbitmq.Run()
}

func onError(err error) {
	if err != nil {
		log.Fatalf("[main] %s", err)
		panic(fmt.Sprintf("%s", err))
	}
}