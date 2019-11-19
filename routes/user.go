package routes

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/ElementAI/gin-gonic-spike/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func hashAndSalt(pwd string) (string, error) {
	bs := []byte(pwd)
	hash, err := bcrypt.GenerateFromPassword(bs, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func userRoutes(router *gin.RouterGroup, db *gorm.DB) {
	router.GET("/users", func(ctx *gin.Context) {
		users := make([]models.User, 0)
		db.Find(&users)
		ctx.JSON(http.StatusOK, users)
	})

	router.POST("/user", func(ctx *gin.Context) {
		user := models.User{}
		err := ctx.BindJSON(&user)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		hashedPwd, err := hashAndSalt(user.Password)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		user.Password = hashedPwd
		errs := db.Create(&user).GetErrors()
		if len(errs) > 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": errs[0].Error()})
			return
		}
	})
}
