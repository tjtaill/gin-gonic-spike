package main

import (
	"github.com/ElementAI/gin-gonic-spike/models"
	"github.com/ElementAI/gin-gonic-spike/routes"
	"github.com/gin-gonic/gin"
	"log"
)

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
