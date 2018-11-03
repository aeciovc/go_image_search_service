package rabbitmq

import (
	"github.com/streadway/amqp"
	"strings"
	"log"
	"encoding/json"
)

/*************************
	Nameko Serializer
**************************/

type NamekoSerializer struct{
}

type NamekoEncoder struct{
	Result interface{} `json:"result"`
	Error interface{}  `json:"error"`
}

type NamekoError struct{
	ExcPath string `json:"exc_path"`
}

var namekoErrors = map[error]string{
	ErrMethodNotFound: "nameko.exceptions.MethodNotFound",
	ErrRemote: "nameko.exceptions.RemoteError",
}

//Python args
type Params struct{
	Args []string
	//KWargs []string //not supported yet
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

	//Get method params
	params := ns.getParams(d.Body, &Params{})
		
	return Call{ContentType:contentType, ServiceName:serviceName, MethodName:methodName, Params:params}, nil
}

//Publisher
func (ns *NamekoSerializer) Marshall(c Call) (amqp.Delivery, error){
	return amqp.Delivery{}, nil
}

//Encoder
func (ns *NamekoSerializer) Encode(v interface{}, err error) []byte{
	
	//Get map error
	errResult := ns.getError(err)

	//Build result
	encode := &NamekoEncoder{Result:v, Error:errResult}

	result, err := buildJSON(encode)
	if err != nil{
		log.Printf("[NamekoSerializer] Error building JSON response: %s ", err)
		return []byte{}
	}

	log.Println("[NamekoSerializer] Returning result: ", string(result))
	return result
}

func (ns *NamekoSerializer) getParams(body []byte, p *Params) []string{

	if p == nil{
		return []string{}
	}

	log.Println("[NamekoSerializer] Body ", string(body))
	if err := json.Unmarshal(body, p); err != nil{
		log.Printf("[NamekoSerializer] Error unmarshall params: %s", err)
	}

	return p.Args
}

func (ns *NamekoSerializer) getError(err error) interface{}{
	if err != nil{
		return NamekoError{ExcPath:namekoErrors[err]}
	}

	return ""
}