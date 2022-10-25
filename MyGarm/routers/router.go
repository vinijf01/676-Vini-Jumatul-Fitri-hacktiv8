package routers

import (
	"mygram/controllers"
	"mygram/middlewares"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {

	router := gin.Default()

	authRouter := router.Group("/users")
	{
		authRouter.POST("/register", controllers.RegisterUser)
		authRouter.POST("/login", controllers.Login)
	}

	userRouter := router.Group("/users")
	{
		userRouter.Use(middlewares.JWTAuth())
		userRouter.PUT("/", controllers.UpdateUser)
		userRouter.DELETE("/", controllers.DeleteUser)
	}

	photoRouter := router.Group("/photos")
	{
		photoRouter.Use(middlewares.JWTAuth())
		photoRouter.POST("/", controllers.PostPhoto)
		photoRouter.GET("/", controllers.GetPhotos)
		photoRouter.PUT("/:photoId", controllers.UpdatePhoto)
		photoRouter.DELETE("/:photoId", controllers.DeletePhoto)
	}

	commentRouter := router.Group("/comments")
	{
		commentRouter.Use(middlewares.JWTAuth())
		commentRouter.POST("/", controllers.PostComment)
		commentRouter.GET("/", controllers.GetComments)
		commentRouter.PUT("/:commentId", controllers.UpdateComment)
		commentRouter.DELETE("/:commentId", controllers.DeleteComment)

	}
	sosmedRouter := router.Group("/socialmedias")
	{
		sosmedRouter.Use(middlewares.JWTAuth())
		sosmedRouter.POST("/", controllers.CreateSosmed)
		sosmedRouter.GET("/", controllers.GetSosmeds)
		sosmedRouter.PUT("/:socialMediaId", controllers.UpdateSosmed)
		sosmedRouter.DELETE("/:socialMediaId", controllers.DeleteSosmed)

	}

	return router
}
