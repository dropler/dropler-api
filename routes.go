package main

import (
	"dropler/clients"
	"dropler/drops"
	"dropler/oauth"
	"dropler/users"
	"github.com/gin-gonic/gin"
)

func setupRoutes(r *gin.Engine) {
	apiGroup := r.Group("/api")
	oauthGroup := r.Group("/oauth")

	clients.SetupRoutes(apiGroup)
	drops.SetupRoutes(apiGroup)
	users.SetupRoutes(apiGroup)

	oauth.SetupRoutes(oauthGroup)
}
