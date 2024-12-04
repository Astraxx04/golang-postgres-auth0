package controllers

import (
	"net/http"
	"go-postgres-auth0/database"
	"go-postgres-auth0/models"
	"github.com/gin-gonic/gin"
)

func GetAllUsers(ctx *gin.Context) {
	users := []models.User{}
	if err := database.DB.Find(&users).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H {
			"message": "Failed to get users!",
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H {
		"message": "Fetched all users successfully!",
		"response": &users,
	})
}

func GetUser(ctx *gin.Context) {
	var user models.User
	if err := database.DB.Where("id=?", ctx.Param("id")).First(&user).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H {
			"message": "Failed to get the user!",
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H {
		"message": "Fetched user successfully!",
		"response": &user,
	})
}

func CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H {
			"message": "Failed to create user. Invalid request body!",
			"error": err.Error(),
		})
		return
	}
	if err := database.DB.Create(&user).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusConflict, gin.H {
			"message": "Failed to create the user. User might already exist!",
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H {
		"message": "User created successfully!",
		"response": &user,
	})
}

func DeleteUser(ctx *gin.Context) {
	var user models.User
	if err := database.DB.Where("id=?", ctx.Param("id")).Delete(&user).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H {
			"message": "Failed to delete user. User does not exist!",
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H {
		"message": "User deleted successfully!",
		"response": &user,
	})
}

func UpdateUser(ctx *gin.Context) {
	var user models.User
	if err := database.DB.Where("id=?", ctx.Param("id")).First(&user).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H {
			"message": "Failed to update user. Could not find the requested user!",
			"error": err.Error(),
		})
		return
	}
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H {
			"message": "Failed to update user. Invalid request body!",
			"error": err.Error(),
		})
		return
	}
	if err := database.DB.Save(&user).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H {
			"message": "Could not update user. Saving to db failed!",
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H {
		"message": "User updated successfully!",
		"response": &user,
	})
}