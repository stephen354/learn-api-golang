package repository

import (
	"belajar-api/model"
	"errors"
)

type CategoryRepository interface {
	FindAll() []model.Category
	FindById(id int) (model.Category, error)
	Save(category model.Category) model.Category
	Update(id int, category model.Category) error
	Delete(id int) error
	// Helper for product validation
	Exists(id int) bool
}

type categoryRepository struct {
	categories []model.Category
}

func NewCategoryRepository() CategoryRepository {
	return &categoryRepository{
		categories: []model.Category{
			{ID: 1, Nama: "Minuman"},
			{ID: 2, Nama: "Makanan"},
		},
	}
}

func (r *categoryRepository) FindAll() []model.Category {
	return r.categories
}

func (r *categoryRepository) FindById(id int) (model.Category, error) {
	for _, c := range r.categories {
		if c.ID == id {
			return c, nil
		}
	}
	return model.Category{}, errors.New("category not found")
}

func (r *categoryRepository) Save(category model.Category) model.Category {
	maxID := 0
	for _, c := range r.categories {
		if c.ID > maxID {
			maxID = c.ID
		}
	}
	category.ID = maxID + 1
	r.categories = append(r.categories, category)
	return category
}

func (r *categoryRepository) Update(id int, category model.Category) error {
	for i, c := range r.categories {
		if c.ID == id {
			category.ID = id
			r.categories[i] = category
			return nil
		}
	}
	return errors.New("category not found")
}

func (r *categoryRepository) Delete(id int) error {
	for i, c := range r.categories {
		if c.ID == id {
			r.categories = append(r.categories[:i], r.categories[i+1:]...)
			return nil
		}
	}
	return errors.New("category not found")
}

func (r *categoryRepository) Exists(id int) bool {
	for _, c := range r.categories {
		if c.ID == id {
			return true
		}
	}
	return false
}
