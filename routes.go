package main

import "github.com/gin-gonic/gin"

func setupRoutes(r *gin.Engine) {

	// API versioning
	v := r.Group("/v1")

	o := r.Group("/oauth")

	// api group requires user authentication
	v.Use(AuthRequired())

	// Application Routes
	o.GET("/access_token", GetAccessToken)

	v.GET("/users", ListUser)
	v.POST("/users", CreateUser)
	v.GET("/users/:id", GetUser)

	v.GET("/drops", ListDrop)
	v.POST("/drops", CreateDrop)
	v.GET("/drops/:id", GetDrop)

	v.GET("/clients", ListClient)
	v.POST("/clients", CreateClient)
	v.GET("/clients/:id", GetClient)

	v.GET("/account", ListAccount)
	v.PUT("/account", UpdateAccount)

	v.GET("/media/:id", GetMedia)
	v.POST("/media", UploadMedia)
}
