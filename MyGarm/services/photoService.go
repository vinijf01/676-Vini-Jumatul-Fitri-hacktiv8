package services

import (
	"errors"
	"mygram/databases"
	"mygram/models"
	"time"
)

func PostPhoto(input models.PhotoForm, IdUser int) (models.Photo, error) {
	photo := models.Photo{}
	db := databases.GetDb()

	photo.Title = input.Title
	photo.Caption = input.Caption
	photo.PhotoURL = input.PhotoURL
	photo.CreatedAt = time.Now()
	photo.UserID = uint(IdUser)

	errdb := db.Create(&photo).Error
	if errdb != nil {
		return photo, errdb
	}
	return photo, nil
}

func GetPhotos(idUser uint) ([]models.Photo, error) {
	var photos []models.Photo
	db := databases.GetDb()
	err := db.Where("user_id = ?", idUser).Find(&photos).Error
	if err != nil {
		return photos, err
	}
	return photos, nil
}

func UpdatePhoto(input models.UpdatePhotoForm, idPhoto uint) (models.Photo, error) {
	var photo models.Photo
	db := databases.GetDb()

	err := db.Where("id_photo = ?", idPhoto).Preload("Comments").Find(&photo).Error
	if err != nil {
		return photo, err
	}
	photo.Caption = input.Caption
	photo.Title = input.Title
	photo.PhotoURL = input.PhotoURL
	photo.UpdatedAt = time.Now()
	err = db.Save(&photo).Error
	if err != nil {
		return photo, err
	}
	return photo, nil
}

func DeletePhoto(idPhoto int) (string, error) {
	var photo models.Photo
	var comment models.Comments
	db := databases.GetDb()
	err := db.Where("id_photo = ?", idPhoto).Preload("Comments").Find(&photo).Error
	if err != nil {
		return "Data tidak ditemukan", err
	}
	err = db.Where("photo_id = ?", idPhoto).Delete(comment).Error
	if err != nil {
		return "Data comment Gagal dihapus", err
	}

	err = db.Where("id_photo = ?", idPhoto).Delete(photo).Error
	if err != nil {
		return "Data photo Gagal dihapus", err
	}
	return "message : Data photo berhasil di hapus", nil

}

func GetPhotosById(IdPhoto uint) (models.Photo, error) {
	var photo models.Photo
	db := databases.GetDb()
	err := db.Where("id_photo = ?", IdPhoto).Find(&photo).Error
	if err != nil {
		return photo, err
	}
	if photo.IdPhoto == 0 {
		return photo, errors.New("photo data not found")
	}
	return photo, nil
}

func GetPhotoURLById(IdPhoto uint) (*string, error) {
	var photo models.Photo
	db := databases.GetDb()
	err := db.Where("id_photo = ?", IdPhoto).Select("photo_url").Find(&photo).Error
	if err != nil {
		return &photo.PhotoURL, err
	}

	return &photo.PhotoURL, nil
}

func GetPhotoURLByIdUser(Iduser uint) (*string, error) {
	var photo models.Photo
	db := databases.GetDb()
	err := db.Where("user_id = ?", Iduser).Select("photo_url").Find(&photo).Error
	if err != nil {
		return &photo.PhotoURL, err
	}
	return &photo.PhotoURL, nil
}
