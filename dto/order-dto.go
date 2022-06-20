package dto

type OrderUpdateDTO struct {
	ID        uint64 `json:"id" form:"id"`
	ProductID string `json:"product_id,omitempty" form:"product_id,omitempty"`
	CartID    string `json:"cart_id,omitempty" form:"cart_id,omitempty"`
}

type OrderCreateDTO struct {
	ProductID string `json:"product_id,omitempty" form:"product_id,omitempty"`
	CartID    string `json:"cart_id,omitempty" form:"cart_id,omitempty"`
}
