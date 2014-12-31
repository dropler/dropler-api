package oauth

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.RouterGroup) {
	r.GET("/authorize", Authorize)
	r.GET("/token", Token)
}
