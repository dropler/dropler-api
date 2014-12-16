package drops

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// List Method for returning all users as array of objects
func List(c *gin.Context) {
	fmt.Println("In Drop List Controller")
	d := DropList{}

	err := d.List()
	if err != nil {
		c.JSON(500, gin.H{"error": "Problem fetching drops. No database connection?"})
	}

	c.JSON(200, d)
}

func Create(c *gin.Context) {
	d := Drop{}

	c.BindWith(&d, binding.Form)

	err := d.Insert()
	if err != nil {
		c.JSON(500, gin.H{"error": "Problem fetching drops. No database connection?"})
		return
	}

	c.JSON(200, d)
}

func GetDrop(c *gin.Context) {
	d := Drop{}
	id := c.Params.ByName("id")

	err := d.GetById(id)
	if err != nil {
		c.JSON(500, gin.H{"error": "Problem fetching drops. No database connection?"})
		return
	}

	c.JSON(200, d)
}
