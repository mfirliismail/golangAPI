package dto

type ProductUpdateDTO struct {
	ID          uint64 `json:"id" form:"id" binding:"required"`
	ProductName string `json:"email" form:"email" binding:"required"`
	Jumlah      uint64 `json:"jumlah" form:"jumlah" binding:"required"`
	Price       uint64 `json:"price" form:"price" binding:"required"`
	Category    string `json:"category" form:"category" binding:"required"`
}

type ProductCreateDTO struct {
	ProductName string `json:"email" form:"email" binding:"required"`
	Jumlah      uint64 `json:"jumlah" form:"jumlah" binding:"required"`
	Price       uint64 `json:"price" form:"price" binding:"required"`
	Category    string `json:"category" form:"category" binding:"required"`
}
