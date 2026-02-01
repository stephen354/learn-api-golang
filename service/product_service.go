package service

import (
	"belajar-api/model"
	"belajar-api/repository"
	"errors"
)

type ProductService interface {
	GetAll() []model.ProdukResponse
	GetById(id int) (model.ProdukResponse, error)
	Create(product model.Produk) (model.Produk, error)
	Update(id int, product model.Produk) error
	Delete(id int) error
}

type productService struct {
	productRepo  repository.ProductRepository
	categoryRepo repository.CategoryRepository
}

func NewProductService(pRepo repository.ProductRepository, cRepo repository.CategoryRepository) ProductService {
	return &productService{
		productRepo:  pRepo,
		categoryRepo: cRepo,
	}
}

func (s *productService) GetAll() []model.ProdukResponse {
	products := s.productRepo.FindAll()
	var responses []model.ProdukResponse

	for _, p := range products {
		category, _ := s.categoryRepo.FindById(p.CategoryID)
		responses = append(responses, model.ProdukResponse{
			ID:         p.ID,
			Nama:       p.Nama,
			Harga:      p.Harga,
			CategoryID: p.CategoryID,
			Category:   category,
		})
	}
	return responses
}

func (s *productService) GetById(id int) (model.ProdukResponse, error) {
	product, err := s.productRepo.FindById(id)
	if err != nil {
		return model.ProdukResponse{}, err
	}

	category, _ := s.categoryRepo.FindById(product.CategoryID)
	return model.ProdukResponse{
		ID:         product.ID,
		Nama:       product.Nama,
		Harga:      product.Harga,
		CategoryID: product.CategoryID,
		Category:   category,
	}, nil
}

func (s *productService) Create(product model.Produk) (model.Produk, error) {
	if !s.categoryRepo.Exists(product.CategoryID) {
		return model.Produk{}, errors.New("invalid CategoryID")
	}
	return s.productRepo.Save(product), nil
}

func (s *productService) Update(id int, product model.Produk) error {
	if product.CategoryID != 0 {
		if !s.categoryRepo.Exists(product.CategoryID) {
			return errors.New("invalid CategoryID")
		}
	}
	return s.productRepo.Update(id, product)
}

func (s *productService) Delete(id int) error {
	return s.productRepo.Delete(id)
}
