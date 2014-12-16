package users

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// List Method for returning all users as array of objects
func List(c *gin.Context) {
	fmt.Println("In List Controller")
	u := UserList{}

	err := u.List()
	if err != nil {
		c.JSON(500, gin.H{"error": "Problem fetching users. No database connection?"})
	}

	c.JSON(200, u)
}

func Create(c *gin.Context) {
	u := User{}

	c.BindWith(&u, binding.Form)

	// Get password from form value directly and
	// not through the BindWith method.
	password := c.Request.FormValue("password")

	err := u.Insert(password)
	if err != nil {
		c.JSON(500, gin.H{"error": "Problem fetching users. No database connection?"})
		return
	}

	c.JSON(200, u)
}

func GetUser(c *gin.Context) {
	u := User{}
	id := c.Params.ByName("id")

	err := u.GetById(id)
	if err != nil {
		c.JSON(500, gin.H{"error": "Problem fetching users. No database connection?"})
		return
	}

	c.JSON(200, u)
}
