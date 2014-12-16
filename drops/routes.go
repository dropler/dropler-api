package drops

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.RouterGroup) {
	dropGroup := r.Group("/drops")
	dropGroup.GET("", List)
	dropGroup.POST("", Create)
	dropGroup.GET("/:id", GetDrop)
}
