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

func CreateSosmed(c *gin.Context) {
	var input models.SosmedInput
	currentUser := c.MustGet("currentUser").(jwt.MapClaims)
	IdUser := int(currentUser["id"].(float64))

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helpers.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helpers.APIResponse("Create Social Media Failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newSosmed, err := services.CreateSosmed(input, uint(IdUser))
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helpers.APIResponse("Create Social Media Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := models.FormatResponseCreateSosmed(newSosmed)
	response := helpers.APIResponse("Create Social Media Successfully", http.StatusCreated, "success", formatter)

	c.JSON(http.StatusCreated, response)
}

func GetSosmeds(c *gin.Context) {
	var sosmeds []models.SocialMedia
	var datauser models.UserComment

	currentUser := c.MustGet("currentUser").(jwt.MapClaims)
	IdUser := uint(currentUser["id"].(float64))

	user, _ := services.GetUserByID(IdUser)

	datauser.IdUser = user.IdUser
	datauser.Email = user.Email
	datauser.Username = user.Username

	sosmeds, err := services.GetSosmeds(IdUser)
	if err != nil {
		errors := helpers.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helpers.APIResponse("Get Social Media Failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := models.FormatResponseGetSosmeds(sosmeds, datauser)
	response := helpers.APIResponse("Get Social Media Successfully", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func UpdateSosmed(c *gin.Context) {
	var input models.SosmedInput
	strid := c.Param("socialMediaId")
	idSosmed, _ := strconv.Atoi(strid)
	currentUser := c.MustGet("currentUser").(jwt.MapClaims)
	IdUser := uint(currentUser["id"].(float64))

	sosmed, err := services.GetSosmedById(IdUser)
	if err != nil {

		c.JSON(http.StatusBadRequest, err)
		return
	}

	if sosmed.UserID != IdUser {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		errors := helpers.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helpers.APIResponse("Update Social Media Failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	updateSosmed, err := services.UpdateSosmed(input, uint(idSosmed))
	if err != nil {
		errors := helpers.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helpers.APIResponse("Update Social Media Failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := models.FormatResponseCreateSosmed(updateSosmed)
	response := helpers.APIResponse("Update Social Media Successfully", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func DeleteSosmed(c *gin.Context) {
	strid := c.Param("socialMediaId")
	idComment, _ := strconv.Atoi(strid)
	currentUser := c.MustGet("currentUser").(jwt.MapClaims)
	IdUser := uint(currentUser["id"].(float64))

	comment, err := services.GetSosmedById(IdUser)
	if err != nil {

		c.JSON(http.StatusBadRequest, err)
		return
	}

	if comment.UserID != IdUser {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	res, err := services.DeleteSosmed(uint(idComment))
	if err != nil {
		errors := helpers.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helpers.APIResponse("Delete Social Media Failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.APIResponse("Delete Social Media Successfully", http.StatusOK, "success", res)

	c.JSON(http.StatusOK, response)
}
