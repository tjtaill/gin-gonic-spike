package middleware

import (
	"time"

	"github.com/ElementAI/gin-gonic-spike/models"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

const identityKey = "id"

type Credentials struct {
	Username string `json: "username" binding:"required" validate:"empty=false"`
	Password string `json: "password" binding:"required validate:"empty=false""`
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

func login(db *gorm.DB) func(*gin.Context) (interface{}, error) {
	return func(ctx *gin.Context) (interface{}, error) {
		var credPayload Credentials
		err := ctx.ShouldBind(&credPayload)
		if err != nil {
			return "", jwt.ErrMissingLoginValues
		}
		dbUser := models.User{}
		errs := db.Where("name = ?", credPayload.Username).Find(&dbUser).GetErrors()
		if len(errs) > 0 {
			return nil, jwt.ErrFailedAuthentication
		}

		if badPassword(dbUser.Password, credPayload.Password) {
			return nil, jwt.ErrFailedAuthentication
		}
		return &dbUser, nil
	}
}

func permCheck(data interface{}, c *gin.Context) bool {
	if _, ok := data.(*models.User); ok {
		return true
	}
	return false
}

func unauthorized(ctx *gin.Context, code int, message string) {
	ctx.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}

func identity(ctx *gin.Context) interface{} {
	claims := jwt.ExtractClaims(ctx)
	user := &models.User{Name: claims[identityKey].(string)}
	ctx.Request.SetBasicAuth(user.Name, "")
	return &models.User{Name: user.Name}
}

func payload(data interface{}) jwt.MapClaims {
	if user, ok := data.(*models.User); ok {
		return jwt.MapClaims{identityKey: user.Name}
	}
	return jwt.MapClaims{}
}

func auth(db *gorm.DB) (*jwt.GinJWTMiddleware, error) {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:                 "spike",
		SigningAlgorithm:      "",
		Key:                   []byte("not secret"), // TODO: read from ENV, File or Store
		Timeout:               time.Hour,
		MaxRefresh:            time.Hour,
		Authenticator:         login(db),
		Authorizator:          permCheck,
		PayloadFunc:           payload,
		Unauthorized:          unauthorized,
		LoginResponse:         nil,
		RefreshResponse:       nil,
		IdentityHandler:       identity,
		IdentityKey:           identityKey,
		TokenLookup:           "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:         "Bearer",
		TimeFunc:              time.Now,
		HTTPStatusMessageFunc: nil,
		PrivKeyFile:           "",
		PubKeyFile:            "",
		SendCookie:            false,
		SecureCookie:          false,
		CookieHTTPOnly:        false,
		CookieDomain:          "",
		SendAuthorization:     false,
		DisabledAbort:         false,
		CookieName:            "",
	})
	if err != nil {
		return nil, err
	}
	return authMiddleware, nil
}
