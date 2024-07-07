package repository

import (
	"github.com/DanilMankiev/SofiaApplication/internal/entity"
	"github.com/jmoiron/sqlx"
)

const (
	categoryTable="categories"
	usersTable="users"
)


type Authorization interface{
	SignUp(entity.RegiterInput) error
}

type Category interface {
	CreateCategory(input entity.Category) (int, error)
	GetAllCategorys() ([]entity.CategoryResult, error)
	GetCategoryById(id int) (entity.CategoryResult, error)
	UpdateCategory(id int, input entity.UpdateCategory) error
	DeleteCategory(id int) error
}
type Repository struct {
	Category
	Authorization
}

func New(db *sqlx.DB) *Repository{
	return &Repository{
		Category: newCategoryPostgres(db),
		Authorization: newAuthorizationPostgres(db),
	}
}