package model

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

type ProdukResponse struct {
	ID         int      `json:"id"`
	Nama       string   `json:"nama"`
	Harga      int      `json:"harga"`
	CategoryID int      `json:"category_id"`
	Category   Category `json:"category"`
}
