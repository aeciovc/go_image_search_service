package main

import (
	"log"
	//"strings"
	"fmt"
    "github.com/streadway/amqp"

	//"github.com/jroimartin/rpcmq"
)

/*
func main() {
	s := rpcmq.NewServer("amqp://rabbitmq:rabbitmq@localhost",
		"rpc-storage_service", "nameko-rpc", "topic")
	if err := s.Register("storage_service.ping", ping); err != nil {
		log.Fatalf("Register: %v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("Init: %v", err)
	}
	defer s.Shutdown()

	forever := make(chan struct{})
	<-forever
}

func ping(id string, data []byte) ([]byte, error) {
	log.Printf("Received (%v): ping(%v)\n", id, string(data))
	return []byte(strings.ToUpper(string(data))), nil
}*/

func failOnError(err error, msg string) {
	if err != nil {
			log.Fatalf("%s: %s", msg, err)
			panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func ping() string {
	return "{\"result\":\"pong from go\"}"
}

func main() {
	conn, err := amqp.Dial("amqp://rabbitmq:rabbitmq@localhost")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
			"rpc-storage_service", // name
			true,       // durable
			false,       // delete when usused
			false,       // exclusive
			false,       // no-wait
			nil,         // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.Qos(
			1,     // prefetch count
			0,     // prefetch size
			false, // global
	)
	failOnError(err, "Failed to set QoS")

	msgs, err := ch.Consume(
			q.Name, // queue
			"",     // consumer
			false,  // auto-ack
			false,  // exclusive
			false,  // no-local
			false,  // no-wait
			nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
			for d := range msgs {
					n := string(d.Body)
					failOnError(err, "Failed to convert body to integer")

					log.Println(" [.] ping()", n)
					response := ping()
					
					err = ch.Publish(
							"nameko-rpc", // exchange
							d.ReplyTo, // routing key
							false,     // mandatory
							false,     // immediate
							amqp.Publishing{
									ContentType:   "application/json",
									CorrelationId: d.CorrelationId,
									Body:          []byte(response),
							})
					failOnError(err, "Failed to publish a message")

					d.Ack(false)
			}
	}()

	log.Printf(" [*] Awaiting RPC requests")
	<-forever
}