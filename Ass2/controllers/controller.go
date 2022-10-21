package controllers

import (
	"ass2/models"
	"ass2/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	var req models.Request
	var res models.Response
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	res, err := services.SaveOrder(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	c.JSON(http.StatusOK, res)
}

func GetOrder(c *gin.Context) {
	var order []models.Order
	order, err := services.GetOrder()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if len(order) > 0 {
		c.JSON(http.StatusOK, order)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": "Data Kosong",
		})
	}

}

func GetOrderbyId(c *gin.Context) {
	strid := c.Param("idOrder")
	id, _ := strconv.Atoi(strid)
	order, err := services.GetOrderByID(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, order)
}

func UpdateOrder(c *gin.Context) {
	strid := c.Param("idOrder")
	id, _ := strconv.Atoi(strid)
	var req models.Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	res, err := services.UpdateOrder(req, uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	c.JSON(http.StatusOK, res)
}

func DeleteOrder(c *gin.Context) {
	strid := c.Param("idOrder")
	id, _ := strconv.Atoi(strid)
	res, err := services.DeleteOrder(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	c.JSON(http.StatusOK, res)
}
