package models

type Items struct {
	IdItem      uint   `gorm:"primaryKey;autoIncrement" json:"idItem"`
	ItemCode    string `gorm:"not null;type:varchar(50)" json:"itemCode"`
	Description string `grom:"not null; type:TEXT" json:"description"`
	Quantity    uint   `gorm:"not null" json:"quantity"`
	OrderId     uint
}
