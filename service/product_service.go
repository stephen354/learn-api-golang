package service

import (
	"belajar-api/model"
	"belajar-api/repository"
)

type ProductService interface {
	GetAll() []model.ProductResponse
	GetById(id int) (model.ProductResponse, error)
	Create(product model.Product) model.Product
	Update(id int, product model.Product) error
	Delete(id int) error
}

type productService struct {
	productRepo  repository.ProductRepository
	categoryRepo repository.CategoryRepository
}

func NewProductService(productRepo repository.ProductRepository, categoryRepo repository.CategoryRepository) ProductService {
	return &productService{
		productRepo:  productRepo,
		categoryRepo: categoryRepo,
	}
}

func (s *productService) GetAll() []model.ProductResponse {
	products := s.productRepo.FindAll()
	var responses []model.ProductResponse

	for _, p := range products {
		category, _ := s.categoryRepo.FindById(p.CategoryID)
		responses = append(responses, model.ProductResponse{
			ID:         p.ID,
			Name:       p.Name,
			Price:      p.Price,
			CategoryID: p.CategoryID,
			Category:   category,
		})
	}
	return responses
}

func (s *productService) GetById(id int) (model.ProductResponse, error) {
	product, err := s.productRepo.FindById(id)
	if err != nil {
		return model.ProductResponse{}, err
	}

	category, _ := s.categoryRepo.FindById(product.CategoryID)
	return model.ProductResponse{
		ID:         product.ID,
		Name:       product.Name,
		Price:      product.Price,
		CategoryID: product.CategoryID,
		Category:   category,
	}, nil
}

func (s *productService) Create(product model.Product) model.Product {
	return s.productRepo.Save(product)
}

func (s *productService) Update(id int, product model.Product) error {
	return s.productRepo.Update(id, product)
}

func (s *productService) Delete(id int) error {
	return s.productRepo.Delete(id)
}
