package rabbitmq

import "errors"

//RPC
var (
	// ErrMethodNotFound is returned when there is no method to invoke
	ErrMethodNotFound = errors.New("Method not found")

	// ErrRemote is returned when an unkonw error occurred
	ErrRemote = errors.New("Unknow error")
)