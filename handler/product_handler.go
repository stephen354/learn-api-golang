package handler

import (
	"belajar-api/model"
	"belajar-api/service"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type ProductHandler struct {
	service service.ProductService
}

func NewProductHandler(service service.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) HandleProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		h.GetAll(w, r)
	} else if r.Method == "POST" {
		h.Create(w, r)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h *ProductHandler) HandleProductByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if r.Method == "GET" {
		h.GetByID(w, r, id)
	} else if r.Method == "PUT" {
		h.Update(w, r, id)
	} else if r.Method == "DELETE" {
		h.Delete(w, r, id)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h *ProductHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(h.service.GetAll())
}

func (h *ProductHandler) GetByID(w http.ResponseWriter, r *http.Request, id int) {
	product, err := h.service.GetById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	var p model.Produk
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	createdProduct, err := h.service.Create(p)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdProduct)
}

func (h *ProductHandler) Update(w http.ResponseWriter, r *http.Request, id int) {
	var p model.Produk
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := h.service.Update(id, p); err != nil {
		// Could be not found or invalid category
		if err.Error() == "product not found" {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, err.Error())
		}
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Product updated successfully")
}

func (h *ProductHandler) Delete(w http.ResponseWriter, r *http.Request, id int) {
	if err := h.service.Delete(id); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Product deleted successfully")
}
