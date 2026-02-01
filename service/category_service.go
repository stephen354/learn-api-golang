package service

import (
	"belajar-api/model"
	"belajar-api/repository"
)

type CategoryService interface {
	GetAll() []model.Category
	GetById(id int) (model.Category, error)
	Create(category model.Category) model.Category
	Update(id int, category model.Category) error
	Delete(id int) error
}

type categoryService struct {
	categoryRepo repository.CategoryRepository
}

func NewCategoryService(categoryRepo repository.CategoryRepository) CategoryService {
	return &categoryService{categoryRepo: categoryRepo}
}

func (s *categoryService) GetAll() []model.Category {
	return s.categoryRepo.FindAll()
}

func (s *categoryService) GetById(id int) (model.Category, error) {
	return s.categoryRepo.FindById(id)
}

func (s *categoryService) Create(category model.Category) model.Category {
	return s.categoryRepo.Save(category)
}

func (s *categoryService) Update(id int, category model.Category) error {
	return s.categoryRepo.Update(id, category)
}

func (s *categoryService) Delete(id int) error {
	return s.categoryRepo.Delete(id)
}
