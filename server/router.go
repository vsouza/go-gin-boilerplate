package server

import (
	"github.com/gin-gonic/gin"
	"github.com/vsouza/go-gin-boilerplate/controllers"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(controllers.HealthController)
	router.GET("/health", health.Status)
	router.Use(AuthMiddleware())
	v1 := router.Group("v1")
	{
		userGroup := v1.Group("user")
		{
			user := new(controllers.UserController)
			userGroup.GET("/:id", user.Retrieve)
			userGroup.POST("/", user.Signup)
			userGroup.DELETE("/:id", user.Delete)
			userGroup.PUT("/:id", user.Update)
		}
	}
	return router
}
