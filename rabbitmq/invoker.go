package rabbitmq

import(
	"reflect"
	"log"
	"errors"
)

type Invoker struct{
}

//Invoker
func (inv *Invoker) Invoke(method interface{}, args ...string) (interface{}, error){
	
	var result interface{}
	var error error

	//Get registered method
	//method := GetServiceByName(methodName)

	response := invoke(method, args...)

	value := response.(reflect.Value)
	typeOfValue := value.Type()
	
	log.Println("[Invoker] Kind of value: ", typeOfValue.Kind())
	
	if typeOfValue.Kind() == reflect.String{
		result = value.String()
	}else if typeOfValue.Kind() == reflect.Slice{
		slice, ok := value.Interface().([]string)
		if !ok{
			log.Println("[Invoker] Error to convert slice ", slice)
		}
		result = slice
	}else{ 
		//Unsupported type to return
		result = nil
		error = errors.New("Unsupported type to return")
	}

	return result, error
}

func invoke(fn interface{}, args ...string) interface{}{
	log.Printf("[invoker] args received %s ", args)
	
	//Parse values
	v := reflect.ValueOf(fn)
    rargs := make([]reflect.Value, len(args))
    for i, a := range args {
        rargs[i] = reflect.ValueOf(a)
    }
	
	//Invoke method
	result := v.Call(rargs)
	
	log.Printf("[invoker] Result Type from call %s ", result)
	return result[0]
}

func GetInvoker() *Invoker{
	return &Invoker{}
}