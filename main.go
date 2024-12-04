package main

import (
	"fmt"
	"go-postgres-auth0/database"
	"go-postgres-auth0/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error while loading environment!")
	}

	router := gin.New()

	database.ConnectToPostgresDB()

	routes.UserRoutes(router)
	routes.AuthRoutes(router)

	router.Run(":8000")

	fmt.Println("Server started successfully!")
}
