package handler

import (
	"belajar-api/model"
	"belajar-api/service"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type CategoryHandler interface {
	HandleCategories(w http.ResponseWriter, r *http.Request)
	HandleCategoryByID(w http.ResponseWriter, r *http.Request)
}

type categoryHandler struct {
	categoryService service.CategoryService
}

func NewCategoryHandler(categoryService service.CategoryService) CategoryHandler {
	return &categoryHandler{categoryService: categoryService}
}

func (h *categoryHandler) HandleCategories(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		categories := h.categoryService.GetAll()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(categories)
	case http.MethodPost:
		var c model.Category
		if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}

		createdCategory := h.categoryService.Create(c)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(createdCategory)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *categoryHandler) HandleCategoryByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid category ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		category, err := h.categoryService.GetById(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(category)
	case http.MethodPut:
		var c model.Category
		if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}

		if err := h.categoryService.Update(id, c); err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	case http.MethodDelete:
		if err := h.categoryService.Delete(id); err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}
