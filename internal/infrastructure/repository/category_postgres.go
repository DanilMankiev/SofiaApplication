package repository

import (
	"github.com/DanilMankiev/SofiaApplication/internal/entity"
	"github.com/jmoiron/sqlx"
)

type CategoryPostgres struct {
	db *sqlx.DB
}

func newCategoryPostgres(db *sqlx.DB) *CategoryPostgres{
	return &CategoryPostgres{db:db}
}

func (cp *CategoryPostgres) CreateCategory(category entity.Category) (int, error){
	return 0,nil
}

func (cp *CategoryPostgres) GetAllCategorys() ([]entity.Category, error){
	return []entity.Category{},nil
}

func (cp *CategoryPostgres) GetCategoryById(id int) (entity.Category, error){
	return entity.Category{},nil
}

func (cp *CategoryPostgres) UpdateCategory(id int, input entity.UpdateCategory) error{
	return nil
}

func (cp *CategoryPostgres) DeleteCategory(id int) error{ 
	return nil
}