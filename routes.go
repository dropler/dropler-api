package main

import (
	"dropler-new/clients"
	"dropler-new/drops"
	"dropler-new/users"
	"github.com/gin-gonic/gin"
)

func setupRoutes(r *gin.Engine) {
	apiGroup := r.Group("/api")

	clients.SetupRoutes(apiGroup)
	drops.SetupRoutes(apiGroup)
	users.SetupRoutes(apiGroup)
}
