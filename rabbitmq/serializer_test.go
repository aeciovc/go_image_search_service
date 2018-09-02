package rabbitmq

import (
	"testing"
	"github.com/streadway/amqp"
	test "github.com/aeciovc/go-image-search/test"
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

	test.Equals(t, responseEncoded, []byte(expectedResult))
}

func TestNamekoSerializerEncodeEmptyStringFromMethodSuccess(t *testing.T) {

	//Input
	methodResult := ""
	
	expectedResult := "{\"result\":\"\"}"

	serializer := &NamekoSerializer{}
	responseEncoded := serializer.Encode(methodResult)

	test.Equals(t, responseEncoded, []byte(expectedResult))
}

//TODO Implementing support to return int type
func TestNamekoSerializerEncodeIntegerFromMethodSuccess(t *testing.T) {

	//Input
	methodResult := 10
	
	expectedResult := "{\"result\":\"10\"}"

	serializer := &NamekoSerializer{}
	responseEncoded := serializer.Encode(methodResult)

	test.Equals(t, responseEncoded, []byte(expectedResult))
}