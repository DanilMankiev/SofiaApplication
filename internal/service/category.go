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
	return cs.repo.CreateCategory(category)
}
func (cs *CategoryService) 	GetAllCategorys() ([]entity.CategoryResult, error) {
	return cs.repo.GetAllCategorys()
}
func (cs *CategoryService) GetCategoryById(id int) (entity.CategoryResult, error) {
	return cs.repo.GetCategoryById(id)
}
func (cs *CategoryService) UpdateCategory(id int, input entity.UpdateCategory) error{
	return cs.repo.UpdateCategory(id,input)
}
func (cs* CategoryService) DeleteCategory(id int) error{
	return cs.repo.DeleteCategory(id)
}