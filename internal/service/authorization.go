package service

import (
	"github.com/DanilMankiev/SofiaApplication/internal/entity"
	"github.com/DanilMankiev/SofiaApplication/internal/infrastructure/repository"
	"github.com/DanilMankiev/SofiaApplication/pkg/otp"
)

type AuthService struct {
	repo repository.Authorization
	otpGenerator otp.CodeGenerator
	codeLength int
}

func newAuthorizationService(repo repository.Authorization, otpGenerator otp.CodeGenerator, codeLength int) *AuthService{
	return &AuthService{
		repo: repo,
		otpGenerator: otpGenerator,
		codeLength: codeLength,
	}
}

func (au *AuthService) Register(input entity.RegiterInput) error{
	if err:=entity.ValidateRegisterInput(input);err!=nil{
		return err
	}



	return au.repo.SignUp(input)
}