package rabbitmq

import(
	"reflect"
	"log"
)

type Invoker struct{
}

//Invoker
func (inv *Invoker) Invoke(method interface{}, args ...string) (interface{}, error){
	
	var error error

	response := invoke(method, args...)

	value := response.(reflect.Value)

	typeOfValue := value.Type()
	log.Println("[Invoker] Kind of value: ", typeOfValue.Kind())
	
	return value.Interface(), error
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

func fillStruct(t reflect.Type, v reflect.Value) {
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		ft := t.Field(i)
		switch ft.Type.Kind() {
		case reflect.Slice:
			f.Set(reflect.MakeSlice(ft.Type, 0, 0))
		case reflect.Struct:
			fillStruct(ft.Type, f)
		case reflect.Ptr:
			fv := reflect.New(ft.Type.Elem())
			fillStruct(ft.Type.Elem(), fv.Elem())
			f.Set(fv)
		default:
		}
	}
}

func GetInvoker() *Invoker{
	return &Invoker{}
}