package account

import "github.com/gin-gonic/gin"

func List(c *gin.Context) {
	u := c.MustGet("user").(gin.H)
	c.JSON(200, u)
}

func Update(c *gin.Context) {}
