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

	//Encode method response to reply
	Encode() []byte
}


/************************
	Nameko Serializer
*************************/

type NamekoSerializer struct{}

type NamekoEncoder struct{
	Result string `json:"result"`
}

//Receiver
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

//Publisher
func (ns *NamekoSerializer) Marshall(c Call) (amqp.Delivery, error){
	return amqp.Delivery{}, nil
}

//Encoder
func (ns *NamekoSerializer) Encode(value interface{}) []byte{

	encode := NamekoEncoder{Result:value.(string)}

	result, err := buildJSON(encode)
	if err != nil{
		return []byte{}
	}

	return result
}