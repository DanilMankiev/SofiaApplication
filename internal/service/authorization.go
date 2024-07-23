package service

import (

	"github.com/DanilMankiev/SofiaApplication/internal/entity"
	"github.com/DanilMankiev/SofiaApplication/internal/infrastructure/repository"
	"github.com/DanilMankiev/SofiaApplication/pkg/otp"
	"github.com/DanilMankiev/SofiaApplication/pkg/rabbitmq"
)

type AuthService struct {
	repo repository.Authorization
	otpGenerator otp.CodeGenerator
	codeLength int
	rabbitmq *rabbitmq.Client
}

const(
	notificationExchange="notification"
	routingKeyEmail="email"
	routingKeySMS="sms"
)
func newAuthorizationService(repo repository.Authorization,client *rabbitmq.Client, otpGenerator otp.CodeGenerator, codeLength int) *AuthService{
	return &AuthService{
		repo: repo,
		otpGenerator: otpGenerator,
		codeLength: codeLength,
		rabbitmq: client,
	}
}

func (au *AuthService) Register(input entity.RegiterInput) error{
	// if err:=entity.ValidateRegisterInput(input);err!=nil{
	// 	return err
	// }

	return au.repo.Register(input)
}
 
func(au *AuthService) SendCodeEmail(email string) error{
	confirmCode:= au.otpGenerator.GenerateCode(au.codeLength)
	go au.rabbitmq.Publish(notificationExchange,routingKeyEmail,false,false,[]byte(confirmCode))
	return au.repo.SendCodeEmail(email,confirmCode)
}

func(au *AuthService) SendCodeSMS( phone string) error {
	confirmCode:=au.otpGenerator.GenerateCode(au.codeLength)
	go au.rabbitmq.Publish(notificationExchange,routingKeySMS,false,false,[]byte(confirmCode))
	return au.repo.SendCodeSMS(phone,confirmCode)
}