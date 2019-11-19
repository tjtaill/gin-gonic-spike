package routes

import (
	"github.com/ElementAI/gin-gonic-spike/middleware"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Register(router *gin.Engine, db *gorm.DB) error {
	authhMiddleware, err := middleware.Auth(db)
	if err != nil {
		return err
	}
	loginRoutes(router, authhMiddleware)
	api := router.Group("/api/v1")
	api.Use(authhMiddleware.MiddlewareFunc())
	userRoutes(api, db)
	return nil
}
