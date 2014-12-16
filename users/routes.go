package users

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.RouterGroup) {
	userGroup := r.Group("/users")
	userGroup.GET("", List)
	userGroup.POST("", Create)
	userGroup.GET("/:id", GetUser)
}
