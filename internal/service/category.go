package service

import (
	"github.com/DanilMankiev/SofiaApplication/internal/entity"
	"github.com/DanilMankiev/SofiaApplication/internal/infrastructure/repository"
)

type CategoryService struct {
	repo repository.Category
}

func newCategoryService(repo repository.Category) *CategoryService{
	return &CategoryService{repo:repo}
}

func (cs *CategoryService) CreateCategory(category entity.Category) (int, error){
	return 0,nil

}
func (cs *CategoryService) 	GetAllCategorys() ([]entity.Category, error) {
	return []entity.Category{},nil
}
func (cs *CategoryService) GetCategoryById(id int) (entity.Category, error) {
	return entity.Category{},nil
}

func (cs *CategoryService) UpdateCategory(id int, input entity.UpdateCategory) error{
	return nil
}
func (cs* CategoryService) DeleteCategory(id int) error{
	return nil
}