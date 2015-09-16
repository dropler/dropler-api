package main

import (
	"fmt"

	"github.com/RangelReale/osin"
	"github.com/gin-gonic/gin"
)

// Authentication Middleware
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Get the code from either the post body or query param
		code := c.Query("access_token")

		if b := osin.CheckBearerAuth(c.Request); b != nil {
			code = b.Code
		}

		if code == "" {
			c.JSON(401, gin.H{"message": "Not authorized", "error": "No access_token provided"})
			c.AbortWithStatus(401)
			return
		}

		u, err := GetUserByAccessToken(code)

		if err != nil {
			fmt.Println("Error: Failed to get user for access token: ", code, err)
			c.JSON(401, gin.H{"message": "Not authorized", "error": err.Error()})
			c.AbortWithStatus(401)
			return
		}

		c.Set("user", u)
	}
}
