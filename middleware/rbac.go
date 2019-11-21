package middleware

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v2"
	"github.com/gin-contrib/authz"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func rbac(db *gorm.DB) (*gin.HandlerFunc, error) {
	adapter, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		return nil, err
	}
	enforcer, err := casbin.NewEnforcer("rbac_model.conf", adapter)
	if err != nil {
		return nil, err
	}
	err = enforcer.LoadPolicy()
	if err != nil {
		return nil, err
	}
	enforcer.AddPolicy("xyz", "/api/v1/users", "GET")
	enforcer.SavePolicy()
	authorizer := authz.NewAuthorizer(enforcer)
	return &authorizer, nil
}
