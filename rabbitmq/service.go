package rabbitmq

var services map[string] interface{}

func init() {
	services = make(map[string] interface{})
}

func Register(methodName string, function interface{}){
	services[methodName] = function
}

func GetServices() map[string] interface{} {
	return services
}

func GetServiceByName(name string) interface{} {
	return services[name]
}