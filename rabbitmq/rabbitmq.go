package rabbitmq

import (
	"log"
	"fmt"
    "github.com/streadway/amqp"
)

var onError func(err error)
var serverConfig ServerConfig
var queueConfig QueueConfig

var consumeChannel <-chan amqp.Delivery
var channel *amqp.Channel

func Init(serverConf ServerConfig, queueConf QueueConfig){
	serverConfig = serverConf
	queueConfig = queueConf
}

func Connect(errFunc func(err error)){

	onError = errFunc

	//Connect
	conn, err := amqp.Dial(serverConfig.URI)
	if err != nil{
		log.Println("[rabbitmq] Failed to connect to RabbitMQ")
		onError(err)
		return
	}

	//Channel
	channel, err = conn.Channel()
	if err != nil{
		log.Println("[rabbitmq] Failed to open a channel")
		onError(err)
		return
	}

	//Queue
	q, err := channel.QueueDeclare(
			queueConfig.Name, // name
			true,       // durable
			false,       // delete when usused
			false,       // exclusive
			false,       // no-wait
			nil,         // arguments
	)
	if err != nil{
		log.Println("[rabbitmq] Failed to declare a queue")
		onError(err)
		return
	}

	//QoS
	err = channel.Qos(
			1,     // prefetch count
			0,     // prefetch size
			false, // global
	)
	if err != nil{
		log.Println("[rabbitmq] Failed to set QoS")
		onError(err)
		return
	}

	//Consumer
	consumeChannel, err = channel.Consume(
			q.Name, // queue
			"",     // consumer
			false,  // auto-ack
			false,  // exclusive
			false,  // no-local
			false,  // no-wait
			nil,    // args
	)
	if err != nil{
		log.Println("[rabbitmq] Failed to register a consumer")
		onError(err)
		return
	}
}

func Run() error{

	forever := make(chan bool)

	go func() {
			for d := range consumeChannel {
					n := string(d.Body)
					log.Println(" [.] details:", d.Exchange)
					log.Println(" [.] details:", d.RoutingKey) // this is the method
					log.Println(" [.] details:", d.ReplyTo)
					log.Println(" [.] details:", d)

					log.Println(" [.] ping()", n)
					f := GetServiceByName("ping")
					response := f()
					
					err := channel.Publish(
							"nameko-rpc", // exchange
							d.ReplyTo, // routing key
							false,     // mandatory
							false,     // immediate
							amqp.Publishing{
									ContentType:   "application/json",
									CorrelationId: d.CorrelationId,
									Body:          []byte(response),
							})
					
					if err != nil{
						log.Println("[rabbitmq] Failed to reply the message")
						onError(err)
						return
					}

					d.Ack(false)
			}
	}()

	log.Printf(" [*] Awaiting RPC requests")
	<-forever

	return nil
}

func failOnError(err error, msg string) {
	if err != nil {
			log.Fatalf("%s: %s", msg, err)
			panic(fmt.Sprintf("%s: %s", msg, err))
	}
}