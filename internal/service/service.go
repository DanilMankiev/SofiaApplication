package service

import (
	"github.com/DanilMankiev/SofiaApplication/internal/entity"
	"github.com/DanilMankiev/SofiaApplication/internal/infrastructure/repository"
	"github.com/DanilMankiev/SofiaApplication/internal/service/notification"
	"github.com/DanilMankiev/SofiaApplication/internal/service/pap"
)


type Category interface {
	CreateCategory(category entity.Category) (int, error)
	GetAllCategorys() ([]entity.CategoryResult, error)
	GetCategoryById(id int) (entity.CategoryResult, error)
	UpdateCategory(id int, input entity.UpdateCategory) error
	DeleteCategory(id int) error
}

type Notification interface {
	Get() error
}

type Pa interface{
	Get() error
}

type Authorization interface{
	SignUp(entity.RegiterInput) error
}

type Service struct {
	Category
	Authorization
	Notification
	Pa
}


func New(repo *repository.Repository) *Service {
	return &Service{
		Category: newCategoryService(repo.Category),
		Authorization: newAuthorizationService(repo.Authorization),
		Notification: notification.NewNoti(),
		Pa: pa.NewNoti(),
	}
}