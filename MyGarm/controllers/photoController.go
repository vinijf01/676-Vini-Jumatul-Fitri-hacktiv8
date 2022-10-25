package controllers

import (
	"log"
	"mygram/helpers"
	"mygram/models"
	"mygram/services"
	"net/http"
	"os"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func PostPhoto(c *gin.Context) {
	var photo models.PhotoForm

	currentUser := c.MustGet("currentUser").(jwt.MapClaims)
	IdUser := int(currentUser["id"].(float64))
	userID := strconv.Itoa(IdUser)

	file, err := c.FormFile("photo_url")
	if err != nil {
		err := "photo_url cannot be empty"
		errorMessage := gin.H{"errors": err}

		response := helpers.APIResponse("Post Photo Failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	log.Println(file.Filename)

	fileName := userID + "." + file.Filename

	err = c.SaveUploadedFile(file, "assets/img/"+fileName)
	if err != nil {
		log.Fatal(err)
	}

	photo.Title = c.Request.FormValue("title")
	if photo.Title == "" {
		err := "Title cannot be empty"
		errorMessage := gin.H{"errors": err}

		response := helpers.APIResponse("Post Photo Failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	photo.Caption = c.Request.FormValue("caption")
	photo.PhotoURL = fileName

	newPhoto, err := services.PostPhoto(photo, IdUser)
	if err != nil {
		errors := helpers.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helpers.APIResponse("Post Photo Failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := models.FormatResponsePostPhoto(newPhoto)
	response := helpers.APIResponse("Post Photo Successfully", http.StatusCreated, "success", formatter)

	c.JSON(http.StatusCreated, response)

}

func GetPhotos(c *gin.Context) {
	var photos []models.Photo
	var datauser models.DataUser
	currentUser := c.MustGet("currentUser").(jwt.MapClaims)
	IdUser := uint(currentUser["id"].(float64))
	user, _ := services.GetUserByID(IdUser)

	datauser.Email = user.Email
	datauser.Username = user.Username
	log.Println(datauser.Email)
	log.Println(datauser.Username)

	photos, err := services.GetPhotos(IdUser)
	if err != nil {
		errors := helpers.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helpers.APIResponse("Get Photo Failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := models.FormatResponseGetPhotos(photos, datauser)
	response := helpers.APIResponse("Get Photo Successfully", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func UpdatePhoto(c *gin.Context) {
	var photo models.UpdatePhotoForm

	currentUser := c.MustGet("currentUser").(jwt.MapClaims)
	IdUser := int(currentUser["id"].(float64))
	userID := strconv.Itoa(IdUser)

	strid := c.Param("photoId")
	idPhoto, _ := strconv.Atoi(strid)
	log.Println(idPhoto)
	oldUrl, err := services.GetPhotoURLById(uint(idPhoto))
	if err != nil {
		log.Fatal(err)
	}

	file, err := c.FormFile("photo_url")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(file.Filename)

	if file.Filename != "" {
		err := DeleteURLPhoto(c, idPhoto)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fileName := userID + "." + file.Filename
		err = c.SaveUploadedFile(file, "assets/img/"+fileName)
		if err != nil {
			log.Fatal(err)
		}
		photo.PhotoURL = fileName

	} else {
		photo.PhotoURL = *oldUrl
	}

	photo.Title = c.Request.FormValue("title")
	photo.Caption = c.Request.FormValue("caption")

	newPhoto, err := services.UpdatePhoto(photo, uint(idPhoto))
	if err != nil {
		errors := helpers.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helpers.APIResponse("Update Photo Failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := models.FormatResponsePostPhoto(newPhoto)
	response := helpers.APIResponse("Update Photo Successfully", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}

func DeletePhoto(c *gin.Context) {
	strid := c.Param("photoId")
	idPhoto, _ := strconv.Atoi(strid)

	currentUser := c.MustGet("currentUser").(jwt.MapClaims)
	IdUser := uint(currentUser["id"].(float64))

	photo, err := services.GetPhotosById(uint(idPhoto))
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helpers.APIResponse("Delete Photo Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if photo.UserID != IdUser {
		errors := helpers.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helpers.APIResponse("Delete Photo Failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err = DeleteURLPhoto(c, idPhoto)
	if err != nil {
		errorMessage := gin.H{"errors": err}

		response := helpers.APIResponse("Delete File Photo Failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	res, err := services.DeletePhoto(idPhoto)
	if err != nil {
		errorMessage := gin.H{"errors": err}

		response := helpers.APIResponse("Delete Photo Failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	response := helpers.APIResponse("Delete Photo Successfully", http.StatusOK, "success", res)

	c.JSON(http.StatusOK, response)
}

func DeleteURLPhoto(c *gin.Context, id int) error {

	oldUrl, err := services.GetPhotoURLById(uint(id))
	if err != nil {
		log.Fatal(err)
	}

	filepath := "assets/img/" + *oldUrl

	if *oldUrl != "" {
		if err := os.Remove(filepath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			panic(err)
		}
		return err
	}
	return nil
}
