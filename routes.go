package main

import (
	"dropler-new/users"
	"github.com/gin-gonic/gin"
)

func setupRoutes(r *gin.Engine) {
	apiGroup := r.Group("/api")

	users.SetupRoutes(apiGroup)
}
