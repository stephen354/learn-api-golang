package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Category struct {
	ID   int    `json:"id"`
	Nama string `json:"nama"`
}

type Produk struct {
	ID         int    `json:"id"`
	Nama       string `json:"nama"`
	Harga      int    `json:"harga"`
	CategoryID int    `json:"category_id"`
}

var categories = []Category{
	{ID: 1, Nama: "Minuman"},
	{ID: 2, Nama: "Makanan"},
}

var produk = []Produk{
	{ID: 1, Nama: "Kopi Tubruk", Harga: 15000, CategoryID: 1},
	{ID: 2, Nama: "Kopi Susu", Harga: 20000, CategoryID: 1},
	{ID: 3, Nama: "Nasi Goreng", Harga: 25000, CategoryID: 2},
}

func main() {

	// API Category By ID (GET, PUT, DELETE)
	http.HandleFunc("/api/categories/", func(w http.ResponseWriter, r *http.Request) {
		idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if r.Method == "GET" {
			for _, c := range categories {
				if c.ID == id {
					w.Header().Set("Content-Type", "application/json")
					json.NewEncoder(w).Encode(c)
					return
				}
			}
			w.WriteHeader(http.StatusNotFound)
		} else if r.Method == "PUT" {
			var updatedCat Category
			if err := json.NewDecoder(r.Body).Decode(&updatedCat); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			for i, c := range categories {
				if c.ID == id {
					updatedCat.ID = id
					categories[i] = updatedCat
					w.WriteHeader(http.StatusOK)
					fmt.Fprint(w, "Category updated successfully")
					return
				}
			}
			w.WriteHeader(http.StatusNotFound)
		} else if r.Method == "DELETE" {
			for i, c := range categories {
				if c.ID == id {
					categories = append(categories[:i], categories[i+1:]...)
					w.WriteHeader(http.StatusOK)
					fmt.Fprint(w, "Category deleted successfully")
					return
				}
			}
			w.WriteHeader(http.StatusNotFound)
		}
	})

	// API Categories (GET, POST)
	http.HandleFunc("/api/categories", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(categories)
		} else if r.Method == "POST" {
			var c Category
			if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			// Generate ID based on current max ID to avoid duplicates after deletions
			maxID := 0
			for _, cat := range categories {
				if cat.ID > maxID {
					maxID = cat.ID
				}
			}
			c.ID = maxID + 1

			categories = append(categories, c)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(c)
		}
	})

	// API Produk By ID (GET, PUT, DELETE)
	http.HandleFunc("/api/produk/", func(w http.ResponseWriter, r *http.Request) {
		idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if r.Method == "GET" {
			for _, p := range produk {
				if p.ID == id {
					w.Header().Set("Content-Type", "application/json")
					json.NewEncoder(w).Encode(p)
					return
				}
			}
			w.WriteHeader(http.StatusNotFound)
		} else if r.Method == "PUT" {
			var updatedProd Produk
			if err := json.NewDecoder(r.Body).Decode(&updatedProd); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			for i, p := range produk {
				if p.ID == id {
					updatedProd.ID = id
					produk[i] = updatedProd
					w.WriteHeader(http.StatusOK)
					fmt.Fprint(w, "Product updated successfully")
					return
				}
			}
			w.WriteHeader(http.StatusNotFound)
		} else if r.Method == "DELETE" {
			for i, p := range produk {
				if p.ID == id {
					produk = append(produk[:i], produk[i+1:]...)
					w.WriteHeader(http.StatusOK)
					fmt.Fprint(w, "Product deleted successfully")
					return
				}
			}
			w.WriteHeader(http.StatusNotFound)
		}
	})

	// API Produk (GET, POST)
	http.HandleFunc("/api/produk", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(produk)
		} else if r.Method == "POST" {
			var p Produk
			if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			// Validate CategoryID
			categoryExists := false
			for _, c := range categories {
				if c.ID == p.CategoryID {
					categoryExists = true
					break
				}
			}

			if !categoryExists {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprint(w, "Invalid CategoryID")
				return
			}

			// Generate ID
			maxID := 0
			for _, prod := range produk {
				if prod.ID > maxID {
					maxID = prod.ID
				}
			}
			p.ID = maxID + 1

			produk = append(produk, p)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(p)
		}
	})

	fmt.Println("Server sedang berjalan di http://localhost:8080")
	fmt.Println("Tekan Ctrl+C untuk menghentikan server")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Gagal menjalankan server:", err)
	}
}
