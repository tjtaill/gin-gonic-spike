package middleware

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Create(db *gorm.DB) (*jwt.GinJWTMiddleware, *gin.HandlerFunc, error) {
	authMiddleware, err := auth(db)
	if err != nil {
		return nil, nil, err
	}
	rbacMiddleware, err := rbac(db)
	if err != nil {
		return nil, nil, err
	}
	return authMiddleware, rbacMiddleware, nil
}
