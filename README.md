# gin-gonic-spike
A spike to evaluate gin-gonic frame work to do rest API in go

# Technology usage 
* Uses GORM for handling database as best practice https://github.com/jinzhu/gorm
* Uses jwt middleware for authentication https://github.com/appleboy/gin-jwt
* Uses air for hot reloading https://github.com/cosmtrek/air
* Uses gin-swagger for docs https://github.com/swaggo/gin-swagger
* Uses validator for validation https://github.com/go-playground/validator

# Generate Docs
to generate docs

with the go utility swag https://github.com/swaggo/swag

`swag init` this places the static docs in the docs directory which is served on the /docs/index.html

# TODO
* use casbin for role based security https://github.com/gin-contrib/authz

# Considerations

* experimenting with hydra https://github.com/ory/hydra
* experimenting with graphql https://github.com/99designs/gqlgen

