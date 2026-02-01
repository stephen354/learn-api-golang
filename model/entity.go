package model

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Product struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	CategoryID int    `json:"category_id"`
}

type ProductResponse struct {
	ID         int      `json:"id"`
	Name       string   `json:"name"`
	Price      int      `json:"price"`
	CategoryID int      `json:"category_id"`
	Category   Category `json:"category"`
}
