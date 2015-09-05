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
	apiGroup := r.Group("/api")
	oauthGroup := r.Group("/oauth")

	// api group requires user authentication
	apiGroup.Use(AuthRequired())

	clients.SetupRoutes(apiGroup)
	drops.SetupRoutes(apiGroup)
	account.SetupRoutes(apiGroup)
	users.SetupRoutes(apiGroup)

	oauth.SetupRoutes(oauthGroup)
}
