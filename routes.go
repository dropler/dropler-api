package main

import (
	"dropler-new/clients"
	"dropler-new/drops"
	"dropler-new/oauth"
	"dropler-new/users"
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
