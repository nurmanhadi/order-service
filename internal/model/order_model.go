package model

type OrderRequestData struct {
	OrderId   string `json:"order_id" validate:"required,min=34,max=36"`
	UserId    string `json:"user_id" validate:"required,min=34,max=36"`
	ProductId int    `json:"product_id" validate:"required"`
	Price     int    `json:"price" validate:"required"`
	Quantity  int    `json:"quantity" validate:"required"`
}
