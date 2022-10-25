package models

import "time"

type SocialMedia struct {
	IdSosmed       uint   `gorm:"primaryKey;autoIncrement" json:"idSosmed"`
	Name           string `gorm:"not null" json:"name" valid:"required"`
	SocialMediaURL string `gorm:"not null" json:"social_media_url" valid:"required"`
	UserID         uint
	CreatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updatedAt"`
}
