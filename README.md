# gin-gonic-spike
A spike to evaluate gin-gonic frame work to do rest API in go

Uses GORM for handling database as best practice https://github.com/jinzhu/gorm
Uses jwt middleware for authentication https://github.com/appleboy/gin-jwt
Uses air for hot reloading https://github.com/cosmtrek/air
Uses gin-swagger for docs https://github.com/swaggo/gin-swagger

need to generate docs

with the go utility swag https://github.com/swaggo/swag

`swag init` this places the static docs in the docs directory which is served on the /docs/index.html

TODO: use casbin for role based security https://github.com/gin-contrib/authz
TODO: do validation with https://github.com/go-playground/validator or https://github.com/xeipuuv/gojsonschema
