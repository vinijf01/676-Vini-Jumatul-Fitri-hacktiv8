package services

import (
	"errors"
	"mygram/auth"
	"mygram/databases"
	"mygram/models"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(input models.RegisterUserInput) (models.User, error) {
	user := models.User{}
	db := databases.GetDb()
	if input.Age <= 8 {
		return user, errors.New("make sure you are over 8 years old")
	}

	user.Age = input.Age
	user.Email = input.Email
	passcheck := len(input.Password)

	if passcheck < 6 {
		return user, errors.New("password must be at least 6 characters long")
	}

	passwordhash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.Password = string(passwordhash)
	user.Username = input.Username
	user.CreatedAt = time.Now()
	errdb := db.Create(&user).Error
	if errdb != nil {
		return user, errdb
	}

	return user, nil
}

func Login(input models.LoginInput) (string, error) {
	db := databases.GetDb()
	var user models.User
	//mencari user dg email yang diinputkan
	err := db.Where("email = ?", input.Email).Find(&user).Error
	if err != nil {
		return "user not found", err
	}

	if user.IdUser == 0 {
		return "error: ", errors.New("no user found on that email")
	}

	//mencocokan password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return "password not match", err
	}

	token, _ := auth.GenerateToken(user.IdUser, user.Email)
	return token, err

}

func UpdateUser(req models.UpdateUserInput, IdUser uint) (models.User, error) {
	var user models.User
	db := databases.GetDb()

	err := db.Where("id_user = ?", IdUser).Find(&user).Error
	if err != nil {
		return user, err
	}

	user.Username = req.Username
	user.Email = req.Email
	user.UpdatedAt = time.Now()

	err = db.Save(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil

}

func DeleteUser(idUser int) (string, error) {
	var (
		user    models.User
		photo   models.Photo
		comment models.Comments
		sosmed  models.SocialMedia
	)
	db := databases.GetDb()
	err := db.Where("id_user= ?", idUser).Preload("Photos").Preload("Comments").Preload("SocialMedias").Find(&user).Error
	if err != nil {
		return "Data not found", err
	}

	err = db.Where("user_id = ?", idUser).Delete(comment).Error
	if err != nil {
		return "Comment data failed to delete", err
	}

	err = db.Where("user_id = ?", idUser).Delete(sosmed).Error
	if err != nil {
		return "Sosmed data failed to delete", err
	}

	err = db.Where("user_id = ?", idUser).Delete(photo).Error
	if err != nil {
		return "Photo data failed to delete", err
	}

	err = db.Where("id_user = ?", idUser).Delete(user).Error
	if err != nil {
		return "User data failed to delete", err
	}
	return "message : User data has been deleted successfully", nil

}

func GetUserByID(IdUser uint) (models.User, error) {
	var user models.User
	db := databases.GetDb()

	err := db.Where("id_user = ?", IdUser).Preload("Photos").Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
