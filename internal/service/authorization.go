package service

import (
	"github.com/DanilMankiev/SofiaApplication/internal/entity"
	"github.com/DanilMankiev/SofiaApplication/internal/infrastructure/repository"
)

type AuthService struct {
	repo repository.Authorization
}

func newAuthorizationService(repo repository.Authorization) *AuthService{
	return &AuthService{repo: repo}
}

func (au *AuthService) SignUp(input entity.RegiterInput) error{
	return au.repo.SignUp(input)
}