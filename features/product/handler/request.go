package handler

type ProductRequest struct {
	Product_image string  `json:"product_image" form:"product_image"`
	Product_name  string  `json:"product_name" form:"product_name"`
	Description   string  `json:"description" form:"description"`
	Quantity      int     `json:"quantity" form:"quantity"`
	Price         string `json:"price" form:"price"`
}
