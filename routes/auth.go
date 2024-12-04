package routes

import (
	"go-postgres-auth0/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(server *gin.Engine) {
	server.POST("/login", controllers.Login)
	server.POST("/register", controllers.Register)
}
