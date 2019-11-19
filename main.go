package main

import (
	"log"

	"github.com/ElementAI/gin-gonic-spike/models"
	"github.com/ElementAI/gin-gonic-spike/routes"
	"github.com/gin-gonic/gin"
)

// @title Spike API
// @version 1.0
// @description This is a simple API to show how to use gin
// @contact.name Troy Taillefer
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	r := gin.Default()
	db, err := models.Register()
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		db.Close()
	}()
	routes.Register(r, db)
	r.Run(":8080")
}
