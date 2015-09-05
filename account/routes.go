package account

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.RouterGroup) {
	userGroup := r.Group("/account")
	userGroup.GET("", List)
	userGroup.PUT("", Update)
}
