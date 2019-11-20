package routes

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
)

type TokenResponse struct {
	Code   int    `json:"code"`
	Expire string `json:"expire"`
	Token  string `json:"token"`
}

func badPassword(hashedPwd string, plainPwd string) bool {
	byteHashed := []byte(hashedPwd)
	bytePlain := []byte(plainPwd)
	err := bcrypt.CompareHashAndPassword(byteHashed, bytePlain)
	if err != nil {
		return true
	}
	return false
}

// @Summary login
// @Description login
// @Tags login
// @Accept  json
// @Param credentials body middleware.Credentials true "credentials"
// @Produce  json
// @Success 200 {object} routes.TokenResponse
// @Router /login [post]
func loginRoutes(router *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) {
	router.POST("/login", authMiddleware.LoginHandler)
}
