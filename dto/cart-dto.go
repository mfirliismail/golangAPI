package dto

type CartUpdateDTO struct {
	ID         uint64 `json:"id" form:"id" binding:"required"`
	CostumerID string `json:"costumer_id,omitempty" form:"costumer_id,omitempty"`
}

type CartCreateDTO struct {
	CostumerID string `json:"costumer_id,omitempty" form:"costumer_id,omitempty"`
}
