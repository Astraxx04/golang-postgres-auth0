package controllers

import (
	"net/http"
	"go-postgres-auth0/database"
	"go-postgres-auth0/models"
	"github.com/gin-gonic/gin"
	"go-postgres-auth0/utils"
)

func Login(ctx *gin.Context) {
	var loginRequest struct {
		Email string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	var user models.User

	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H {
			"message": "Invalid request body!",
			"error": err.Error(),
		})
		return
	}

	if err := database.DB.Where("email=?", loginRequest.Email).First(&user).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H {
			"message": "Invalid email!",
			"error": err.Error(),
		})
		return
	}

	if !utils.CheckPasswordHash(loginRequest.Password, user.Password) {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H {
			"message": "Invalid password!",
		})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if(err != nil) {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Could not authenticate user!",
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H {
		"message": "Login successfull!",
		"token": token,
	})
}

func Register(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H {
			"message": "Failed to create user. Invalid request body!",
			"error": err.Error(),
		})
		return
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to hash password!",
			"error":   err.Error(),
		})
		return
	}

	user.Password = hashedPassword

	if err := database.DB.Save(&user).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H {
			"message": "Could not register user!",
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully!",
		"response": &user,
	})
}