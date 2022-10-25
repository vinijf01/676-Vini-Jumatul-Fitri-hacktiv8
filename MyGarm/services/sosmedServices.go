package services

import (
	"mygram/databases"
	"mygram/models"
	"time"
)

func CreateSosmed(input models.SosmedInput, userId uint) (models.SocialMedia, error) {
	sosmed := models.SocialMedia{}
	db := databases.GetDb()

	sosmed.Name = input.Name
	sosmed.SocialMediaURL = input.SocialMediaURL
	sosmed.UserID = userId
	sosmed.CreatedAt = time.Now()
	errdb := db.Create(&sosmed).Error
	if errdb != nil {
		return sosmed, errdb
	}
	return sosmed, nil

}

func GetSosmeds(idUser uint) ([]models.SocialMedia, error) {
	var sosmed []models.SocialMedia
	db := databases.GetDb()
	err := db.Where("user_id = ?", idUser).Find(&sosmed).Error
	if err != nil {
		return sosmed, err
	}
	return sosmed, nil
}

func UpdateSosmed(input models.SosmedInput, IdSosmed uint) (models.SocialMedia, error) {
	var sosmed models.SocialMedia
	db := databases.GetDb()

	err := db.Where("id_sosmed = ?", IdSosmed).Find(&sosmed).Error
	if err != nil {
		return sosmed, err
	}

	sosmed.Name = input.Name
	sosmed.SocialMediaURL = input.SocialMediaURL
	sosmed.UpdatedAt = time.Now()
	err = db.Save(&sosmed).Error
	if err != nil {
		return sosmed, err
	}
	return sosmed, nil

}

func DeleteSosmed(IdSosmed uint) (string, error) {
	var sosmed models.SocialMedia
	db := databases.GetDb()
	err := db.Model(&sosmed).Where("id_sosmed = ?", IdSosmed).Find(&sosmed).Error
	if err != nil {
		return "Data tidak ditemukan", err
	}

	err = db.Model(&sosmed).Where("id_sosmed = ?", IdSosmed).Delete(&sosmed).Error
	if err != nil {
		return "Comment Gagal dihapus", err
	}
	return "Your sosmed has been successfully deleted", nil

}

func GetSosmedById(IdUser uint) (models.SocialMedia, error) {
	var sosmed models.SocialMedia
	db := databases.GetDb()
	err := db.Where("user_id = ?", IdUser).Find(&sosmed).Error
	if err != nil {
		return sosmed, err
	}
	return sosmed, nil
}
