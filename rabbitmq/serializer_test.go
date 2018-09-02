package rabbitmq

import (
	"testing"
	"github.com/streadway/amqp"
	test "github.com/aeciovc/go-image-search/test"
	"log"
)

func TestNamekoSerializerUnmarshallSuccess(t *testing.T) {

	d := amqp.Delivery{ContentType:"application/json", RoutingKey:"search_image_service.ping", Body: []byte("{\"args\": [], \"kwargs\": {}}")}

	serializer := &NamekoSerializer{}
	call, err := serializer.Unmarshall(d)

	test.Equals(t, err, nil)
	test.Equals(t, call.MethodName, "ping")
	test.Equals(t, call.ServiceName, "search_image_service")
	test.Equals(t, call.ContentType, "application/json")
	test.Equals(t, call.Params, []string{})
}


func TestNamekoSerializerEncodeStringFromMethodSuccess(t *testing.T) {

	//Input
	methodResult := "pong"
	
	expectedResult := "{\"result\":\"pong\"}"

	serializer := &NamekoSerializer{}
	responseEncoded := serializer.Encode(methodResult)
	log.Println("Result ", string(responseEncoded))
	test.Equals(t, responseEncoded, []byte(expectedResult))
}