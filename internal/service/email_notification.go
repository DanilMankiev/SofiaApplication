package service

import (
	
	"github.com/DanilMankiev/SofiaApplication/pkg/rabbitmq"
)

type EmailService struct {
	msgBroker *rabbitmq.Rabbitmq
	
}

func newEmailService(msgBroker *rabbitmq.Rabbitmq) *EmailService {
	return &EmailService{
		msgBroker: msgBroker,
	}
}

func (n *EmailService) Get() error {
	return nil
}