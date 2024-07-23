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

func newAuthorizationService(repo repository.Authorization,client *rabbitmq.Client, otpGenerator otp.CodeGenerator, codeLength int) *AuthService{
	return &AuthService{
		repo: repo,
		otpGenerator: otpGenerator,
		codeLength: codeLength,
		rabbitmq: client,
	}
}

func (au *AuthService) Register(input entity.RegiterInput) error{
	if err:=entity.ValidateRegisterInput(input);err!=nil{
		return err
	}

	return au.repo.Register(input)
}
 
func(au *AuthService) SendCodeEmail(email string) error{
	confirmCode:= au.otpGenerator.GenerateCode(au.codeLength)
	return au.repo.SendCodeEmail(email,confirmCode)
}

func(au *AuthService) SendCodeSMS( phone string) error {

}