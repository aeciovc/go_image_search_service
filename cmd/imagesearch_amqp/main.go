package main

import (
    "log"

	"github.com/micro/go-micro/cmd"
	_"github.com/micro/go-plugins/broker/rabbitmq"
	"github.com/micro/go-micro/broker"
)

func main() {
	
	// Parse CLI flags
	cmd.Init()

	if err := broker.Init(); err != nil {
		log.Fatalf("Broker Init error: %v", err)
	}
	if err := broker.Connect(); err != nil {
		log.Fatalf("Broker Connect error: %v", err)
	}

	_, err := broker.Subscribe("service.topic", func(p broker.Publication) error {
		log.Println("[sub] received message:", string(p.Message().Body), "header", p.Message().Header)
		return nil
	})
	
	if err != nil {
		log.Println(err)
	}

	forever := make(chan struct{})
	<-forever
	//r := rabbitmq.NewBroker(broker.Addrs("amqp://guest:guest@rabbit:5672/"))

    log.Println("Image Search Service is up!")
}

