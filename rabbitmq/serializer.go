package rabbitmq

type Serializer interface {
	GetName() string
	Serialize() string
	GetResponse() (string, error)
}