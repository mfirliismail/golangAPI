package dto

type CostumerUpdateDTO struct {
	ID       uint64 `json:"id" form:"id" binding:"required"`
	Email    string `json:"email" form:"email" validate:"email"`
	FullName string `json:"fullName" form:"fullName"`
	Password string `json:"password,omitempty" form:"password,omitempty" validate:"min:6"`
}

type CostumerCreateDTO struct {
	Email    string `json:"email" form:"email" validate:"email"`
	FullName string `json:"fullName" form:"fullName"`
	Password string `json:"password,omitempty" form:"password,omitempty" validate:"min:6" binding:"required"`
}
