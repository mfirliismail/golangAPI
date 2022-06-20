package entity

type Costumer struct {
	ID       uint64 `gorm: "primary_key: auto_increment" json:"id"`
	Email    string `gorm: "uniqueIndex; type: varchar(255)" json:"email"`
	FullName string `gorm: "type: varchar(255)" json:"fullName"`
	Password string `gorm: "->;<-;not null" json:"-"`
	Token    string `gorm: "-" json:"token,omitempty"`
}
