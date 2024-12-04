package middlewares

import (
	"go-postgres-auth0/utils"
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
)

func Authenticate(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")

	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": " Route not authorized!"})
		return
	}

	parts := strings.Split(token, " ")
    if len(parts) != 2 || parts[0] != "Bearer" {
        ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid token format!"})
        return
    }

    actualToken := parts[1]

	userId, err := utils.VerifyToken(actualToken)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized. Invalid token!"})
		return
	}

	ctx.Set("userId", userId)
	ctx.Next()
}