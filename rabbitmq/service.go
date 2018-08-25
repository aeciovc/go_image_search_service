package rabbitmq

var services map[string]func() string

func init() {
	services = make(map[string]func() string)
}

func Register(methodName string, function func() string){
	services[methodName] = function
}

func GetServices() map[string]func() string {
	return services
}

func GetServiceByName(name string) func() string {
	return services[name]
}