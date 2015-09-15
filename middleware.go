package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Authentication Middleware
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the code from either the post body or query param
		code := c.PostForm("access_token")
		if code == "" {
			code = c.Query("access_token")
		}

		u, err := GetUserByAccessToken(code)

		if err != nil {
			fmt.Println("Error: Failed to get user for access token: ", code, err)
			c.JSON(401, gin.H{"message": "Not authorized", "error": err.Error()})
			c.AbortWithStatus(401)
			return
		}

		//TODO: Authenticate user and token and load user into context state c.Set('user', User{})
		c.Set("user", u)
	}
}
