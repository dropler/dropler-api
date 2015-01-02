package clients

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// List Method for returning all clients as array of objects
func List(c *gin.Context) {
	d := ClientList{}

	err := d.List()
	if err != nil {
		c.JSON(500, gin.H{"error": "Problem fetching clients. No database connection?"})
		return
	}

	c.JSON(200, d)
	return
}

// Create Method for creating a new client
func Create(c *gin.Context) {
	d := Client{}

	// Bind the incoming form data with the client model
	c.BindWith(&d, binding.Form)

	err := d.Create()
	if err != nil {
		panic(err)
	}

	// Return the client model as json to the client
	c.JSON(200, d)
	return
}

func GetClient(c *gin.Context) {
	d := Client{}
	id := c.Params.ByName("id")

	err := d.GetById(id)
	if err != nil {
		c.JSON(500, gin.H{"error": "Problem fetching client. No database connection?"})
		return
	}

	c.JSON(200, d)
}
