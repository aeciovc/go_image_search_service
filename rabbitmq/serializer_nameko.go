package rabbitmq

import (
	"github.com/streadway/amqp"
	"strings"
	"log"
	"reflect"
	"encoding/json"
)

/*************************
	Nameko Serializer
**************************/

type NamekoSerializer struct{}

type NamekoEncoder struct{
	Result interface{} `json:"result"`
	Error string 	   `json:"error"`
}

//Python args
type Params struct{
	Args []string
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
func (ns *NamekoSerializer) Encode(v interface{}) []byte{

	value := v.(reflect.Value)
	typeOfValue := value.Type()
	
	log.Println("[NamekoSerializer] Kind of value: ", typeOfValue.Kind())
	
	encode := &NamekoEncoder{}
	if typeOfValue.Kind() == reflect.String{
		encode = &NamekoEncoder{Result:value.String()}
	}else if typeOfValue.Kind() == reflect.Slice{
		slice, ok := value.Interface().([]string)
		if !ok{
			log.Println("[NamekoSerializer] Error to convert slice ", slice)
		}
		encode = &NamekoEncoder{Result:slice}
	}else{
		//Unsupported type to return
		encode = &NamekoEncoder{Error:"Unsupported type to return"}
	}

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

	if err := json.Unmarshal(body, p); err != nil{
		log.Fatalf("[NamekoSerializer] Error unmarshall params: %s", err)
	}

	return p.Args
}