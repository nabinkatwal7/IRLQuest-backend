package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/nabinkatwal7/irlquest/controller"
	"github.com/nabinkatwal7/irlquest/middleware"
)

func ServeApplication(){
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/", func (c *gin.Context){
		c.JSON(200, gin.H{
			"message": "Server responded with status code 200",
		})
	})

	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/register", controller.Register)
	publicRoutes.POST("/login", controller.Login)
	publicRoutes.GET("/users", controller.GetAllUsers)

	protectedRoutes := router.Group("/api")
	protectedRoutes.Use(middleware.JWTAuthMiddleware())

	router.Run()
	fmt.Println("Server started on port 8080")
}