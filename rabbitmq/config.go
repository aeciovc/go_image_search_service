package rabbitmq

type ServerConfig struct {
	URI 	string
}

type QueueConfig struct {
	Name 	string
}

type SerializerConfig struct{
	Serializer Serializer
}