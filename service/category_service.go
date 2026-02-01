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
	repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) CategoryService {
	return &categoryService{repo: repo}
}

func (s *categoryService) GetAll() []model.Category {
	return s.repo.FindAll()
}

func (s *categoryService) GetById(id int) (model.Category, error) {
	return s.repo.FindById(id)
}

func (s *categoryService) Create(category model.Category) model.Category {
	return s.repo.Save(category)
}

func (s *categoryService) Update(id int, category model.Category) error {
	return s.repo.Update(id, category)
}

func (s *categoryService) Delete(id int) error {
	return s.repo.Delete(id)
}
