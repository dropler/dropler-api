package main

import "github.com/gin-gonic/gin"

// Authentication Middleware
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		//TODO: Authenticate user and token and load user into context state c.Set('user', User{})
		c.Set("user", gin.H{"Hello": "World"})
	}
}
