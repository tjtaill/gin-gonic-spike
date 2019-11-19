package routes

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
)

func badPassword(hashedPwd string, plainPwd string) bool {
	byteHashed := []byte(hashedPwd)
	bytePlain := []byte(plainPwd)
	err := bcrypt.CompareHashAndPassword(byteHashed, bytePlain)
	if err != nil {
		return true
	}
	return false
}

func loginRoutes(router *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) {
	router.POST("/login", authMiddleware.LoginHandler)
}
