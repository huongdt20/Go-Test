package dto

type CreateProductRequest struct {
	ProductName string  `json:"product_name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int64   `json:"quantity"`
	CategoryID  int64   `json:"category_id"`
}
