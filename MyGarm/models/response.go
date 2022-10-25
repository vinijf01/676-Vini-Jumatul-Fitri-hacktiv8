package models

import "time"

//USER
type ResponseRegister struct {
	Age      uint   `json:"age"`
	Email    string `json:"email"`
	IdUser   uint   `json:"id"`
	Username string `json:"username"`
}

func FormatResponseRegister(user User) ResponseRegister {
	formatter := ResponseRegister{
		Age:      user.Age,
		Email:    user.Email,
		IdUser:   user.IdUser,
		Username: user.Username,
	}
	return formatter
}

type ResponseUpdateUser struct {
	IdUser    uint      `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Age       uint      `json:"age"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FormatResponseUpdateUser(user User) ResponseUpdateUser {
	formatter := ResponseUpdateUser{
		IdUser:    user.IdUser,
		Email:     user.Email,
		Username:  user.Username,
		Age:       user.Age,
		UpdatedAt: user.UpdatedAt,
	}
	return formatter
}

//PHOTO
type ResponsePostPhoto struct {
	IdPhoto   uint      `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoURL  string    `json:"photo_url"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

func FormatResponsePostPhoto(photo Photo) ResponsePostPhoto {
	formatter := ResponsePostPhoto{
		IdPhoto:   photo.IdPhoto,
		Title:     photo.Title,
		Caption:   photo.Caption,
		PhotoURL:  photo.PhotoURL,
		UserID:    photo.UserID,
		CreatedAt: photo.CreatedAt,
	}
	return formatter
}

type ResponseGetPhoto struct {
	IdPhoto   uint      `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoURL  string    `json:"photo_url"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      DataUser  `json:"User"`
}

type DataUser struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

//satu photo
func FormatResponseGetPhoto(photo Photo, user DataUser) ResponseGetPhoto {
	formatter := ResponseGetPhoto{
		IdPhoto:   photo.IdPhoto,
		Title:     photo.Title,
		Caption:   photo.Caption,
		PhotoURL:  photo.PhotoURL,
		UserID:    photo.UserID,
		CreatedAt: photo.CreatedAt,
		UpdatedAt: photo.UpdatedAt,
		User: DataUser{
			Email:    user.Email,
			Username: user.Username,
		},
	}
	return formatter
}

//banyak photo
func FormatResponseGetPhotos(photos []Photo, user DataUser) []ResponseGetPhoto {
	photosFormatter := []ResponseGetPhoto{}

	for _, photo := range photos {
		responseGetPhoto := FormatResponseGetPhoto(photo, user)
		photosFormatter = append(photosFormatter, responseGetPhoto)
	}
	return photosFormatter
}

//COMMENT
type ResponseComment struct {
	IdComment uint      `json:"id"`
	Message   string    `json:"message"`
	PhotoId   uint      `json:"photo_id"`
	UserId    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

func FormatResponseComment(comment Comments) ResponseComment {
	formatter := ResponseComment{
		IdComment: comment.IdComment,
		Message:   comment.Message,
		PhotoId:   comment.PhotoID,
		UserId:    comment.UserID,
		CreatedAt: comment.CreatedAt,
	}
	return formatter
}

type UserComment struct {
	IdUser   uint   `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type PhotoComment struct {
	IdPhoto  uint   `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
	UserId   uint   `json:"user_id"`
}

type ResponseGetComment struct {
	IdComment uint         `json:"id"`
	Message   string       `json:"message"`
	PhotoId   uint         `json:"photo_id"`
	UserId    uint         `json:"user_id"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	User      UserComment  `json:"User"`
	Photo     PhotoComment `json:"Photo"`
}

func FormatResponseGetComment(comment Comments, user UserComment, photo PhotoComment) ResponseGetComment {
	formatter := ResponseGetComment{
		IdComment: comment.IdComment,
		Message:   comment.Message,
		PhotoId:   comment.PhotoID,
		UserId:    comment.UserID,
		CreatedAt: comment.CreatedAt,
		UpdatedAt: comment.UpdatedAt,
		User: UserComment{
			IdUser:   user.IdUser,
			Email:    user.Email,
			Username: user.Username,
		},
		Photo: PhotoComment{
			IdPhoto:  photo.IdPhoto,
			Title:    photo.Title,
			Caption:  photo.Caption,
			PhotoUrl: photo.PhotoUrl,
			UserId:   photo.UserId,
		},
	}
	return formatter
}

//banyak Comment
func FormatResponseGetComments(comments []Comments, user UserComment, photo PhotoComment) []ResponseGetComment {
	commentsFormatter := []ResponseGetComment{}

	for _, comment := range comments {
		responseGetComment := FormatResponseGetComment(comment, user, photo)
		commentsFormatter = append(commentsFormatter, responseGetComment)
	}
	return commentsFormatter
}

type ResponseUpdateComment struct {
	IdComment uint      `json:"id"`
	Message   string    `json:"message"`
	PhotoId   uint      `json:"photo_id"`
	UserId    uint      `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FormatResponseUpdateComment(comment Comments) ResponseUpdateComment {
	formatter := ResponseUpdateComment{
		IdComment: comment.IdComment,
		Message:   comment.Message,
		PhotoId:   comment.PhotoID,
		UserId:    comment.UserID,
		UpdatedAt: comment.UpdatedAt,
	}
	return formatter
}

//SOCIAL MEDIA
type ResponseCreateSosmed struct {
	IdSosmed       uint      `json:"id"`
	Name           string    `json:"name"`
	SocialMediaURL string    `json:"social_media_url"`
	UserId         uint      `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
}

func FormatResponseCreateSosmed(sosmed SocialMedia) ResponseCreateSosmed {
	formatter := ResponseCreateSosmed{
		IdSosmed:       sosmed.IdSosmed,
		Name:           sosmed.Name,
		SocialMediaURL: sosmed.SocialMediaURL,
		UserId:         sosmed.UserID,
		CreatedAt:      sosmed.CreatedAt,
	}
	return formatter
}

type ResponseGetSosmed struct {
	IdSosmed       uint        `json:"id"`
	Name           string      `json:"name"`
	SocialMediaURL string      `json:"social_media_url"`
	UserID         uint        `json:"user_id"`
	CreatedAt      time.Time   `json:"created_at"`
	UpdatedAt      time.Time   `json:"updated_at"`
	User           UserComment `json:"User"`
}

func FormatResponseGetSosmed(sosmed SocialMedia, user UserComment) ResponseGetSosmed {
	formatter := ResponseGetSosmed{
		IdSosmed:       sosmed.IdSosmed,
		Name:           sosmed.Name,
		SocialMediaURL: sosmed.SocialMediaURL,
		UserID:         sosmed.UserID,
		CreatedAt:      sosmed.CreatedAt,
		UpdatedAt:      sosmed.UpdatedAt,
		User: UserComment{
			IdUser:   user.IdUser,
			Email:    user.Email,
			Username: user.Username,
		},
	}
	return formatter
}

func FormatResponseGetSosmeds(sosmed []SocialMedia, user UserComment) []ResponseGetSosmed {
	sosmedsFormatter := []ResponseGetSosmed{}

	for _, sosmed := range sosmed {
		responseGetPhoto := FormatResponseGetSosmed(sosmed, user)
		sosmedsFormatter = append(sosmedsFormatter, responseGetPhoto)
	}
	return sosmedsFormatter
}
