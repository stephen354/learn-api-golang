package repository

import (
	"belajar-api/model"
	"errors"
)

type ProductRepository interface {
	FindAll() []model.Produk
	FindById(id int) (model.Produk, error)
	Save(product model.Produk) model.Produk
	Update(id int, product model.Produk) error
	Delete(id int) error
}

type productRepository struct {
	products []model.Produk
}

func NewProductRepository() ProductRepository {
	return &productRepository{
		products: []model.Produk{
			{ID: 1, Nama: "Kopi Tubruk", Harga: 15000, CategoryID: 1},
			{ID: 2, Nama: "Kopi Susu", Harga: 20000, CategoryID: 1},
			{ID: 3, Nama: "Nasi Goreng", Harga: 25000, CategoryID: 2},
		},
	}
}

func (r *productRepository) FindAll() []model.Produk {
	return r.products
}

func (r *productRepository) FindById(id int) (model.Produk, error) {
	for _, p := range r.products {
		if p.ID == id {
			return p, nil
		}
	}
	return model.Produk{}, errors.New("product not found")
}

func (r *productRepository) Save(product model.Produk) model.Produk {
	maxID := 0
	for _, p := range r.products {
		if p.ID > maxID {
			maxID = p.ID
		}
	}
	product.ID = maxID + 1
	r.products = append(r.products, product)
	return product
}

func (r *productRepository) Update(id int, product model.Produk) error {
	for i, p := range r.products {
		if p.ID == id {
			product.ID = id
			r.products[i] = product
			return nil
		}
	}
	return errors.New("product not found")
}

func (r *productRepository) Delete(id int) error {
	for i, p := range r.products {
		if p.ID == id {
			r.products = append(r.products[:i], r.products[i+1:]...)
			return nil
		}
	}
	return errors.New("product not found")
}
