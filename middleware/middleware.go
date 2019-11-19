package middleware

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/jinzhu/gorm"
)

func Auth(db *gorm.DB) (*jwt.GinJWTMiddleware, error) {
	return auth(db)
}
