package entity

type Product struct {
	ID          uint64 `gorm: "primary_key: auto_increment" json:"id"`
	ProductName string `gorm: "type: varchar(255)" json:"email"`
	Jumlah      uint64 `gorm: "type: integer" json:"jumlah"`
	Price       uint64 `gorm: "type: integer" json:"price"`
	Category    string `gorm: "type: varchar(255)" json:"category"`
}
