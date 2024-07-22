package rabbitmq

import "github.com/rabbitmq/amqp091-go"

type Exchange struct {
	Name                                  string
	Kind                                  string
	Durable, AutoDelete, Internal, NoWait bool
	Args                                  amqp091.Table
}

type Queue struct {
	Name                                   string
	Durable, AutoDelete, Exclusive, NoWait bool
	Args                               	  amqp091.Table
}

type Binding struct {
	ExchangeName string
	QueueName    string
	RoutingKey   string
	NoWait       bool
	Args         amqp091.Table
}

type Config struct{
	Exchanges []Exchange
	Queues []Queue
	Bindings []Binding
}

