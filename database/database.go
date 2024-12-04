package database

import (
	"fmt"
	"go-postgres-auth0/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToPostgresDB() {
	//dsn := "user=gopostgres password=gopostgres host=localhost port=5432 dbname=gopostgres sslmode=disable"
	uri := "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.User{})

	DB = db
	fmt.Println("Connected to postgres database successfully")
}
