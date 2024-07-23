package service

import (
	"github.com/DanilMankiev/SofiaApplication/internal/entity"
	"github.com/DanilMankiev/SofiaApplication/internal/infrastructure/repository"
	"github.com/DanilMankiev/SofiaApplication/pkg/otp"
	"github.com/DanilMankiev/SofiaApplication/pkg/rabbitmq"
)


type Category interface {
	CreateCategory(category entity.Category) (int, error)
	GetAllCategorys() ([]entity.CategoryResult, error)
	GetCategoryById(id int) (entity.CategoryResult, error)
	UpdateCategory(id int, input entity.UpdateCategory) error
	DeleteCategory(id int) error
}

type Email interface {
	
}

type SMS interface{

}

type Authorization interface{
	Register(entity.RegiterInput) error
	SendCodeEmail(email string) error
	SendCodeSMS(phone string) error
}

type Service struct {
	Category
	Authorization
	Email

}

func New(repo *repository.Repository, msgBroker *rabbitmq.Client, otp otp.CodeGenerator, verificationLength int) *Service {
	return &Service{
		Category: newCategoryService(repo.Category),
		Authorization: newAuthorizationService(repo.Authorization,msgBroker,otp,verificationLength),
	}
}