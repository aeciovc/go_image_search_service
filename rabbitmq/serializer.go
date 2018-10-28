package rabbitmq

import (
	"github.com/streadway/amqp"
)

type Call struct{
	MethodName string
	ServiceName string
	ContentType string
	Params []string
}

type Serializer interface {
	Marshall(Call) (amqp.Delivery, error)
	Unmarshall(amqp.Delivery) (Call, error)

	Encode(interface{}) []byte
}

func GetSerializer() Serializer{
	return &NamekoSerializer{}
}