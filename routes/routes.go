package routes

import (
	_ "github.com/ElementAI/gin-gonic-spike/docs"
	"github.com/ElementAI/gin-gonic-spike/middleware"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Register(router *gin.Engine, db *gorm.DB) error {
	authMiddleware, rbacMiddleware, err := middleware.Create(db)
	if err != nil {
		return err
	}
	url := ginSwagger.URL("http://localhost:8080/docs/doc.json")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	loginRoutes(router, authMiddleware)
	api := router.Group("/api/v1")
	api.Use(authMiddleware.MiddlewareFunc())
	api.Use(*rbacMiddleware)
	userRoutes(api, db)
	return nil
}
