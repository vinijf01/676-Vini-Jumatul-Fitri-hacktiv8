package models

import "time"

type Order struct {
	IdOrder      uint      `gorm:"primaryKey;autoIncrement" json:"idOrder"`
	CustomerName string    `gorm:"not null; type:varchar(50)" json:"customerName"`
	OrderedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"orderedAt"`
	Items        []Items   `grom:"constraint:OnUpdate:CASCADE,OnDelete: SET NULL" json:"items"`
}
