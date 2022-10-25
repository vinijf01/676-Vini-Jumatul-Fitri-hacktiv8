package models

import "time"

type User struct {
	IdUser       uint          `gorm:"primaryKey;autoIncrement" json:"idUser"`
	Username     string        `gorm:"not null;unique;type:varchar(100)" json:"username" valid:"required"`
	Email        string        `gorm:"not null;unique" json:"email" valid:"email,required"`
	Password     string        `gorm:"not null" json:"password" valid:"required"`
	Age          uint          `gorm:"not null" json:"age" valid:"required"`
	CreatedAt    time.Time     `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt    time.Time     `gorm:"default:CURRENT_TIMESTAMP" json:"updatedAt"`
	Photos       []Photo       `grom:"constraint:OnUpdate:CASCADE,OnDelete: SET NULL" json:"photos"`
	Comments     []Comments    `grom:"constraint:OnUpdate:CASCADE,OnDelete: SET NULL" json:"comments"`
	SocialMedias []SocialMedia `grom:"constraint:OnUpdate:CASCADE,OnDelete: SET NULL" json:"socialMedia"`
}
