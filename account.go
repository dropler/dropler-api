package main

import "github.com/gin-gonic/gin"

func ListAccount(c *gin.Context) {
	u := c.MustGet("user").(gin.H)
	c.JSON(200, u)
}

func UpdateAccount(c *gin.Context) {}
