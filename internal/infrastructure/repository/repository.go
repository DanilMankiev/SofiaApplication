package repository

import (
	"github.com/DanilMankiev/SofiaApplication/internal/entity"
	"github.com/jmoiron/sqlx"
)

type Category interface {
	CreateCategory(category entity.Category) (int, error)
	GetAllCategorys() ([]entity.Category, error)
	GetCategoryById(id int) (entity.Category, error)
	UpdateCategory(id int, input entity.UpdateCategory) error
	DeleteCategory(id int) error
}
type Repository struct {
	Category
}

func New(db *sqlx.DB) *Repository{
	return &Repository{
		Category: newCategoryPostgres(db),
	}
}