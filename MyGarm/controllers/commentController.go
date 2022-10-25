package controllers

import (
	"mygram/helpers"
	"mygram/models"
	"mygram/services"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func PostComment(c *gin.Context) {
	var input models.CommentInput
	currentUser := c.MustGet("currentUser").(jwt.MapClaims)
	IdUser := int(currentUser["id"].(float64))

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helpers.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helpers.APIResponse("Post Comment Failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newComment, err := services.PostComment(input, uint(IdUser))
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helpers.APIResponse("Post Comment Failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := models.FormatResponseComment(newComment)
	response := helpers.APIResponse("Post Comment Success", http.StatusCreated, "success", formatter)

	c.JSON(http.StatusCreated, response)
}

func GetComments(c *gin.Context) {
	var comments []models.Comments
	var datauser models.UserComment
	var dataphoto models.PhotoComment

	currentUser := c.MustGet("currentUser").(jwt.MapClaims)
	IdUser := uint(currentUser["id"].(float64))

	user, _ := services.GetUserByID(IdUser)
	photo, _ := services.GetPhotosById(IdUser)

	datauser.IdUser = user.IdUser
	datauser.Email = user.Email
	datauser.Username = user.Username

	dataphoto.IdPhoto = photo.IdPhoto
	dataphoto.Title = photo.Title
	dataphoto.Caption = photo.Caption
	dataphoto.PhotoUrl = photo.PhotoURL
	dataphoto.UserId = photo.UserID

	comments, err := services.GetComments(IdUser)
	if err != nil {
		errors := helpers.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helpers.APIResponse("Get Comment Failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := models.FormatResponseGetComments(comments, datauser, dataphoto)
	response := helpers.APIResponse("post Photo success", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func UpdateComment(c *gin.Context) {
	var input models.UpdateCommentInput
	strid := c.Param("commentId")
	idComment, _ := strconv.Atoi(strid)
	currentUser := c.MustGet("currentUser").(jwt.MapClaims)
	IdUser := uint(currentUser["id"].(float64))

	comment, err := services.GetCommentById(IdUser)
	if err != nil {

		c.JSON(http.StatusBadRequest, err)
		return
	}

	if comment.UserID != IdUser {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		errors := helpers.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helpers.APIResponse("Update Comment Failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	updateComment, err := services.UpdateComment(input, uint(idComment))
	if err != nil {
		errors := helpers.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helpers.APIResponse("Update Comment Failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := models.FormatResponseUpdateComment(updateComment)
	response := helpers.APIResponse("Update Comment Success", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func DeleteComment(c *gin.Context) {
	strid := c.Param("commentId")
	idComment, _ := strconv.Atoi(strid)
	currentUser := c.MustGet("currentUser").(jwt.MapClaims)
	IdUser := uint(currentUser["id"].(float64))

	comment, err := services.GetCommentById(IdUser)
	if err != nil {

		errors := helpers.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helpers.APIResponse("Delete Comment Failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if comment.UserID != IdUser {
		errors := helpers.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helpers.APIResponse("Delete Comment Failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	res, err := services.DeleteComment(uint(idComment))
	if err != nil {
		errors := helpers.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helpers.APIResponse("Delete Comment Failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helpers.APIResponse("Delete Comment Success", http.StatusOK, "success", res)

	c.JSON(http.StatusOK, response)
}
