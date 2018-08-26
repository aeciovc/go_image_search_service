package rabbitmq

import (
	"github.com/streadway/amqp"
	"strings"
)

type Call struct{
	MethodName string
	ServiceName string
	ContentType string
	Params []string //d.Body
}

type Serializer interface {
	Marshall(Call) (amqp.Delivery, error)
	Unmarshall(amqp.Delivery) (Call, error)
}


/************************
	Nameko Serializer
*************************/

type NamekoSerializer struct{}

func (ns *NamekoSerializer) Unmarshall(d amqp.Delivery) (Call, error){
	
	contentType := d.ContentType
	serviceAndMethodName := strings.Split(d.RoutingKey, ".")
	
	serviceName := ""
	methodName := ""

	if (len(serviceAndMethodName) == 2){
		serviceName = serviceAndMethodName[0]
		methodName = serviceAndMethodName[1]
	}

	params := []string{}
		
	return Call{ContentType:contentType, ServiceName:serviceName, MethodName:methodName, Params:params}, nil
}

func (ns *NamekoSerializer) Marshall(c Call) (amqp.Delivery, error){
	return amqp.Delivery{}, nil
}