package services

import (
	"mygram/databases"
	"mygram/models"
	"time"
)

func PostComment(input models.CommentInput, userId uint) (models.Comments, error) {
	comment := models.Comments{}
	db := databases.GetDb()

	comment.Message = input.Message
	comment.PhotoID = input.PhotoID
	comment.UserID = userId
	comment.CreatedAt = time.Now()
	errdb := db.Create(&comment).Error
	if errdb != nil {
		return comment, errdb
	}
	return comment, nil

}

func GetComments(idUser uint) ([]models.Comments, error) {
	var comments []models.Comments
	db := databases.GetDb()
	err := db.Where("user_id = ?", idUser).Find(&comments).Error
	if err != nil {
		return comments, err
	}
	return comments, nil
}

func UpdateComment(input models.UpdateCommentInput, idComment uint) (models.Comments, error) {
	var comment models.Comments
	db := databases.GetDb()

	err := db.Where("id_comment = ?", idComment).Find(&comment).Error
	if err != nil {
		return comment, err
	}

	comment.Message = input.Message
	comment.UpdatedAt = time.Now()
	err = db.Save(&comment).Error
	if err != nil {
		return comment, err
	}
	return comment, nil

}

func DeleteComment(IdComment uint) (string, error) {
	var comment models.Comments
	db := databases.GetDb()
	err := db.Model(&comment).Where("id_comment = ?", IdComment).Find(&comment).Error
	if err != nil {
		return "Data tidak ditemukan", err
	}

	err = db.Model(&comment).Where("id_comment = ?", IdComment).Delete(&comment).Error
	if err != nil {
		return "Comment Gagal dihapus", err
	}
	return "Your comment has been successfully deleted", nil

}

func GetCommentById(IdUser uint) (models.Comments, error) {
	var comment models.Comments
	db := databases.GetDb()
	err := db.Where("user_id = ?", IdUser).Find(&comment).Error
	if err != nil {
		return comment, err
	}
	return comment, nil
}
