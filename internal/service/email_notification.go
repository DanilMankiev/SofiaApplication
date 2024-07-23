package service

import (
	
	"github.com/DanilMankiev/SofiaApplication/pkg/rabbitmq"
)

type EmailService struct {
	clientBroker *rabbitmq.Client
	
}

func newEmailService(clientBroker *rabbitmq.Client) *EmailService {
	return &EmailService{
		clientBroker: clientBroker,
	}
}

func (n *EmailService) Get() error {
	return nil
}