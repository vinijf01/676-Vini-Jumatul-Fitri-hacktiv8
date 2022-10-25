package controllers

import (
	"log"
	"mygram/helpers"
	"mygram/models"
	"mygram/services"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func UpdateUser(c *gin.Context) {
	var req models.UpdateUserInput
	currentUser := c.MustGet("currentUser").(jwt.MapClaims)
	IdUser := uint(currentUser["id"].(float64))

	if err := c.ShouldBindJSON(&req); err != nil {
		errors := helpers.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helpers.APIResponse("Update Data User Failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	updateUser, err := services.UpdateUser(req, IdUser)
	if err != nil {
		errors := helpers.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helpers.APIResponse("Update Data User Failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := models.FormatResponseUpdateUser(updateUser)
	response := helpers.APIResponse("Update Data User Successfully", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}

func DeleteUser(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(jwt.MapClaims)
	IdUser := uint(currentUser["id"].(float64))

	err := DeleteFilePhoto(c, int(IdUser))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := services.DeleteUser(int(IdUser))
	if err != nil {
		errors := helpers.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helpers.APIResponse("Delete User Failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helpers.APIResponse("Delete Data User Success", http.StatusOK, "success", res)

	c.JSON(http.StatusOK, response)

}

func DeleteFilePhoto(c *gin.Context, id int) error {

	oldUrl, err := services.GetPhotoURLByIdUser(uint(id))
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
