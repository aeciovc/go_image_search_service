package rabbitmq

import(
	"encoding/json"
	"reflect"
	"log"
)

func buildJSON(struc interface{}) ([]byte, error) {
	resp, err := json.Marshal(struc)

	if err != nil {
		log.Printf("[commons] Couldn't marshall object. %s", err.Error())
		return nil, err
	}

	return resp, err
}

func invoke(fn interface{}, args ...string) interface{}{
	log.Printf("[commons] args received %s ", args)
	
	//Parse values
	v := reflect.ValueOf(fn)
    rargs := make([]reflect.Value, len(args))
    for i, a := range args {
        rargs[i] = reflect.ValueOf(a)
    }
	
	//Invoke method
	result := v.Call(rargs)
	
	log.Printf("[commons] Result Type from call %s ", result)
	return result[0]
}