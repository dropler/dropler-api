package main

import (
	"dropler/account"
	"dropler/clients"
	"dropler/drops"
	"dropler/oauth"
	"dropler/users"

	"github.com/gin-gonic/gin"
)

func setupRoutes(r *gin.Engine) {

	// API versioning
	v := r.Group("/v1")

	apiGroup := v.Group("/api")
	oauthGroup := v.Group("/oauth")

	// api group requires user authentication
	apiGroup.Use(AuthRequired())

	clients.SetupRoutes(apiGroup)
	drops.SetupRoutes(apiGroup)
	account.SetupRoutes(apiGroup)
	users.SetupRoutes(apiGroup)

	oauth.SetupRoutes(oauthGroup)
}
