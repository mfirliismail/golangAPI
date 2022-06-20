package entity

type Order struct {
	ID        uint64  `gorm: "primary_key: auto_increment" json:"id"`
	ProductID string  `gorm: "not null" json:"-"`
	Product   Product `gorm: "foreignkey: ProductID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"product"`
	CartID    string  `gorm: "not null" json:"-"`
	Cart      Cart    `gorm: "foreignkey: CartID; constraint:onUpdate:CASCADE, onDelete:CASCADE" json:"cart"`
}
