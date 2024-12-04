package routes

import (
	"go-postgres-auth0/controllers"
	"go-postgres-auth0/middlewares"
	"github.com/gin-gonic/gin"
)

func UserRoutes(server *gin.Engine) {

	authenticated := server.Group("/")

	authenticated.Use(middlewares.Authenticate)

	server.GET("/allUsers", controllers.GetAllUsers)
	authenticated.GET("/user/:id", controllers.GetUser)
	authenticated.POST("/user", controllers.CreateUser)
	authenticated.DELETE("/user/:id", controllers.DeleteUser)
	authenticated.PUT("/user/:id", controllers.UpdateUser)
}
