package repository

import (
	"belajar-api/model"
	"database/sql"
	"errors"
)

type ProductRepository interface {
	FindAll() []model.Product
	FindById(id int) (model.Product, error)
	Save(product model.Product) model.Product
	Update(id int, product model.Product) error
	Delete(id int) error
}

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) FindAll() []model.Product {
	rows, err := r.db.Query("SELECT id, name, price, category_id FROM products ORDER BY id ASC")
	if err != nil {
		return []model.Product{}
	}
	defer rows.Close()

	var products []model.Product
	for rows.Next() {
		var p model.Product
		rows.Scan(&p.ID, &p.Name, &p.Price, &p.CategoryID)
		products = append(products, p)
	}
	return products
}

func (r *productRepository) FindById(id int) (model.Product, error) {
	var p model.Product
	err := r.db.QueryRow("SELECT id, name, price, category_id FROM products WHERE id = $1", id).Scan(&p.ID, &p.Name, &p.Price, &p.CategoryID)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.Product{}, errors.New("product not found")
		}
		return model.Product{}, err
	}
	return p, nil
}

func (r *productRepository) Save(product model.Product) model.Product {
	err := r.db.QueryRow("INSERT INTO products (name, price, category_id) VALUES ($1, $2, $3) RETURNING id", product.Name, product.Price, product.CategoryID).Scan(&product.ID)
	if err != nil {
		return model.Product{}
	}
	return product
}

func (r *productRepository) Update(id int, product model.Product) error {
	res, err := r.db.Exec("UPDATE products SET name = $1, price = $2, category_id = $3 WHERE id = $4", product.Name, product.Price, product.CategoryID, id)
	if err != nil {
		return err
	}
	count, _ := res.RowsAffected()
	if count == 0 {
		return errors.New("product not found")
	}
	return nil
}

func (r *productRepository) Delete(id int) error {
	res, err := r.db.Exec("DELETE FROM products WHERE id = $1", id)
	if err != nil {
		return err
	}
	count, _ := res.RowsAffected()
	if count == 0 {
		return errors.New("product not found")
	}
	return nil
}
