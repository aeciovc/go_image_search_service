package main

import (
	"log"
	"fmt"
	"github.com/aeciovc/go-image-search/rabbitmq"
)

func main() {
	log.Println("Initializing service...")

	// Build configs
	serverConfig := rabbitmq.ServerConfig{URI:"amqp://rabbitmq:rabbitmq@localhost"}
	queueConfig := rabbitmq.QueueConfig{Name:"rpc-storage_service"}

	// Register functions services
	rabbitmq.Register("ping", ping)
	log.Println(rabbitmq.GetServices())

	// Initialize the config
	rabbitmq.Init(serverConfig, queueConfig)

	// Connect
	rabbitmq.Connect(onError)

	// Run
	rabbitmq.Run()

}

func onError(err error) {
	if err != nil {
			log.Fatalf("[ImageSearchService] %s", err)
			panic(fmt.Sprintf("%s", err))
	}
}

func ping() string {
	return "{\"result\":\"pong\"}"
}