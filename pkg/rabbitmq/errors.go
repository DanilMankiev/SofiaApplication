package rabbitmq

import "errors"

var (
	errNotConnected = errors.New("not connected to server")
	errNotValidMsg = errors.New("not valid message") 
	errShutDown = errors.New("client is shutting down")
)