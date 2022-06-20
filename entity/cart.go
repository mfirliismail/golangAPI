package entity

type Cart struct {
	ID         uint64   `gorm: "primary_key: auto_increment" json:"id"`
	CostumerID string   `gorm: "not null" json:"-"`
	Costumer   Costumer `gorm: "foreignkey: CostumerID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"costumer"`
}
