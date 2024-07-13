package rabbitmq

import (
	"errors"

	"github.com/rabbitmq/amqp091-go"
)

type Rabbitmq struct {
	Connect *amqp091.Connection
	Channel *amqp091.Channel
}

func NewRabbitMQ(url string) (*Rabbitmq, error){
	conn,err:=amqp091.Dial(url)
	if err!=nil{
		return nil, err
	}
	ch,err:= conn.Channel()
	if err!=nil{
		return nil, err
	}

	err=ch.ExchangeDeclare(
		"notification",
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	if err!=nil{
		return nil, errors.New("error: declare exchange")
	}
	
	if err=ch.Qos(
		1,
		0,
		false,
	); err!=nil{
		return nil, errors.New("error: Qos options")
	}
	return &Rabbitmq{Connect: conn, Channel: ch},nil
}

