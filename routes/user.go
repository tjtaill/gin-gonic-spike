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

// @Summary get a list of users
// @Description get a list of users
// @Tags user
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {array} models.User
// @Router /api/v1/users [get]
func getuser(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		users := make([]models.User, 0)
		db.Find(&users)
		ctx.JSON(http.StatusOK, users)
	}
}

// @Summary get a list of users
// @Description get a list of users
// @Tags user
// @Accept  json
// @Param credentials body models.User true "user"
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {object} middleware.Credentials
// @Router /api/v1/user [post]
func postuser(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := models.User{}
		err := ctx.BindJSON(&user)
		if err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		hashedPwd, err := hashAndSalt(user.Password)
		if err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		user.Password = hashedPwd
		errs := db.Create(&user).GetErrors()
		if len(errs) > 0 {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": errs[0].Error()})
			return
		}
	}
}

func userRoutes(router *gin.RouterGroup, db *gorm.DB) {
	router.GET("/users", getuser(db))
	router.POST("/user", postuser(db))
}
