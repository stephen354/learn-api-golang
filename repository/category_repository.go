package repository

import (
	"belajar-api/model"
	"database/sql"
	"errors"
)

type CategoryRepository interface {
	FindAll() []model.Category
	FindById(id int) (model.Category, error)
	Save(category model.Category) model.Category
	Update(id int, category model.Category) error
	Delete(id int) error
	Exists(id int) bool
}

type categoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) FindAll() []model.Category {
	rows, err := r.db.Query("SELECT id, name FROM categories ORDER BY id ASC")
	if err != nil {
		return []model.Category{}
	}
	defer rows.Close()

	var categories []model.Category
	for rows.Next() {
		var c model.Category
		rows.Scan(&c.ID, &c.Name)
		categories = append(categories, c)
	}
	return categories
}

func (r *categoryRepository) FindById(id int) (model.Category, error) {
	var c model.Category
	err := r.db.QueryRow("SELECT id, name FROM categories WHERE id = $1", id).Scan(&c.ID, &c.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.Category{}, errors.New("category not found")
		}
		return model.Category{}, err
	}
	return c, nil
}

func (r *categoryRepository) Save(category model.Category) model.Category {
	err := r.db.QueryRow("INSERT INTO categories (name) VALUES ($1) RETURNING id", category.Name).Scan(&category.ID)
	if err != nil {
		return model.Category{}
	}
	return category
}

func (r *categoryRepository) Update(id int, category model.Category) error {
	res, err := r.db.Exec("UPDATE categories SET name = $1 WHERE id = $2", category.Name, id)
	if err != nil {
		return err
	}
	count, _ := res.RowsAffected()
	if count == 0 {
		return errors.New("category not found")
	}
	return nil
}

func (r *categoryRepository) Delete(id int) error {
	res, err := r.db.Exec("DELETE FROM categories WHERE id = $1", id)
	if err != nil {
		return err
	}
	count, _ := res.RowsAffected()
	if count == 0 {
		return errors.New("category not found")
	}
	return nil
}

func (r *categoryRepository) Exists(id int) bool {
	var exists bool
	err := r.db.QueryRow("SELECT EXISTS(SELECT 1 FROM categories WHERE id = $1)", id).Scan(&exists)
	if err != nil {
		return false
	}
	return exists
}
