package clients

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.RouterGroup) {
	clientGroup := r.Group("/clients")
	clientGroup.GET("", List)
	clientGroup.POST("", Create)
	clientGroup.GET("/:id", GetClient)
}
