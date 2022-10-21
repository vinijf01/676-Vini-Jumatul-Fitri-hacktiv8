package routers

import (
	"ass2/controllers"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/orders", controllers.CreateOrder)
	router.GET("/orders", controllers.GetOrder)
	router.GET("/orders/:idOrder", controllers.GetOrderbyId)
	router.PUT("/orders/:idOrder", controllers.UpdateOrder)
	router.DELETE("/orders/:idOrder", controllers.DeleteOrder)

	return router
}
